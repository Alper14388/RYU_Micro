package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/netrack/openflow/ofp"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	pb "sdn/common/proto"
	"sync"
	"time"
)

type server struct {
	pb.UnimplementedConnectionManagerServer
	store *Store
}

type Store struct {
	mu   sync.Mutex
	conn net.Conn
}

var store = &Store{}

func (s *server) SendFlowMod(ctx context.Context, req *pb.FlowModRequest) (*pb.FlowModResponse, error) {
	var flowMod ofp.FlowMod
	if err := json.Unmarshal(req.Data, &flowMod); err != nil {
		return &pb.FlowModResponse{
			Success: false,
			Message: fmt.Sprintf("Invalid FlowMod data: %v", err),
		}, nil
	}

	s.store.mu.Lock()
	c := s.store.conn
	s.store.mu.Unlock()

	if c == nil {
		return &pb.FlowModResponse{
			Success: false,
			Message: "No switch connection available",
		}, nil
	}

	if _, err := flowMod.WriteTo(c); err != nil {
		return &pb.FlowModResponse{
			Success: false,
			Message: fmt.Sprintf("Error sending FlowMod: %v", err),
		}, nil
	}

	return &pb.FlowModResponse{
		Success: true,
		Message: "FlowMod sent successfully",
	}, nil
}

func (s *server) SendPacketOut(ctx context.Context, req *pb.PacketOutRequest) (*pb.PacketOutResponse, error) {
	var pktOut ofp.PacketOut
	if err := json.Unmarshal(req.Data, &pktOut); err != nil {
		return &pb.PacketOutResponse{
			Success: false,
			Message: fmt.Sprintf("Invalid PacketOut data: %v", err),
		}, nil
	}

	s.store.mu.Lock()
	c := s.store.conn
	s.store.mu.Unlock()

	if c == nil {
		return &pb.PacketOutResponse{
			Success: false,
			Message: "No switch connection available",
		}, nil
	}

	if _, err := pktOut.WriteTo(c); err != nil {
		return &pb.PacketOutResponse{
			Success: false,
			Message: fmt.Sprintf("Error sending PacketOut: %v", err),
		}, nil
	}

	return &pb.PacketOutResponse{
		Success: true,
		Message: "PacketOut sent successfully",
	}, nil
}

func (s *server) HandlePacketIn(ctx context.Context, req *pb.PacketInRequest) (*pb.PacketInResponse, error) {
	// This will be called by Ryu_go service
	return &pb.PacketInResponse{
		Success: true,
		Message: "PacketIn handled",
	}, nil
}

func main() {
	// Start OpenFlow listener
	go listenAndServeOpenFlow(":6633")

	// Start gRPC server
	lis, err := net.Listen("tcp", ":8094")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterConnectionManagerServer(s, &server{store: store})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// ==============================
// 1) Store yapısı
// ==============================
func listenAndServeOpenFlow(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}
	log.Printf("[CM] Listening OpenFlow on %s...\n", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		log.Println("[CM] New Switch connection from", conn.RemoteAddr())

		// (Opsiyonel) Birden fazla switch bağlanırsa, her accept için
		// ayrı bir goroutine açıp handshake yaparsınız.
		go handleSwitchConnection(conn)
	}
}

// handleSwitchConnection: Basit OpenFlow handshake + read-loop
func handleSwitchConnection(conn net.Conn) {
	// 3.1) Minimal handshake: Okuyup "Hello" al, "Hello" gönder, "FeaturesRequest" vs.
	if err := doMinimalHandshake(conn); err != nil {
		log.Println("Handshake failed:", err)
		conn.Close()
		return
	}
	log.Println("[CM] Handshake OK with switch", conn.RemoteAddr())

	// 3.2) Bu noktada handshake tamam, store.conn'a kaydediyoruz (tek switch varsayımı).
	store.mu.Lock()
	// varsa eski connection'ı kapatabiliriz (örneğin sadece 1 switch'e izin veriyoruz)
	if store.conn != nil {
		store.conn.Close()
	}
	store.conn = conn
	store.mu.Unlock()

	// 3.3) readFromSwitch: PacketIn (veya diğer OpenFlow msg) bekler
	readFromSwitch(conn)
}

// doMinimalHandshake: Hello al, Hello gönder, FeaturesRequest gönder, FeaturesReply bekle
func doMinimalHandshake(conn net.Conn) error {
	// 1) Switch'ten Hello (header) al (8 byte)
	hdr := make([]byte, 8)
	if _, err := io.ReadFull(conn, hdr); err != nil {
		return fmt.Errorf("error reading Hello: %v", err)
	}
	// header[0] = version, header[1] = type=0(Hello?), vs.
	// Bu örnekte kontrolleri minimal tutuyoruz
	log.Printf("[CM] Received Hello: version=%d, type=%d\n", hdr[0], hdr[1])

	// 2) Biz de Hello göndereceğiz (8 byte, type=0)
	hello := make([]byte, 8)
	hello[0] = hdr[0] // version
	hello[1] = 0      // type = HELLO
	binary.BigEndian.PutUint16(hello[2:], 8)
	// xid
	copy(hello[4:], []byte{0x11, 0x22, 0x33, 0x44})

	if _, err := conn.Write(hello); err != nil {
		return fmt.Errorf("error writing Hello: %v", err)
	}
	log.Println("[CM] Sent HelloReply")

	// 3) FeaturesRequest
	featReq := make([]byte, 8)
	featReq[0] = hdr[0] // version
	featReq[1] = 5      // type=FeaturesRequest(5)
	binary.BigEndian.PutUint16(featReq[2:], 8)
	copy(featReq[4:], []byte{0x00, 0x00, 0x00, 0x06})

	if _, err := conn.Write(featReq); err != nil {
		return fmt.Errorf("error writing FeatureRequest: %v", err)
	}
	log.Println("[CM] Sent FeatureRequest")

	// 4) Switch'ten FeaturesReply beklemek optional.
	//    1 sn timeout vs.
	_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	featReplyHdr := make([]byte, 8)
	if _, err := io.ReadFull(conn, featReplyHdr); err != nil {
		log.Println("[CM] Timeout or error reading FeaturesReply (maybe switch didn't send?) - continuing anyway:", err)
		// Optional: not returning error, let's keep going.
	} else {
		log.Printf("[CM] Received FeaturesReply: version=%d, type=%d\n", featReplyHdr[0], featReplyHdr[1])
	}
	conn.SetReadDeadline(time.Time{}) // reset

	return nil
}

// readFromSwitch: Sürekli mesaj bekler, PacketIn veya EchoRequest vb. parse edebilirsiniz
func readFromSwitch(conn net.Conn) {
	defer conn.Close()

	for {
		// 1) Header oku (8 byte)
		hdr := make([]byte, 8)
		_, err := io.ReadFull(conn, hdr)
		if err != nil {
			log.Println("Switch read error:", err)
			return
		}
		msgType := hdr[1]
		length := binary.BigEndian.Uint16(hdr[2:4])
		xid := binary.BigEndian.Uint32(hdr[4:8])

		// 2) Body oku (length - 8)
		bodyLen := int(length) - 8
		body := make([]byte, 0)
		if bodyLen > 0 {
			body = make([]byte, bodyLen)
			if _, err := io.ReadFull(conn, body); err != nil {
				log.Println("Error reading msg body:", err)
				return
			}
		}

		// 3) Mesaj tipine göre işlem
		switch msgType {
		case 10: // OFPT_PACKET_IN
			// PacketIn parsing: ofp.PacketIn
			var pktIn ofp.PacketIn
			if _, err := pktIn.ReadFrom(bytes.NewReader(append(hdr, body...))); err != nil {
				log.Println("Error parsing PacketIn:", err)
				continue
			}
			log.Println("[CM] PacketIn => forwarding to microservice. Xid=", xid)
			forwardPacketIn(pktIn)

		case 2: // EchoRequest
			log.Println("[CM] Received EchoRequest, replying EchoReply")
			// Minimal echo reply
			echoReply := make([]byte, 8)
			echoReply[0] = hdr[0] // version
			echoReply[1] = 3      // type=EchoReply
			binary.BigEndian.PutUint16(echoReply[2:], 8)
			copy(echoReply[4:], hdr[4:8]) // same Xid
			if _, err := conn.Write(echoReply); err != nil {
				log.Println("Error writing EchoReply:", err)
			}

		default:
			log.Printf("[CM] Received msgType=%d, length=%d -> skipping\n", msgType, length)
		}
	}
}

// forwardPacketIn: PacketIn microservisine JSON post
func forwardPacketIn(pktIn ofp.PacketIn) {
	data, err := json.Marshal(pktIn)
	if err != nil {
		log.Println("PacketIn marshal error:", err)
		return
	}
	url := "http://127.0.0.1:8090/packetin"
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Println("Forward PacketIn error:", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("[CM] Forwarded PacketIn to %s, got status=%s\n", url, resp.Status)
}

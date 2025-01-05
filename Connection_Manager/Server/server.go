package Server

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
	"sdn/Connection_Manager/utils"
	pb "sdn/common/proto"
	"sync"
	"time"
)

type Server struct {
	pb.UnimplementedConnectionManagerServer
	Store *Store
}

type Store struct {
	mu   sync.Mutex
	conn net.Conn
}

var tempMod struct {
	Buffer      uint32    `json:"Buffer"`
	Command     uint32    `json:"Command"`
	Match       ofp.Match `json:"Match"`
	IdleTimeout uint16    `json:"IdleTimeout"`
	HardTimeout uint16    `json:"HardTimeout"`
	Priority    uint16    `json:"Priority"`
}
var store = &Store{}

func ListenAndServeOpenFlow(addr string) {
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

		// Birden fazla switch bağlanırsa, her accept için
		// ayrı bir goroutine
		go handleSwitchConnection(conn)
	}
}

func handleSwitchConnection(conn net.Conn) {
	if err := doMinimalHandshake(conn); err != nil {
		log.Println("Handshake failed:", err)
		conn.Close()
		return
	}
	log.Println("[CM] Handshake OK with switch", conn.RemoteAddr())
	store.mu.Lock()
	if store.conn != nil {
		store.conn.Close()
	}
	store.conn = conn
	store.mu.Unlock()
	if err := buildDefaultFlow(conn); err != nil {
		log.Println("Build default flow failed:", err)
	}
	readFromSwitch(conn)
}

func doMinimalHandshake(conn net.Conn) error {
	hdr := make([]byte, 8)
	if _, err := io.ReadFull(conn, hdr); err != nil {
		return fmt.Errorf("error reading Hello: %v", err)
	}

	log.Printf("[CM] Received Hello: version=%d, type=%d\n", hdr[0], hdr[1])

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
	featReq[0] = hdr[0]
	featReq[1] = 5
	binary.BigEndian.PutUint16(featReq[2:], 8)
	copy(featReq[4:], []byte{0x00, 0x00, 0x00, 0x06})

	if _, err := conn.Write(featReq); err != nil {
		return fmt.Errorf("error writing FeatureRequest: %v", err)
	}
	log.Println("[CM] Sent FeatureRequest")

	_ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	featReplyHdr := make([]byte, 8)
	if _, err := io.ReadFull(conn, featReplyHdr); err != nil {
		log.Println("[CM] Timeout or error reading FeaturesReply (maybe switch didn't send?) - continuing anyway:", err)
	} else {
		log.Printf("[CM] Received FeaturesReply: version=%d, type=%d\n", featReplyHdr[0], featReplyHdr[1])
	}
	conn.SetReadDeadline(time.Time{}) // reset

	return nil
}

func readFromSwitch(conn net.Conn) {
	defer conn.Close()

	for {
		hdr := make([]byte, 8)
		_, err := io.ReadFull(conn, hdr)
		if err != nil {
			log.Println("Switch read error:", err)
			return
		}
		msgType := hdr[1]
		length := binary.BigEndian.Uint16(hdr[2:4])
		xid := binary.BigEndian.Uint32(hdr[4:8])

		bodyLen := int(length) - 8
		body := make([]byte, 0)
		if bodyLen > 0 {
			body = make([]byte, bodyLen)
			if _, err := io.ReadFull(conn, body); err != nil {
				log.Println("Error reading msg body:", err)
				return
			}
		}

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

func forwardPacketIn(pktIn ofp.PacketIn) {
	log.Printf("Forwarding PacketIn via gRPC: %+v", pktIn)

	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to PacketHandler service: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewPacketHandlerClient(conn)

	req := &pb.PacketInRequest{
		BufferId: pktIn.Buffer,
		Data:     pktIn.Data,
		Cookie:   pktIn.Cookie,
		TableId:  uint32(pktIn.Table),
	}

	resp, err := client.HandlePacketIn(context.Background(), req)
	if err != nil {
		log.Printf("Error sending PacketIn via gRPC: %v", err)
		return
	}

	log.Printf("PacketIn successfully forwarded via gRPC: %+v", resp)
}

func (s *Server) SendFlowMod(ctx context.Context, req *pb.FlowModRequest) (*pb.FlowModResponse, error) {

	if err := json.Unmarshal(req.Data, &tempMod); err != nil {
		log.Printf("Failed to unmarshal FlowMod data: %v", err)
		return &pb.FlowModResponse{
			Success: false,
			Message: fmt.Sprintf("Invalid FlowMod data: %v", err),
		}, nil
	}

	instructions, err := utils.GetInstruction(req)
	if err != nil {
		log.Printf("Failed to get instructions: %v", err)
		return &pb.FlowModResponse{
			Success: false,
			Message: "get instructions error",
		}, err
	}

	flowMod := ofp.FlowMod{
		Buffer:       tempMod.Buffer,
		Command:      ofp.FlowModCommand(tempMod.Command),
		Match:        tempMod.Match,
		IdleTimeout:  tempMod.IdleTimeout,
		HardTimeout:  tempMod.HardTimeout,
		Priority:     tempMod.Priority,
		Instructions: instructions,
	}
	flowMod.Buffer = ofp.NoBuffer
	s.Store.mu.Lock()
	c := s.Store.conn
	s.Store.mu.Unlock()
	if c == nil {
		return &pb.FlowModResponse{
			Success: false,
			Message: "No switch connection available",
		}, nil
	}
	if err := utils.SendHeaderandBuffer(c, flowMod); err != nil {
		return &pb.FlowModResponse{
			Success: false,
			Message: "FlowMod sent failed",
		}, err
	}

	return &pb.FlowModResponse{
		Success: true,
		Message: "FlowMod sent successfully",
	}, nil
}

func (s *Server) HandlePacketIn() (*pb.PacketInResponse, error) {
	return &pb.PacketInResponse{
		Success: true,
		Message: "PacketIn handled",
	}, nil
}

func GetStore() *Store {
	return store
}

func buildDefaultFlow(conn net.Conn) error {
	defaultFlow := ofp.FlowMod{
		Buffer:   ofp.NoBuffer,
		Table:    0,
		Command:  ofp.FlowAdd,
		Priority: 0,
		Match: ofp.Match{
			Type:   ofp.MatchTypeXM,
			Fields: []ofp.XM{},
		},
		IdleTimeout: 0,
		HardTimeout: 0,
		Instructions: ofp.Instructions{
			&ofp.InstructionApplyActions{
				Actions: []ofp.Action{
					&ofp.ActionOutput{
						Port:   ofp.PortController,
						MaxLen: 65535,
					},
				},
			},
		},
	}

	if err := utils.SendHeaderandBuffer(conn, defaultFlow); err != nil {
		log.Printf("Failed to send default FlowMod: %v", err)
		return err
	}

	log.Println("[CM] Default FlowMod sent successfully")
	return nil
}

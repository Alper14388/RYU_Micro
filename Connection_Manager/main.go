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
	var tempMod struct {
		Buffer      uint32    `json:"Buffer"`
		Command     uint32    `json:"Command"`
		Match       ofp.Match `json:"Match"`
		IdleTimeout uint16    `json:"IdleTimeout"`
		HardTimeout uint16    `json:"HardTimeout"`
		Priority    uint16    `json:"Priority"`
	}

	if err := json.Unmarshal(req.Data, &tempMod); err != nil {
		log.Printf("Failed to unmarshal FlowMod data: %v", err)
		return &pb.FlowModResponse{
			Success: false,
			Message: fmt.Sprintf("Invalid FlowMod data: %v", err),
		}, nil
	}

	var instructions ofp.Instructions
	for _, inst := range req.Instructions {
		switch inst.Type {
		case uint32(ofp.InstructionTypeApplyActions): // ApplyActions
			var actions ofp.Actions
			for _, action := range inst.Actions {
				switch action.Type {
				case uint32(ofp.ActionTypeOutput): // ActionOutput
					actions = append(actions, &ofp.ActionOutput{
						Port:   ofp.PortNo(action.Port),
						MaxLen: uint16(action.MaxLen),
					})
				default:
					log.Printf("Unsupported action type: %v", action.Type)
					return &pb.FlowModResponse{
						Success: false,
						Message: fmt.Sprintf("Unsupported action type: %v", action.Type),
					}, nil
				}
			}
			instructions = append(instructions, &ofp.InstructionApplyActions{
				Actions: actions,
			})
		default:
			log.Printf("Unsupported instruction type: %v", inst.Type)
			return &pb.FlowModResponse{
				Success: false,
				Message: fmt.Sprintf("Unsupported instruction type: %v", inst.Type),
			}, nil
		}
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

	s.store.mu.Lock()
	c := s.store.conn
	s.store.mu.Unlock()
	if c == nil {
		return &pb.FlowModResponse{
			Success: false,
			Message: "No switch connection available",
		}, nil
	}
	flowMod.Buffer = ofp.NoBuffer
	header := make([]byte, 8)
	header[0] = 4
	header[1] = 14
	size, _ := calculateFlowModSize(&flowMod)
	binary.BigEndian.PutUint16(header[2:], uint16(8+size)) // Header + Body uzunluğu
	binary.BigEndian.PutUint32(header[4:], uint32(12345))  // Xid örnek olarak 12345

	var buf bytes.Buffer
	if _, err := flowMod.WriteTo(&buf); err != nil {
		return &pb.FlowModResponse{
			Success: false,
			Message: fmt.Sprintf("Error sending FlowMod: %v", err),
		}, nil
	}

	if _, err := c.Write(append(header, buf.Bytes()...)); err != nil {
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

func (s *server) HandlePacketIn() (*pb.PacketInResponse, error) {
	return &pb.PacketInResponse{
		Success: true,
		Message: "PacketIn handled",
	}, nil
}

func main() {
	go listenAndServeOpenFlow(":6633")

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

	// Serialize PacketIn into JSON
	data, err := json.Marshal(pktIn)
	if err != nil {
		log.Println("PacketIn marshal error:", err)
		return
	}

	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to PacketHandler service: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewPacketHandlerClient(conn)

	req := &pb.PacketInRequest{
		BufferId: pktIn.Buffer,
		Data:     data,
	}

	resp, err := client.HandlePacketIn(context.Background(), req)
	if err != nil {
		log.Printf("Error sending PacketIn via gRPC: %v", err)
		return
	}

	log.Printf("PacketIn successfully forwarded via gRPC: %+v", resp)
}

func calculateFlowModSize(flowMod *ofp.FlowMod) (int, error) {
	var buf bytes.Buffer
	if _, err := flowMod.WriteTo(&buf); err != nil {
		return 0, err
	}
	return buf.Len(), nil
}

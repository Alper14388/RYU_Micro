package ofp

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"

	fopb "FlowOperation/proto"
	"Ryu_go/common"
	pb "Ryu_go/proto"
	"Ryu_go/utils"
)

const (
	OFPPFlood   = 0xfffffffb
	OFPNoBuffer = 0xffffffff
	FlowOpAddr  = "localhost:8092"
)

var macToPort = make(map[uint64]map[string]uint32)
var macToPortLock sync.RWMutex

func updateMacToPort(dpid uint64, src string, inPort uint32) {
	macToPortLock.Lock()
	defer macToPortLock.Unlock()

	log.Println("updateMacToPort", src, inPort)

	if _, exists := macToPort[dpid]; !exists {
		macToPort[dpid] = make(map[string]uint32)
	}

	macToPort[dpid][src] = inPort

	log.Printf("Updated MAC-to-Port mapping: DPID=%d, SRC=%s, IN_PORT=%d", dpid, src, inPort)
}

func outPortLookup(dpid uint64, dst string) uint32 {
	macToPortLock.RLock()
	defer macToPortLock.RUnlock()
	log.Println("outPortLookup", dst)
	if ports, exists := macToPort[dpid]; exists {
		if outPort, found := ports[dst]; found {
			log.Printf("Out port found: DPID=%d, DST=%s, OUT_PORT=%d", dpid, dst, outPort)
			return outPort
		}
	}

	log.Printf("Out port not found, flooding: DPID=%d, DST=%s", dpid, dst)
	return OFPPFlood
}

func HandlePacketInGRPC(req *pb.PacketInRequest) (*pb.PacketInResponse, error) {
	packet := common.PacketInWrapper{
		Buffer: req.BufferId,
		Length: uint16(req.Length),
		Reason: uint8(req.Reason),
		Table:  uint8(req.TableId),
		Cookie: req.Cookie,
		Match: struct {
			Type   uint16 `json:"Type"`
			Fields []struct {
				Class uint16      `json:"Class"`
				Type  uint8       `json:"Type"`
				Value string      `json:"Value"`
				Mask  interface{} `json:"Mask"`
			} `json:"Fields"`
		}{
			Type: uint16(req.Match.Type),
		},
		Data: req.Data,
	}

	// Convert match fields
	for _, field := range req.Match.Fields {
		packet.Match.Fields = append(packet.Match.Fields, struct {
			Class uint16      `json:"Class"`
			Type  uint8       `json:"Type"`
			Value string      `json:"Value"`
			Mask  interface{} `json:"Mask"`
		}{
			Class: uint16(field.Class),
			Type:  uint8(field.Type),
			Value: field.Value,
			Mask:  field.Mask,
		})
	}

	log.Printf("Packet-In received: %+v\n", packet)

	packetInfo, err := utils.ExtractData(packet)
	if err != nil {
		log.Printf("Failed to extract data: %v", err)
		return &pb.PacketInResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}
	log.Printf("Extracted Packet: %+v\n", packetInfo)

	if packetInfo.IsLLDP {
		go notifyTopology(packetInfo)
		return &pb.PacketInResponse{
			Success: true,
			Message: "LLDP packet processed",
		}, nil
	}

	updateMacToPort(packetInfo.DPID, packetInfo.Src, packetInfo.InPort)

	outPort := outPortLookup(packetInfo.DPID, packetInfo.Dst)
	go addFlowEntry(packetInfo, outPort, packetInfo.InPort)
	go sendPacketOut(packetInfo, outPort, packetInfo.InPort)

	return &pb.PacketInResponse{
		Success: true,
		Message: "Packet processed successfully",
	}, nil
}

func notifyTopology(packet utils.PacketData) {
	// TODO: Convert to gRPC when topology service is updated
	log.Println("Topology notification skipped - waiting for gRPC implementation")
}

func addFlowEntry(packet utils.PacketData, outPort uint32, inPort uint32) {
	log.Println("Adding flow entry for packet.")

	conn, err := grpc.Dial(FlowOpAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to FlowOperation service: %v", err)
		return
	}
	defer conn.Close()

	client := fopb.NewFlowOperationClient(conn)
	req := &fopb.FlowAddRequest{
		SwitchId:    packet.DPID,
		Src:         packet.Src,
		Dst:         packet.Dst,
		InPort:      inPort,
		OutPort:     outPort,
		Priority:    1,
		HardTimeout: 30,
		IdleTimeout: 30,
		BufferId:    packet.BufferID,
	}

	resp, err := client.AddFlow(context.Background(), req)
	if err != nil {
		log.Printf("Error adding flow entry: %v", err)
		return
	}
	log.Printf("Flow entry added successfully: %v", resp)
}

func sendPacketOut(packet utils.PacketData, outPort uint32, inPort uint32) {
	log.Println("Sending packet out for packet.")

	conn, err := grpc.Dial(FlowOpAddr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to FlowOperation service: %v", err)
		return
	}
	defer conn.Close()

	client := fopb.NewFlowOperationClient(conn)
	req := &fopb.PacketOutRequest{
		SwitchId: packet.DPID,
		InPort:   inPort,
		OutPort:  outPort,
		Data:     packet.EncodedData,
		BufferId: packet.BufferID,
	}

	resp, err := client.SendPacketOut(context.Background(), req)
	if err != nil {
		log.Printf("Error sending packet out: %v", err)
		return
	}
	log.Printf("Packet out sent successfully: %v", resp)
}

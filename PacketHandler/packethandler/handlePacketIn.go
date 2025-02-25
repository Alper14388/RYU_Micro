package ryu_go

import (
	"context"
	"log"
	"sync"

	"sdn/Ryu_go/utils"
	pb "sdn/common/proto"

	"google.golang.org/grpc"
)

const (
	FlowOpAddr = "flowadd-service.default.svc.cluster.local:8092"
)

var macToPort = make(map[uint64]map[string]uint32)
var macToPortLock sync.RWMutex

func updateMacToPort(dpid uint64, src string, inPort uint32) {
	macToPortLock.Lock()
	defer macToPortLock.Unlock()

	if _, exists := macToPort[dpid]; !exists {
		macToPort[dpid] = make(map[string]uint32)
	}

	macToPort[dpid][src] = inPort
	log.Printf("Updated MAC-to-Port mapping: DPID=%d, SRC=%s, IN_PORT=%d", dpid, src, inPort)
}

func outPortLookup(dpid uint64, dst string) uint32 {
	macToPortLock.RLock()
	defer macToPortLock.RUnlock()

	if ports, exists := macToPort[dpid]; exists {
		if outPort, found := ports[dst]; found {
			return outPort
		}
	}
	return 0xfffffffb
}

func HandlePacketIn(req *pb.PacketInRequest) (*pb.PacketInResponse, error) {
	log.Printf("Received PacketIn: %+v\n", req)

	packetInfo, err := utils.ExtractDataFromPacketIn(req)
	if err != nil {
		log.Printf("Failed to extract data: %v", err)
		return &pb.PacketInResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	if packetInfo.IsLLDP {
		log.Printf("Ignoring LLDP packet for DPID=%d\n", packetInfo.DPID)
		return &pb.PacketInResponse{
			Success: true,
			Message: "LLDP packet ignored",
		}, nil
	}
	log.Println("packetInfo:", packetInfo)
	updateMacToPort(packetInfo.DPID, packetInfo.Src, packetInfo.InPort)

	outPort := outPortLookup(packetInfo.DPID, packetInfo.Dst)
	if err := addFlowEntry(packetInfo, outPort); err != nil {
		log.Printf("Error adding flow entry: %v", err)
	}

	return &pb.PacketInResponse{
		Success: true,
		Message: "Packet processed successfully",
	}, nil
}

func addFlowEntry(packet utils.PacketData, outPort uint32) error {

	conn, err := grpc.Dial(FlowOpAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	priority := uint32(1)
	if packet.EtherType == 0x0800 && packet.IPProto == 1 { // ICMP kontrolü
		priority = 10
	}
	client := pb.NewFlowOperationClient(conn)
	req := &pb.FlowAddRequest{
		SwitchId:    packet.DPID,
		Src:         packet.Src,
		Dst:         packet.Dst,
		InPort:      packet.InPort,
		OutPort:     outPort,
		Priority:    priority,
		HardTimeout: 300,
		IdleTimeout: 60,
		BufferId:    packet.BufferID,
		EthType:     uint64(packet.EtherType),
		IPProto:     uint64(packet.IPProto),
	}

	_, err = client.AddFlow(context.Background(), req)
	return err
}

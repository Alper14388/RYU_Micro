package utils

import (
	"fmt"
	"log"
	pb "sdn/common/proto"
)

type PacketData struct {
	BufferID  uint32 `json:"buffer_id"`
	Data      []byte `json:"data"`
	DPID      uint64 `json:"dpid"`
	IsLLDP    bool   `json:"is_lldp"`
	Dst       string `json:"dst,omitempty"`
	Src       string `json:"src,omitempty"`
	InPort    uint32 `json:"in_port,omitempty"`
	EtherType uint16 `json:"ether_type"`
	IPProto   uint8  `json:"ip_proto"`
}

type EthernetFrame struct {
	Dst       string
	Src       string
	Ethertype uint16
	IPProto   uint8
}

// EtherTypes constants
const (
	EthTypeLLDP = 0x88CC
	EthTypeIPv4 = 0x0800
	EthTypeARP  = 0x0806
)

// ExtractDataFromPacketIn extracts relevant data from the PacketInRequest.
func ExtractDataFromPacketIn(req *pb.PacketInRequest) (PacketData, error) {
	const ethernetStart = 20
	if len(req.Data) < ethernetStart+14 {
		return PacketData{}, fmt.Errorf("packet data too short to parse Ethernet frame")
	}

	frame := ParseEthernetFrame(req.Data[ethernetStart:])

	packetInfo := PacketData{
		BufferID:  req.BufferId,
		Data:      req.Data,
		DPID:      req.SwitchId,
		IsLLDP:    frame.Ethertype == EthTypeLLDP,
		InPort:    req.InPort,
		EtherType: frame.Ethertype,
	}

	if !packetInfo.IsLLDP {
		packetInfo.Dst = frame.Dst
		packetInfo.Src = frame.Src
		packetInfo.IPProto = frame.IPProto
	}

	log.Printf("Extracted Packet Data: %+v", packetInfo)
	return packetInfo, nil
}

// ParseEthernetFrame parses raw Ethernet frame bytes.
func ParseEthernetFrame(data []byte) EthernetFrame {
	if len(data) < 14 {
		log.Println("Frame too short to parse Ethernet header")
		return EthernetFrame{}
	}

	dst := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", data[0], data[1], data[2], data[3], data[4], data[5])
	src := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", data[6], data[7], data[8], data[9], data[10], data[11])
	ethertype := uint16(data[12])<<8 | uint16(data[13])

	log.Printf("Parsed Ethernet Frame: Dst=%s, Src=%s, Ethertype=0x%04x", dst, src, ethertype)

	frame := EthernetFrame{
		Dst:       dst,
		Src:       src,
		Ethertype: ethertype,
	}

	if ethertype == EthTypeIPv4 && len(data) >= 23 {
		frame.IPProto = data[23]
	}

	return frame
}

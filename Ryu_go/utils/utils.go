package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	pb "sdn/common/proto"
)

type PacketData struct {
	BufferID    uint32 `json:"buffer_id"`
	EncodedData string `json:"encoded_data"`
	Data        []byte `json:"data"`
	DPID        uint64 `json:"dpid"`
	IsLLDP      bool   `json:"is_lldp"`
	Dst         string `json:"dst,omitempty"`
	Src         string `json:"src,omitempty"`
	InPort      uint32 `json:"in_port,omitempty"`
}

type EthernetFrame struct {
	Dst       string
	Src       string
	Ethertype uint16
}

// EtherTypes constants
const (
	EthTypeLLDP = 0x88CC
)

// ExtractDataFromPacketIn extracts relevant data from the PacketInRequest.
func ExtractDataFromPacketIn(req *pb.PacketInRequest) (PacketData, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(req.Data))
	if err != nil {
		return PacketData{}, fmt.Errorf("failed to decode base64 data: %w", err)
	}

	frame := ParseEthernetFrame(decodedData)

	packetInfo := PacketData{
		BufferID:    req.BufferId,
		EncodedData: string(req.Data),
		Data:        decodedData,
		DPID:        req.SwitchId,
		IsLLDP:      frame.Ethertype == EthTypeLLDP,
		InPort:      req.InPort,
	}

	if !packetInfo.IsLLDP {
		packetInfo.Dst = frame.Dst
		packetInfo.Src = frame.Src
	}

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

	return EthernetFrame{
		Dst:       dst,
		Src:       src,
		Ethertype: ethertype,
	}
}

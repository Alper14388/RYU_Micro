package utils

import (
	"Project/common"
	"encoding/base64"
	"fmt"
	"log"
)

type PacketData struct {
	BufferID    uint32 `json:"buffer_id"`
	EncodedData string `json:"encoded_data"`
	Data        []byte `json:"data"`
	DPID        uint32 `json:"dpid"`
	IsLLDP      bool   `json:"is_lldp"`
	Dst         string `json:"dst,omitempty"`
	Src         string `json:"src,omitempty"`
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

func ExtractData(packet common.PacketInWrapper) (PacketData, error) {
	// Decode the base64 data
	decodedData, err := base64.StdEncoding.DecodeString(packet.OFPPacketIn.Data)
	if err != nil {
		return PacketData{}, fmt.Errorf("failed to decode base64 data: %w", err)
	}

	// Parse the Ethernet frame
	frame := ParseEthernetFrame(decodedData)

	// Populate PacketData
	packetInfo := PacketData{
		BufferID:    packet.OFPPacketIn.BufferID,
		EncodedData: packet.OFPPacketIn.Data,
		Data:        decodedData,
		DPID:        packet.DatapathID,
		IsLLDP:      frame.Ethertype == EthTypeLLDP,
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

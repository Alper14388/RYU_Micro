package utils

import (
	"Ryu_go/common"
	"encoding/base64"
	"fmt"
	"log"
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

// ExtractData extracts relevant data from the PacketInWrapper.
func ExtractData(packet common.PacketInWrapper) (PacketData, error) {
	// Decode the base64 data
	decodedData, err := base64.StdEncoding.DecodeString(packet.Data)
	if err != nil {
		return PacketData{}, fmt.Errorf("failed to decode base64 data: %w", err)
	}

	// Parse the Ethernet frame
	frame := ParseEthernetFrame(decodedData)

	// Extract in_port from match fields
	inPort, err := extractInPort(packet)
	if err != nil {
		log.Println("Error extracting in_port:", err)
		inPort = 0 // Default to 0 if extraction fails
	}

	// Populate PacketData
	packetInfo := PacketData{
		BufferID:    packet.Buffer,
		EncodedData: packet.Data,
		Data:        decodedData,
		DPID:        1,
		IsLLDP:      frame.Ethertype == EthTypeLLDP,
		InPort:      inPort,
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

// extractInPort extracts the in_port from match fields in the packet.
func extractInPort(packet common.PacketInWrapper) (uint32, error) {
	for _, field := range packet.Match.Fields {
		if field.Class == 32768 && field.Type == 0 { // "in_port" corresponds to Class: 32768, Type: 0
			decodedValue, err := base64.StdEncoding.DecodeString(field.Value)
			if err != nil {
				return 0, fmt.Errorf("failed to decode in_port value: %w", err)
			}
			if len(decodedValue) >= 4 {
				return uint32(decodedValue[0])<<24 | uint32(decodedValue[1])<<16 | uint32(decodedValue[2])<<8 | uint32(decodedValue[3]), nil
			}
		}
	}
	return 0, fmt.Errorf("in_port not found in match fields")
}

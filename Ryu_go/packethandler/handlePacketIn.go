package ofp

import (
	"Project/common"
	"Project/utils"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const (
	OFPPFlood    = 0xfffffffb
	OFPNoBuffer  = 0xffffffff
	TopologyURL  = "http://127.0.0.1:8093/topology"
	AddFlowURL   = "http://127.0.0.1:8092/flowadd"
	PacketOutUrl = "http://127.0.0.1:8092/packetout"
)

type FlowEntry struct {
	MatchFields map[string]interface{} `json:"match_fields"`
	Actions     []string               `json:"actions"`
	TableID     uint8                  `json:"table_id"`
	Priority    uint16                 `json:"priority"`
}

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

func HandlePacketIn(w http.ResponseWriter, r *http.Request) {
	var packet common.PacketInWrapper
	if err := json.NewDecoder(r.Body).Decode(&packet); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid Packet-In format", http.StatusBadRequest)
		return
	}
	log.Printf("Packet-In received: %+v\n", packet)

	packetInfo, err := utils.ExtractData(packet)
	if err != nil {
		log.Fatalf("Failed to extract data: %v", err)
	}
	log.Printf("Extracted Packet: %+v\n", packetInfo)

	if packetInfo.IsLLDP {
		go notifyTopology(packetInfo)
		return
	}

	updateMacToPort(packetInfo.DPID, packetInfo.Src, extractInPort(packet))

	outPort := outPortLookup(packetInfo.DPID, packetInfo.Dst)
	go addFlowEntry(packetInfo, outPort, extractInPort(packet))
	go sendPacketOut(packetInfo, outPort, extractInPort(packet))
	w.WriteHeader(http.StatusOK)
}

func notifyTopology(packet utils.PacketData) {
	log.Println("Notifying topology service for LLDP packet.")
	data, _ := json.Marshal(packet)
	resp, err := http.Post(TopologyURL, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Printf("Error notifying topology service: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Println("Topology service notified successfully.")
}

func addFlowEntry(packet utils.PacketData, outPort uint32, inPort uint32) {
	log.Println("Adding flow entry for packet.")
	flow := map[string]interface{}{
		"switch_id":    packet.DPID,
		"src":          packet.Src,
		"dst":          packet.Dst,
		"in_port":      inPort,
		"out_port":     outPort,
		"priority":     1,
		"hard_timeout": 30,
		"idle_timeout": 30,
		"bufferID":     packet.BufferID,
	}
	data, _ := json.Marshal(flow)
	resp, err := http.Post(AddFlowURL, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Printf("Error adding flow entry: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Println("Flow entry added successfully.")
}

func sendPacketOut(packet utils.PacketData, outPort uint32, inPort uint32) {
	log.Println("Sending packet out for packet.")
	packetBuild := map[string]interface{}{
		"switch_id": packet.DPID,
		"in_port":   inPort,
		"out_port":  outPort,
		"data":      packet.Data,
		"buffer_id": packet.BufferID,
	}
	data, _ := json.Marshal(packetBuild)
	resp, err := http.Post(PacketOutUrl, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Printf("Error sending packet out: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Println("Packet out sent successfully.")
}

func extractInPort(packet common.PacketInWrapper) uint32 {
	for _, field := range packet.OFPPacketIn.Match.OFPMatch.OxmFields {
		if field.OXMTlv.Field == "in_port" {
			if inPort, ok := field.OXMTlv.Value.(float64); ok {
				return uint32(inPort)
			}
		}
	}
	return 0
}

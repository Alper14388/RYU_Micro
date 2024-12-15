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
	OFPPFlood   = 0xfffffffb
	OFPNoBuffer = 0xffffffff
	TopologyURL = "http://127.0.0.1:8093/topology"
	AddFlowURL  = "http://127.0.0.1:8092/addflow"
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

	updateMacToPort(packetInfo.DPID, packetInfo.Src, packet.OFPPacketIn.InPort)

	outPort := outPortLookup(packetInfo.DPID, packetInfo.Dst)

	if outPort != OFPPFlood {
		go addFlowEntry(packetInfo, outPort, packet.OFPPacketIn.InPort)
	}
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
		"hard_timeout": 5,
		"idle_timeout": 5,
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

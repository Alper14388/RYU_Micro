package ofp

import (
	"Project/common"
	"Project/utils"
	"encoding/json"
	"log"
	"net/http"
)

type FlowEntry struct {
	MatchFields map[string]interface{} `json:"match_fields"`
	Actions     []string               `json:"actions"`
	TableID     uint8                  `json:"table_id"`
	Priority    uint16                 `json:"priority"`
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
	if packet.OFPPacketIn.Reason == 0 {
		//sendPacketOut(packet)
		//addFlowEntry(packet)
	}
}

/*func sendPacketOut(packet PacketInData) {
	packetOut := ofp.PacketOut{
		BufferID: packet.BufferID,
		InPort:   ofp.Port(packet.InPort),
		Actions: []ofp.Action{
			ofp.ActionOutput{Port: ofp.PortFlood},
		},
		Data: []byte(packet.Data),
	}

	//sendToSwitch(packetOut)
}*/

/*func addFlowEntry(packet PacketInData) {
	// Flow ekleme işlemi
	flow := FlowEntry{
		MatchFields: map[string]interface{}{
			"in_port": packet.InPort,
		},
		Actions:  []string{"output:FLOOD"},
		TableID:  0,
		Priority: 1,
	}

	//sendFlowMod(flow)
}*/

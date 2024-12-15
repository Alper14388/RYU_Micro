package packetout

import (
	switchop "FlowOperation/flowops/switch"
	"encoding/json"
	"github.com/netrack/openflow/ofp"
	"log"
	"net/http"
)

type PacketOutRequest struct {
	SwitchID uint64 `json:"switch_id"`
	InPort   uint32 `json:"in_port"`
	OutPort  uint32 `json:"out_port"`
	Data     string `json:"data"`
	BufferID uint32 `json:"buffer_id"`
}

func PacketOut(w http.ResponseWriter, r *http.Request) {
	log.Println("PacketOut Endpoint Hit")
	var request PacketOutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid packet-out request format", http.StatusBadRequest)
		return
	}
	packetOut := newPacket(request)
	if err := switchop.SendToSwitch(request.SwitchID, &packetOut); err != nil {
		http.Error(w, "Failed to send packet out", http.StatusInternalServerError)
		log.Println("Error sending PacketOut to switch:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("Packet sent out successfully:", request)
}

func newPacket(request PacketOutRequest) ofp.PacketOut {
	packetOut := ofp.PacketOut{
		Buffer: request.BufferID,
		InPort: ofp.PortNo(request.InPort),
		Actions: []ofp.Action{
			&ofp.ActionOutput{
				Port: ofp.PortNo(request.OutPort),
			},
		},
		Data: []byte(request.Data),
	}
	return packetOut
}

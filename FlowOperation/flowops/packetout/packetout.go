package packetout

import (
	switchop "FlowOperation/flowops/switch"
	"encoding/base64"
	"encoding/json"
	"github.com/netrack/openflow/ofp"
	"log"
	"net/http"
)

type PacketOutRequest struct {
	SwitchID uint64 `json:"switch_id"`
	InPort   uint32 `json:"in_port"`
	OutPort  uint32 `json:"out_port"`
	Data     string `json:"data"` // Base64 encoded
	BufferID uint32 `json:"buffer_id"`
}

func PacketOut(w http.ResponseWriter, r *http.Request) {
	log.Println("PacketOut Endpoint Hit")
	var request PacketOutRequest

	// Decode incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid packet-out request format", http.StatusBadRequest)
		log.Println("Error decoding request body:", err)
		return
	}

	// Decode base64 data
	data, err := base64.StdEncoding.DecodeString(request.Data)
	if err != nil {
		http.Error(w, "Invalid data format in request (not base64)", http.StatusBadRequest)
		log.Println("Error decoding base64 data:", err)
		return
	}

	// Create PacketOut message
	packetOut := newPacket(request, data)

	// Send PacketOut to switch
	if err := switchop.SendToSwitch(request.SwitchID, &packetOut); err != nil {
		http.Error(w, "Failed to send packet out", http.StatusInternalServerError)
		log.Println("Error sending PacketOut to switch:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("Packet sent out successfully:", request)
}

func newPacket(request PacketOutRequest, data []byte) ofp.PacketOut {
	packetOut := ofp.PacketOut{
		Buffer: request.BufferID,
		InPort: ofp.PortNo(request.InPort),
		Actions: []ofp.Action{
			&ofp.ActionOutput{
				Port: ofp.PortNo(request.OutPort),
			},
		},
		Data: data,
	}
	return packetOut
}

package packetout

import (
	"bytes"
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

	if err := sendPacketToSwitch(&packetOut); err != nil {
		log.Println("Error sending flow add to switch:", err)
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

func sendPacketToSwitch(flowMod *ofp.PacketOut) error {

	data, err := json.Marshal(flowMod)
	if err != nil {
		log.Println("Packetout marshal error:", err)
		return err
	}
	url := "http://127.0.0.1:8094/sendpacketout"
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Println("Forward Packetout error:", err)
		return err
	}
	defer resp.Body.Close()
	log.Printf(" Forwarded packetout to %s, got status=%s\n", url, resp.Status)
	return nil
}

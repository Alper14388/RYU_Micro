package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	of "github.com/netrack/openflow"
	"github.com/netrack/openflow/ofp"
)

type PacketInData struct {
	BufferID   uint32 `json:"buffer_id"`
	Data       string `json:"data"`
	InPort     uint32 `json:"in_port"`
	Reason     uint8  `json:"reason"`
	TotalLen   uint16 `json:"total_len"`
	DatapathID uint64 `json:"dpid"`
}

func sendToMicroservice(packet PacketInData) {
	jsonData, err := json.Marshal(packet)
	if err != nil {
		log.Println("Error marshaling Packet-In to JSON:", err)
		return
	}

	resp, err := http.Post("http://127.0.0.1:8090/packetin", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Packet-In sent to microservice:", resp.Status)
}

func handleHello(rw of.ResponseWriter, r *of.Request) {
	log.Println("Hello message received from switch")
	rw.Write(&of.Header{Type: of.TypeHello}, nil)
}

func handleEchoRequest(rw of.ResponseWriter, r *of.Request) {
	log.Println("Echo message received from switch")
	reply := &of.Header{
		Version:     r.Header.Version,
		Type:        of.TypeEchoReply,
		Length:      r.Header.Length,
		Transaction: r.Header.Transaction,
	}

	if err := rw.Write(reply, nil); err != nil {
		log.Println("Error sending Echo Reply:", err)
		return
	}

	log.Println("Echo Reply sent")
}

func handleFeatureRequest(rw of.ResponseWriter, r *of.Request) {
	log.Println("Feature message received from switch")
}

func handlePacketIn(rw of.ResponseWriter, r *of.Request) {
	var packetIn ofp.PacketIn
	packetIn.ReadFrom(r.Body)
	fmt.Println("data-received-in-packet-in", packetIn.Data, packetIn.Reason, packetIn.Buffer, packetIn.Match.Fields, packetIn.Length, packetIn)

	//log.Printf("Packet-In received: %+v\n", packet)
	//sendToMicroservice(packet)
}

func main() {
	mux := of.NewServeMux()
	mux.HandleFunc(of.TypeMatcher(of.TypeHello), handleHello)
	mux.HandleFunc(of.TypeMatcher(of.TypeEchoRequest), handleEchoRequest)
	mux.HandleFunc(of.TypeMatcher(of.TypeFeaturesRequest), handleFeatureRequest)
	mux.HandleFunc(of.TypeMatcher(of.TypePacketIn), handlePacketIn)

	log.Println("Starting OpenFlow controller on port 6633...")
	if err := of.ListenAndServe(":6633", mux); err != nil {
		log.Fatalf("Error starting OpenFlow controller: %v", err)
	}
}

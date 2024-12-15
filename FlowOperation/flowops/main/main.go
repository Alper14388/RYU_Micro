package main

import (
	"FlowOperation/flowops/flowadd"
	"FlowOperation/flowops/packetout"
	"log"
	"net/http"
)

func main() {
	log.Println("Flow Operation Service running on port 8092...")
	http.HandleFunc("/flowadd", flowadd.AddFlow)
	http.HandleFunc("/packetout", packetout.PacketOut)

	log.Fatal(http.ListenAndServe(":8092", nil))
}

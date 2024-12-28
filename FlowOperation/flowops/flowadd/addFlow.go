package flowadd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/netrack/openflow/ofp"
)

type Request struct {
	SwitchID    uint64 `json:"switch_id"`
	InPort      uint32 `json:"in_port"`
	Src         string `json:"src"`
	Dst         string `json:"dst"`
	OutPort     uint32 `json:"out_port"`
	Priority    uint16 `json:"priority"`
	HardTimeout uint16 `json:"hard_timeout"`
	IdleTimeout uint16 `json:"idle_timeout"`
	BufferID    uint32 `json:"buffer_id"`
}

func AddFlow(w http.ResponseWriter, r *http.Request) {
	log.Println("AddFlow called")
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid flow add request format", http.StatusBadRequest)
		log.Println("Error decoding request body:", err)
		return
	}
	log.Println("New flow add request:", request)

	match := newMatch(request)
	flowMod := newFlowMod(request, match)
	log.Println("FlowMod:", flowMod)

	if err := sendFlowAddToSwitch(flowMod); err != nil {
		log.Println("Error sending flow add to switch:", err)
	}

	w.WriteHeader(http.StatusOK)
	log.Println("Flow added successfully:", request)
}

func uint32ToXMValue(value uint32) ofp.XMValue {
	buf := make([]byte, 4)
	buf[0] = byte(value >> 24)
	buf[1] = byte(value >> 16)
	buf[2] = byte(value >> 8)
	buf[3] = byte(value)
	return ofp.XMValue(buf)
}

func stringToXMValue(value string) ofp.XMValue {
	return ofp.XMValue([]byte(value))
}

func newMatch(request Request) ofp.Match {
	match := ofp.Match{
		Type: ofp.MatchTypeXM,
		Fields: []ofp.XM{
			{
				Class: ofp.XMClassOpenflowBasic,
				Type:  ofp.XMTypeInPort,
				Value: uint32ToXMValue(request.InPort),
			},
			{
				Class: ofp.XMClassOpenflowBasic,
				Type:  ofp.XMTypeEthSrc,
				Value: stringToXMValue(request.Src),
			},
			{
				Class: ofp.XMClassOpenflowBasic,
				Type:  ofp.XMTypeEthDst,
				Value: stringToXMValue(request.Dst),
			},
		},
	}
	return match
}

func newFlowMod(request Request, match ofp.Match) *ofp.FlowMod {
	flowMod := &ofp.FlowMod{
		Buffer:      request.BufferID,
		Command:     ofp.FlowAdd,
		Match:       match,
		IdleTimeout: request.IdleTimeout,
		HardTimeout: request.HardTimeout,
		Priority:    request.Priority,
		Instructions: ofp.Instructions{
			&ofp.InstructionApplyActions{
				Actions: []ofp.Action{
					&ofp.ActionOutput{
						Port:   ofp.PortNo(request.OutPort),
						MaxLen: ofp.ContentLenMax,
					},
				},
			},
		},
	}
	return flowMod
}

func sendFlowAddToSwitch(flowMod *ofp.FlowMod) error {

	data, err := json.Marshal(flowMod)
	if err != nil {
		log.Println("FlowAdd marshal error:", err)
		return err
	}
	log.Println(string(data))
	url := "http://127.0.0.1:8094/sendflowmod"
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		log.Println("Forward FlowAdd error:", err)
		return err
	}
	defer resp.Body.Close()
	log.Printf(" Forwarded FlowAdd to %s, got status=%s\n", url, resp.Status)
	return nil
}

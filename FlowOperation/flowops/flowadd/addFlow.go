package flowadd

import (
	cmpb "Connection_Manager/proto"
	pb "FlowOperation/proto"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/netrack/openflow/ofp"
	"google.golang.org/grpc"
)

func AddFlowGRPC(req *pb.FlowAddRequest) (*pb.FlowAddResponse, error) {
	log.Println("AddFlowGRPC called")
	log.Println("New flow add request:", req)

	match := newMatchFromGRPC(req)
	flowMod := newFlowModFromGRPC(req, match)
	log.Println("FlowMod:", flowMod)

	if err := sendFlowAddToSwitch(flowMod); err != nil {
		log.Println("Error sending flow add to switch:", err)
		return &pb.FlowAddResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	log.Println("Flow added successfully:", req)
	return &pb.FlowAddResponse{
		Success: true,
		Message: "Flow added successfully",
	}, nil
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

func newMatchFromGRPC(request *pb.FlowAddRequest) ofp.Match {
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

func newFlowModFromGRPC(request *pb.FlowAddRequest, match ofp.Match) *ofp.FlowMod {
	flowMod := &ofp.FlowMod{
		Buffer:      request.BufferId,
		Command:     ofp.FlowAdd,
		Match:       match,
		IdleTimeout: uint16(request.IdleTimeout),
		HardTimeout: uint16(request.HardTimeout),
		Priority:    uint16(request.Priority),
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

	conn, err := grpc.Dial("localhost:8094", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to Connection Manager: %v", err)
		return err
	}
	defer conn.Close()

	client := cmpb.NewConnectionManagerClient(conn)
	req := &cmpb.FlowModRequest{
		Data: data,
	}

	resp, err := client.SendFlowMod(context.Background(), req)
	if err != nil {
		log.Printf("Error sending FlowMod: %v", err)
		return err
	}

	if !resp.Success {
		log.Printf("FlowMod failed: %s", resp.Message)
		return fmt.Errorf(resp.Message)
	}

	log.Printf("FlowMod sent successfully")
	return nil
}

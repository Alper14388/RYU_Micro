package flowadd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sdn/FlowOperation/flowops/flowadd/utils"

	pb "sdn/common/proto"

	"github.com/netrack/openflow/ofp"
	"google.golang.org/grpc"
)

func AddFlowGRPC(req *pb.FlowAddRequest) (*pb.FlowAddResponse, error) {
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

func newMatchFromGRPC(request *pb.FlowAddRequest) ofp.Match {
	var fields []ofp.XM
	log.Println("request-eth-frame", request.EthType, request.IPProto)
	if request.EthType == 0x0806 { // ARP
		fields = append(fields, ofp.XM{
			Class: ofp.XMClassOpenflowBasic,
			Type:  ofp.XMTypeEthType,
			Value: utils.Uint16ToXMValue(0x0806),
		})
	}

	if request.EthType == 0x0800 && request.IPProto == 0x01 {
		fields = append(fields,
			ofp.XM{
				Class: ofp.XMClassOpenflowBasic,
				Type:  ofp.XMTypeEthType,
				Value: utils.Uint16ToXMValue(0x0800), // IPv4
			},
			ofp.XM{
				Class: ofp.XMClassOpenflowBasic,
				Type:  ofp.XMTypeIPProto,
				Value: utils.Uint8ToXMValue(0x01), // ICMP
			})
	}

	if request.Src != "" {
		sourceMac := utils.MacStringTo6Byte(request.Src)
		fields = append(fields, ofp.XM{
			Class: ofp.XMClassOpenflowBasic,
			Type:  ofp.XMTypeEthSrc,
			Value: ofp.XMValue(sourceMac[:]),
		})
	}

	if request.Dst != "" {
		destinationMac := utils.MacStringTo6Byte(request.Dst)
		fields = append(fields, ofp.XM{
			Class: ofp.XMClassOpenflowBasic,
			Type:  ofp.XMTypeEthDst,
			Value: ofp.XMValue(destinationMac[:]),
		})
	}

	if request.InPort != 0 {
		fields = append(fields, ofp.XM{
			Class: ofp.XMClassOpenflowBasic,
			Type:  ofp.XMTypeInPort,
			Value: utils.Uint32ToXMValue(request.InPort),
		})
	}

	return ofp.Match{
		Type:   ofp.MatchTypeXM,
		Fields: fields,
	}
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
						Port:   ofp.PortNo(ofp.PortNormal),
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

	conn, err := grpc.Dial("sdn-controller-service.default.svc.cluster.local:8094", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to Connection Manager: %v", err)
		return err
	}
	defer conn.Close()

	client := pb.NewConnectionManagerClient(conn)
	instructions := exportInstruction(flowMod)
	log.Println("Instructions:", instructions)
	req := &pb.FlowModRequest{
		Command:      uint32(flowMod.Command),
		Flags:        uint32(flowMod.Flags),
		TableId:      uint32(flowMod.Table),
		Data:         data,
		Instructions: instructions,
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

func exportInstruction(flowMod *ofp.FlowMod) []*pb.Instruction {
	var instructions []*pb.Instruction
	for _, instr := range flowMod.Instructions {
		switch inst := instr.(type) {
		case *ofp.InstructionApplyActions:
			var actions []*pb.Action
			for _, action := range inst.Actions {
				switch act := action.(type) {
				case *ofp.ActionOutput:
					actions = append(actions, &pb.Action{
						Type:   uint32(ofp.ActionTypeOutput),
						Port:   uint32(act.Port),
						MaxLen: uint32(act.MaxLen),
					})
				default:
					log.Printf("Unsupported action type: %T", act)
				}
			}
			instructions = append(instructions, &pb.Instruction{
				Type:    uint32(ofp.InstructionTypeApplyActions),
				Actions: actions,
			})
		default:
			log.Printf("Unsupported instruction type: %T", instr)
		}
	}
	return instructions
}

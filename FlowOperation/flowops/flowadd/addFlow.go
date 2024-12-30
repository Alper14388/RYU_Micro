package flowadd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	pb "sdn/common/proto"

	"github.com/netrack/openflow/ofp"
	"google.golang.org/grpc"
)

func addBasicFlows(switchId uint64) error {
	log.Printf("Adding basic flows to switch: %d", switchId)

	arpFlow := &pb.FlowAddRequest{
		SwitchId:    switchId,
		Priority:    100,
		HardTimeout: 0,
		IdleTimeout: 0,
		Src:         "",
		Dst:         "",
		InPort:      0,
		OutPort:     uint32(ofp.PortNormal),
	}

	arpMatch := ofp.Match{
		Type: ofp.MatchTypeXM,
		Fields: []ofp.XM{
			{Class: ofp.XMClassOpenflowBasic, Type: ofp.XMTypeEthType, Value: uint16ToXMValue(0x0806)}, // ARP
		},
	}
	arpFlowMod := newFlowModFromGRPC(arpFlow, arpMatch)

	icmpFlow := &pb.FlowAddRequest{
		SwitchId:    switchId,
		Priority:    100,
		HardTimeout: 0,
		IdleTimeout: 0,
		Src:         "",
		Dst:         "",
		InPort:      0,
		OutPort:     uint32(ofp.PortNormal),
	}

	icmpMatch := ofp.Match{
		Type: ofp.MatchTypeXM,
		Fields: []ofp.XM{
			{Class: ofp.XMClassOpenflowBasic, Type: ofp.XMTypeEthType, Value: uint16ToXMValue(0x0800)}, // IPv4
			{Class: ofp.XMClassOpenflowBasic, Type: ofp.XMTypeIPProto, Value: uint8ToXMValue(0x01)},    // ICMP
		},
	}
	icmpFlowMod := newFlowModFromGRPC(icmpFlow, icmpMatch)
	controllerFlow := &pb.FlowAddRequest{
		SwitchId:    switchId,
		Priority:    0,
		HardTimeout: 0,
		IdleTimeout: 0,
		Src:         "",
		Dst:         "",
		InPort:      0,
		OutPort:     0,
	}

	controllerMatch := ofp.Match{
		Type:   ofp.MatchTypeXM,
		Fields: []ofp.XM{}, // Default match
	}
	controllerFlowMod := newFlowModFromGRPC(controllerFlow, controllerMatch)
	controllerFlowMod.Instructions = ofp.Instructions{
		&ofp.InstructionApplyActions{
			Actions: []ofp.Action{
				&ofp.ActionOutput{
					Port:   ofp.PortController,
					MaxLen: ofp.ContentLenMax,
				},
			},
		},
	}
	if err := sendFlowAddToSwitch(arpFlowMod); err != nil {
		return fmt.Errorf("failed to add ARP flow: %w", err)
	}
	if err := sendFlowAddToSwitch(icmpFlowMod); err != nil {
		return fmt.Errorf("failed to add ICMP flow: %w", err)
	}
	if err := sendFlowAddToSwitch(controllerFlowMod); err != nil {
		return fmt.Errorf("failed to add controller flow: %w", err)
	}
	return nil
}
func AddFlowGRPC(req *pb.FlowAddRequest) (*pb.FlowAddResponse, error) {
	log.Println("New flow add request:", req)
	if err := addBasicFlows(req.SwitchId); err != nil {
		log.Println("Error adding basic flows:", err)
		return &pb.FlowAddResponse{
			Success: false,
			Message: "Failed to add basic flows",
		}, err
	}
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

func uint16ToXMValue(value uint16) ofp.XMValue {
	buf := make([]byte, 2)
	buf[0] = byte(value >> 8)
	buf[1] = byte(value)
	return ofp.XMValue(buf)
}
func uint8ToXMValue(value uint8) ofp.XMValue {
	buf := make([]byte, 1)
	buf[0] = byte(value)
	return ofp.XMValue(buf)
}

func macStringTo6Byte(s string) [6]byte {
	hw, err := net.ParseMAC(s)
	if err != nil {
		log.Println("ERROR: parsing MAC:", err)
	}
	var arr [6]byte
	copy(arr[:], hw)
	return arr
}
func newMatchFromGRPC(request *pb.FlowAddRequest) ofp.Match {
	sourceMac := macStringTo6Byte(request.Src)
	destinationMac := macStringTo6Byte(request.Dst)
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
				Value: ofp.XMValue(sourceMac[:]),
			},
			{
				Class: ofp.XMClassOpenflowBasic,
				Type:  ofp.XMTypeEthDst,
				Value: ofp.XMValue(destinationMac[:]),
			},
			{
				Class: ofp.XMClassOpenflowBasic,
				Type:  ofp.XMTypeEthType,
				Value: uint16ToXMValue(0x0800), // IPv4
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

	conn, err := grpc.Dial("localhost:8094", grpc.WithInsecure())
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

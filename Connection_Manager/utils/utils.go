package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/netrack/openflow/ofp"
	"log"
	"net"
	pb "sdn/common/proto"
)

func CalculateFlowModSize(flowMod *ofp.FlowMod) (int, error) {
	var buf bytes.Buffer
	if _, err := flowMod.WriteTo(&buf); err != nil {
		return 0, err
	}
	return buf.Len(), nil
}

func GetInstruction(req *pb.FlowModRequest) (ofp.Instructions, error) {
	var instructions ofp.Instructions
	for _, inst := range req.Instructions {
		switch inst.Type {
		case uint32(ofp.InstructionTypeApplyActions): // ApplyActions
			var actions ofp.Actions
			for _, action := range inst.Actions {
				switch action.Type {
				case uint32(ofp.ActionTypeOutput): // ActionOutput
					actions = append(actions, &ofp.ActionOutput{
						Port:   ofp.PortNo(action.Port),
						MaxLen: uint16(action.MaxLen),
					})
				default:
					log.Printf("Unsupported action type: %v", action.Type)
					return nil, errors.New("unsupported action type")
				}
			}
			instructions = append(instructions, &ofp.InstructionApplyActions{
				Actions: actions,
			})
		default:
			log.Printf("Unsupported instruction type: %v", inst.Type)
			return nil, errors.New("unsupported instruction type")
		}
	}
	return instructions, nil
}

func SendHeaderandBuffer(c net.Conn, flowMod ofp.FlowMod) error {
	header := make([]byte, 8)
	header[0] = 4
	header[1] = 14
	size, _ := CalculateFlowModSize(&flowMod)
	binary.BigEndian.PutUint16(header[2:], uint16(8+size))
	binary.BigEndian.PutUint32(header[4:], uint32(12345))

	var buf bytes.Buffer
	if _, err := flowMod.WriteTo(&buf); err != nil {
		return errors.New(fmt.Sprintf("Error sending FlowMod: %v", err))
	}

	if _, err := c.Write(append(header, buf.Bytes()...)); err != nil {
		return errors.New(fmt.Sprintf("Error sending FlowMod: %v", err))
	}
	return nil
}

package packetout

import (
	"context"
	"fmt"
	"log"
	pb "sdn/common/proto"

	"github.com/netrack/openflow/ofp"
	"google.golang.org/grpc"
)

func PacketOutGRPC(req *pb.PacketOutRequest) (*pb.PacketOutResponse, error) {
	log.Println("PacketOut Endpoint Hit")

	// Create PacketOut message
	packetOut := newPacketFromGRPC(req, req.Data)

	if err := sendPacketToSwitch(&packetOut); err != nil {
		log.Println("Error sending packet out to switch:", err)
		return &pb.PacketOutResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	log.Println("Packet sent out successfully:", req)
	return &pb.PacketOutResponse{
		Success: true,
		Message: "Packet sent successfully",
	}, nil
}

func newPacketFromGRPC(request *pb.PacketOutRequest, data []byte) ofp.PacketOut {
	packetOut := ofp.PacketOut{
		Buffer: request.BufferId,
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

func sendPacketToSwitch(packetOut *ofp.PacketOut) error {

	var actions []*pb.Action
	for _, action := range packetOut.Actions {
		protoAction, err := ofpActionToProto(action)
		if err != nil {
			return err
		}
		actions = append(actions, protoAction)
	}

	conn, err := grpc.Dial("localhost:8094", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to Connection Manager: %v", err)
		return err
	}
	defer conn.Close()
	log.Printf("Sending packet out to switch: %v", string(packetOut.Data))
	client := pb.NewConnectionManagerClient(conn)
	req := &pb.PacketOutRequest{
		Data:     packetOut.Data,
		InPort:   uint32(packetOut.InPort),
		Actions:  actions,
		BufferId: packetOut.Buffer,
	}

	resp, err := client.SendPacketOut(context.Background(), req)
	if err != nil {
		log.Printf("Error sending PacketOut: %v", err)
		return err
	}

	if !resp.Success {
		log.Printf("PacketOut failed: %s", resp.Message)
		return fmt.Errorf(resp.Message)
	}

	log.Printf("PacketOut sent successfully")
	return nil
}

func ofpActionToProto(action ofp.Action) (*pb.Action, error) {
	switch act := action.(type) {
	case *ofp.ActionOutput:
		return &pb.Action{
			Type:   uint32(ofp.ActionTypeOutput),
			Port:   uint32(act.Port),
			MaxLen: uint32(act.MaxLen),
		}, nil
	default:
		return nil, fmt.Errorf("unsupported action type: %v", action.Type())
	}
}

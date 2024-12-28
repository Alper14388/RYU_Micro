package packetout

import (
	"context"
	"encoding/json"
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
	data, err := json.Marshal(packetOut)
	if err != nil {
		log.Println("PacketOut marshal error:", err)
		return err
	}

	conn, err := grpc.Dial("localhost:8094", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to Connection Manager: %v", err)
		return err
	}
	defer conn.Close()

	client := pb.NewConnectionManagerClient(conn)
	req := &pb.PacketOutRequest{
		Data: data,
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

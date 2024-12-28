package main

import (
	"context"
	"log"
	"net"
	"sdn/FlowOperation/flowops/flowadd"
	"sdn/FlowOperation/flowops/packetout"
	pb "sdn/common/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFlowOperationServer
}

func (s *server) AddFlow(ctx context.Context, req *pb.FlowAddRequest) (*pb.FlowAddResponse, error) {
	return flowadd.AddFlowGRPC(req)
}

func (s *server) SendPacketOut(ctx context.Context, req *pb.PacketOutRequest) (*pb.PacketOutResponse, error) {
	return packetout.PacketOutGRPC(req)
}

func main() {
	log.Println("Flow Operation Service running on port 8092...")

	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFlowOperationServer(s, &server{})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package Main

import (
	"context"
	"log"
	"net"
	ofp "sdn/Ryu_go/packethandler"
	pb "sdn/common/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedPacketHandlerServer
}

// HandlePacketIn handles incoming PacketIn gRPC requests
func (s *server) HandlePacketIn(ctx context.Context, req *pb.PacketInRequest) (*pb.PacketInResponse, error) {
	log.Printf("Received PacketIn requests: %+v", req)
	resp, err := ofp.HandlePacketIn(req)
	if err != nil {
		log.Printf("Error handling PacketIn: %v", err)
		return nil, err
	}
	log.Printf("PacketIn response: %+v", resp)
	return resp, nil
}

func main() {
	address := ":8090"
	log.Printf("Starting gRPC server on %s...", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	grpcServer := grpc.NewServer()

	// Register PacketHandler service
	pb.RegisterPacketHandlerServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	// Start serving
	log.Printf("gRPC server is now listening at %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

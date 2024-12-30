package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"sdn/Connection_Manager/Server"
	pb "sdn/common/proto"
)

func main() {
	go Server.ListenAndServeOpenFlow(":6633")

	lis, err := net.Listen("tcp", ":8094")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	store := Server.GetStore()
	pb.RegisterConnectionManagerServer(s, &Server.Server{Store: store})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

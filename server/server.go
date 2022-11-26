package main

import (
	"log"
	"net"
	"uacademy/grpc_example/proto-gen/dice"
	dService "uacademy/grpc_example/server/services/dice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	c := &dService.TutorialService{}
	s := grpc.NewServer()
	dice.RegisterTutorialServer(s, c)
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

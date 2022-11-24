package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"uacademy/grpc_example/proto-gen/dice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// TutorialService is a struct that implements the server interface
type TutorialService struct {
	dice.UnimplementedTutorialServer
}

// RollDice ...
func (s *TutorialService) RollDice(ctx context.Context, req *dice.RollDiceRequest) (*dice.RollDiceResponse, error) {
	var res dice.RollDiceResponse

	if req.Num < 0 {
		return nil, grpc.Errorf(codes.InvalidArgument, "number should be postive")
	}

	for i := 0; i < int(req.Num); i++ {
		res.Dice = append(res.Dice, int32(rand.Intn(100)))
	}

	return &res, nil
}

func (s *TutorialService) HelloMethod(ctx context.Context, req *dice.Empty) (*dice.Hello, error) {
	return &dice.Hello{
		Foo: "hi",
	}, nil
}

func main() {
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	dice.RegisterTutorialServer(s, &TutorialService{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

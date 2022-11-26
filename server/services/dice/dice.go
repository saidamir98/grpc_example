package dice

import (
	"context"
	"log"
	"math/rand"
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

// Ping ...
func (s *TutorialService) Ping(ctx context.Context, req *dice.Empty) (*dice.Pong, error) {
	log.Println("Ping")

	return &dice.Pong{
		Message: "OK",
	}, nil
}

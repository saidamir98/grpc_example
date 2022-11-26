package main

import (
	"context"
	"log"
	"time"

	"uacademy/grpc_example/proto-gen/dice"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := dice.NewTutorialClient(conn)

	for i := 0; i < 10; i++ {
		res, err := c.Ping(context.Background(), &dice.Empty{})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", res.Message)
		time.Sleep(time.Second * 1)
	}

}

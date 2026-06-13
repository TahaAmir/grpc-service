package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/TahaAmir/grpc-service/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("%v", res.Message)
}

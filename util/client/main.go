package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"and-gate/v2/pb"
)

func main() {
	conn, err := grpc.NewClient("localhost:5555",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStartServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetStartNanos(ctx, &pb.GetStartNanosRequest{})
	if err != nil {
		log.Fatalf("GetStartNanos failed: %v", err)
	}

	log.Printf("GetStartNanos response: StartNanos=%d", resp.GetStartNanos())
}

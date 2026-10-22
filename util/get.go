package util

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"and-gate/v2/pb"
)

func GetStartTime(address *string) {
	conn, err := grpc.NewClient(*address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close connection: %v", err)
		}
	}()

	client := pb.NewStartServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetStartNanos(ctx, &pb.GetStartNanosRequest{})
	if err != nil {
		log.Fatalf("GetStartNanos failed: %v", err)
	}

	log.Printf("GetStartNanos response: StartNanos=%d", resp.GetStartNanos())
}

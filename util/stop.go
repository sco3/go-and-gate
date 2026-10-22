package util

import (
	"and-gate/v2/pb"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Stop(address *string) {

	conn, err := grpc.NewClient(*address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStartServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.StopServer(ctx, &pb.StopServerRequest{})
	if err != nil {
		log.Fatalf("Stop server failed: %v", err)
	}

	log.Printf("Stop server response: %v", resp)

}

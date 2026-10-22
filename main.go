package main

import (
	"and-gate/v2/srv"
	"flag"
	"fmt"
	"log"
)

func main() {

	address := flag.String("address", ":5555", "Address for gRPC server to listen on")
	flag.Parse()

	fmt.Printf("Address: %v\n", *address)

	server, err := srv.NewServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

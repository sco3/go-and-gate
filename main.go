package main

import (
	"and-gate/v2/srv"
	"log"
)

func main() {
	server, err := srv.NewServer(":5555")
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}

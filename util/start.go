package util

import (
	"and-gate/v2/srv"
	"log"
)

func Start(address *string, host *string) {

	server, err := srv.NewServer(*address, *host)
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}

}

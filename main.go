package main

import (
	"and-gate/v2/util"
	"flag"
)

func main() {
	startFlag := flag.Bool("start", false, "Start the gRPC server")
	stopFlag := flag.Bool("stop", false, "Stop the running gRPC server")
	getTimeFlag := flag.Bool("get", false, "Get start time")

	address := flag.String("address", ":5555", "Address for gRPC server to listen on")

	host := flag.String("host", "", "Host/Pod name for redis distributed barrier")

	flag.Parse()

	if *startFlag {
		util.PrintParams()
		util.Start(address, host)
	} else if *stopFlag {
		util.PrintParams()
		util.Stop(address)
	} else if *getTimeFlag {
		util.PrintParams()
		util.GetStartTime(address)
	} else {
		flag.Usage()
	}

}

package main

import (
	"flag"
	"net"
	"net/rpc"

	"uk.ac.bris.cs/distributed3/pairbroker/stubs"
)

type Factory struct{}

//TODO: Define a Multiply function to be accessed via RPC.
//Check the previous weeks' examples to figure out how to do this.

func Multiply(pair stubs.Pair) stubs.JobReport {
	out := pair.X * pair.Y
	return stubs.JobReport{Result: out}
}

func main() {
	pAddr := flag.String("ip", "127.0.0.1:8050", "IP and port to listen on")
	brokerAddr := flag.String("broker", "127.0.0.1:8030", "Address of broker instance")
	flag.Parse()
	listener, _ := net.Listen("tcp", *pAddr)
	conn, _ := net.Dial("tcp", *brokerAddr)
	defer listener.Close()
	rpc.Accept(listener)
	defer conn.Close()

	//TODO: You'll need to set up the RPC server, and subscribe to the running broker instance.
}

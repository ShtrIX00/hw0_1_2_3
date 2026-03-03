package main

import (
	"io"
	"log"
	"net"
	"os"
)

const (
	network          = "tcp"
	address          = "localhost:8080"
	expectedResponse = "OK\n"
)

func main() {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	data, err := io.ReadAll(conn)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if string(data) != expectedResponse {
		os.Exit(1)
	}
}

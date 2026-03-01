package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConnection(conn) // можно добавить конкурентность (go), а можно и нет
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte(expectedResponse))
	if err != nil {
		log.Println(err)
	}
}

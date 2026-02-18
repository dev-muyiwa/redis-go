package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening to TCP server: %s\n", err)
		return
	}
	defer listener.Close()

	log.Println("Listening to TCP server on port 8080")

	pool := NewWorkerPool(20)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %s", err)
			continue
		}

		pool.tasks <- conn
	}
}

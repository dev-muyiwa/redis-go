package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening to TCP server: %s\n", err)
		return
	}
	defer listener.Close()

	log.Println("Listening to TCP server on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %s", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatalf("Error reading: %v\n", err)
			return
		}
		log.Printf("Received %s\n", buffer[:n])
		conn.Write([]byte("Message received\n"))
	}
}

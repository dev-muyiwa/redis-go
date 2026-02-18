package main

import (
	"log"
	"net"
)

type WorkerPool struct {
	tasks chan net.Conn
}

func NewWorkerPool(size int) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan net.Conn, 100),
	}
	for i := 0; i < size; i++ {
		go pool.worker()
	}

	return pool
}

func (p *WorkerPool) worker() {
	for conn := range p.tasks {
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	//conn.SetReadDeadline(time.Now().Add(10 * time.Second))
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

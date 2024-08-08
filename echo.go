package main

import (
	"io"
	"log"
	"net"
)

func handleEcho(c net.Conn) {
	defer log.Print("Connection Closed")
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("READ: ", err)
		}
		_, err = c.Write(buf[:n])
		if err != nil {
			log.Fatal("WRITE: ", err)
		}
	}
}

func EchoServer() {
	listener, err := net.Listen("tcp4", ":7007")
	if err != nil {
		log.Fatal("LISTENER: ", err)
	}
	log.Printf("Listening on port: %v", listener.Addr().String())
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("CONN: ", err)
		}
		defer conn.Close()

		go handleEcho(conn)
	}
}

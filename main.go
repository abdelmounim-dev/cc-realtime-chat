package main

import (
	"io"
	"log"
	"net"
)

func main() {
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

		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1024)
			for {
				n, err := c.Read(buf)
				if err != nil {
					if err == io.EOF {
						log.Print("EOF")
						break
					}
					log.Fatal("READ: ", err)
				}
				_, err = c.Write(buf[:n])
				if err != nil {
					log.Fatal("WRITE: ", err)
				}
			}
		}(conn)
	}
}

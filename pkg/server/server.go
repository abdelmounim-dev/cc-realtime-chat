package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func broadcast(conns *[]net.Conn, msgChan chan string) error {
	for {
		msg := <-msgChan
		for _, c := range *conns {
			_, err := c.Write([]byte(msg))
			if err != nil {
				return fmt.Errorf("BROADCAST: %v", err)
			}
		}
	}
}

func handleConn(c net.Conn, msgChan chan string) {
	defer func() {
		log.Print("Connection Closed")
		c.Close() // Close the connection when the function ends
	}()

	reader := bufio.NewReader(c)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("READ: ", err)
			break // Close the connection if there's an error
		}
		log.Println("msg: ", msg)
		msgChan <- msg
	}
}

func establishConnections(listener net.Listener, connChan chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("CONN: ", err)
			continue
		}
		log.Println("Connection Established: ", conn.RemoteAddr())

		connChan <- conn

	}
}

func Server(done chan<- bool, maxConns int) {
	listener, err := net.Listen("tcp4", ":7007")
	if err != nil {
		log.Fatal("LISTENER: ", err)
	}
	log.Printf("Listening on port: %v", listener.Addr().String())
	defer listener.Close()

	connChan := make(chan net.Conn)
	msgChan := make(chan string)
	conns := make([]net.Conn, maxConns)
	conns = conns[:0]
	log.Println(len(conns))

	go establishConnections(listener, connChan)
	go broadcast(&conns, msgChan)

	for {
		if len(conns) == maxConns {
			log.Println("Maximum Connection Number reached: ", maxConns)
			break
		}
		conn := <-connChan
		conns = append(conns, conn)
		go handleConn(conn, msgChan)
	}

	done <- true
}

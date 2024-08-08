package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func handleEcho(c net.Conn) {
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
		_, err = c.Write([]byte(msg))
		if err != nil {
			log.Println("WRITE: ", err)
			break // Close the connection if there's an error
		}
	}
}

func EchoServer(listnerChan chan<- net.Listener) {
	listener, err := net.Listen("tcp4", ":7007")
	if err != nil {
		log.Fatal("LISTENER: ", err)
	}
	log.Printf("Listening on port: %v", listener.Addr().String())
	defer listener.Close()

	listnerChan <- listener

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("CONN: ", err)
			continue
		}
		// Start a goroutine for each connection, don't defer conn.Close() here
		go handleEcho(conn)
	}
}

func getMessage() string {
	fmt.Print("send: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func printMessage(msg string) {
	fmt.Println("recv:", msg)
}

func EchoClient(listener net.Listener, done chan<- bool) {
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		log.Println("DIAL: ", err)
		return
	}

	defer conn.Close()
	reader := bufio.NewReader(conn)

	for msg := getMessage(); msg != "quit"; msg = getMessage() {
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			log.Println("WRITE: ", err)
			continue
		}
		recv, err := reader.ReadString('\n')
		if err != nil {
			log.Println("READ: ", err)
			continue
		}
		printMessage(strings.TrimSpace(recv))
	}
	log.Println("QUIT")
	done <- true
}

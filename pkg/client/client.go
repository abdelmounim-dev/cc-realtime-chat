package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func getMessage(prompt string) string {
	if prompt != "" {
		fmt.Print(prompt + ": ")
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func printMessage(prompt, msg string) {
	fmt.Println(prompt+":", msg)
}

func isTerminationMessage(msg string) bool {
	return msg == "quit" || msg == "exit"
}

func send(conn net.Conn, done chan<- bool) {
	name := getMessage("name")
	for {
		msg := getMessage("")
		if isTerminationMessage(msg) {
			done <- true
			break
		}
		_, err := conn.Write([]byte(name + ": " + msg + "\n"))
		if err != nil {
			log.Println("WRITE: ", err)
			continue
		}
	}
	log.Println("QUIT")
}

func receive(conn net.Conn, done <-chan bool) {
	reader := bufio.NewReader(conn)
	for {
		select {
		case <-done:
			log.Println("Receive routine stopped.")
			return
		default:
			recv, err := reader.ReadString('\n')
			if err != nil {
				log.Println("READ: ", err)
				continue
			}
			fmt.Println(recv)
		}
	}
}

func Client(addr string, done chan<- bool) {
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		log.Println("DIAL: ", err)
		done <- true
		return
	}
	defer conn.Close()

	clientDone := make(chan bool)
	go send(conn, clientDone)
	go receive(conn, clientDone)

	<-clientDone
	done <- true
}

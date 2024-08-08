package main

import "net"

func main() {
	listenerChan := make(chan net.Listener)
	done := make(chan bool)

	go EchoServer(listenerChan)
	listener := <-listenerChan

	go EchoClient(listener, done)
	<-done
}

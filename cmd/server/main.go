package main

import (
	"github.com/abdelmounim-dev/cc-realtime-chat/pkg/server"
)

func main() {
	done := make(chan bool)

	go server.Server(done, 3)
	<-done

}

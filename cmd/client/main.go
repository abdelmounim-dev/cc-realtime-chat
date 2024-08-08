package main

import (
	"github.com/abdelmounim-dev/cc-realtime-chat/pkg/client"
)

func main() {
	done := make(chan bool)

	go client.Client("localhost:7007", done)
	<-done
}

package server

import (
	"bufio"
	"io"
	"log"
	"net"
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
		log.Println("msg: ", msg)
		_, err = c.Write([]byte(msg))
		if err != nil {
			log.Println("WRITE: ", err)
			break // Close the connection if there's an error
		}
	}
}

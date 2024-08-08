package server

import (
	"net"
	"sync"
)

type Connections struct {
	mu    sync.Mutex
	conns []net.Conn
}

func (c *Connections) append(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.conns = append(c.conns, conn)
}

func (c *Connections) remove(index int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.conns = append(c.conns[:index], c.conns[index+1:]...)
}

package udt

import (
	"fmt"
	"net"
	"time"

	"github.com/murphybytes/udt.go/cudt"
)

// Conn implements net.Conn
type Conn struct {
	connectionKey int
	listener      *Listener
	remoteAddress string
}

func (c *Conn) Read(b []byte) (n int, e error) {
	n, e = cudt.Read(c.connectionKey, b)
	fmt.Printf("READ: %s\n", string(b))
	if e != nil {
		return 0, e
	}
	return n, nil
}

func (c *Conn) Write(b []byte) (n int, e error) {
	n, e = cudt.Write(c.connectionKey, b)
	if e != nil {
		return 0, e
	}
	return
}

func (c *Conn) LocalAddr() (a net.Addr) {
	return a
}

func (c *Conn) RemoteAddr() (a net.Addr) {
	return a
}

func (c *Conn) SetDeadline(t time.Time) (e error) {
	return e
}

func (c *Conn) SetReadDeadline(t time.Time) (e error) {
	return e
}

func (c *Conn) SetWriteDeadline(t time.Time) (e error) {
	return e
}

func (c *Conn) Close() (e error) {
	cudt.Close(c.connectionKey)
	return e
}

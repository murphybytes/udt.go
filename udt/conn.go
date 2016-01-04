package udt

import (
	"net"
	"time"
	"unsafe"
)

// Conn implements net.Conn
type Conn struct {
	udtPointer unsafe.Pointer
}

func (c *Conn) Read(b []byte) (n int, e error) {
	return n, e
}

func (c *Conn) Write(b []byte) (n int, e error) {
	return n, e
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

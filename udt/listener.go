package udt

import (
	"fmt"
	"net"
	"strings"
	"unsafe"

	"github.com/murphybytes/udt.go/cudt"
)

// Listener implments net.Listener
type Listener struct {
	udtPointer unsafe.Pointer
}

func Listen(addressString string) (l *Listener, e error) {
	parts := strings.Split(addressString, ":")
	if len(parts) > 2 || len(parts) < 1 {
		e = fmt.Errorf("Address string invalid %s", addressString)
		return
	}

	var ipaddr string
	var port string

	if len(parts) < 2 {
		port = parts[0]
	} else {
		ipaddr = parts[0]
		port = parts[1]
	}

	l, e = cudt.Listen(ipaddr, port)
	return
}

func (l *Listener) Accept() (c Conn, e error) {
	return
}

func (l *Listener) Close() (c Conn, e error) {
	return
}

func (l *Listener) Addr() (a net.Addr) {
	return
}

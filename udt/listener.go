package udt

import (
	"fmt"
	"net"

	"github.com/murphybytes/udt.go/cudt"
)

// Listener implments net.Listener
type Listener struct {
	sessionKey int
	address    string
}

func Listen(addressString string) (l net.Listener, e error) {
	var host, port string
	host, port, e = net.SplitHostPort(addressString)

	if e != nil {
		return
	}

	var sessionKey int
	sessionKey, e = cudt.Listen(host, port)
	fmt.Printf("listen returned key %d\n", sessionKey)
	if e == nil {
		l = &Listener{
			sessionKey: sessionKey,
			address:    addressString,
		}
	}
	return
}

func (l *Listener) Accept() (c net.Conn, e error) {
	var connectionKey int
	var addr string
	connectionKey, addr, e = cudt.Accept(l.sessionKey)
	c = &Conn{
		connectionKey: connectionKey,
		remoteAddress: addr,
		listener:      l,
	}

	return
}

func (l *Listener) Close() (e error) {
	cudt.Close(l.sessionKey)
	return
}

func (l *Listener) Addr() (a net.Addr) {
	return
}

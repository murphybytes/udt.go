package udt

import (
	"fmt"
	"net"
	"strings"

	"github.com/murphybytes/udt.go/cudt"
)

// Listener implments net.Listener
type Listener struct {
	sessionKey int
	address    string
}

func Listen(addressString string) (l net.Listener, e error) {
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

	var sessionKey int
	sessionKey, e = cudt.Listen(ipaddr, port)
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

package cudt

// #cgo CXXFLAGS: -I${SRCDIR}/../vendor/udt/udt4/src
// #cgo LDFLAGS: -L${SRCDIR}/../vendor/udt/udt4/src -ludt -lstdc++ -lpthread -lm
// #include "stdlib.h"
// #include "cudt.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type action int

const (
	put action = iota
	get
	del
	kill
)

type request struct {
	action       action
	key          int
	handle       int
	responseChan chan request
	ok           bool
}

var requestChan chan request

func getUDTHandle(key int) (sock int, e error) {
	req := request{
		action:       get,
		key:          key,
		responseChan: make(chan request),
	}
	requestChan <- req
	resp := <-req.responseChan
	close(req.responseChan)

	if !resp.ok {
		e = fmt.Errorf("No session found for key %d", key)
	} else {
		sock = resp.handle
	}

	return
}

func saveUDTHandle(sock int) int {
	req := request{
		action:       put,
		handle:       sock,
		responseChan: make(chan request),
	}
	requestChan <- req
	resp := <-req.responseChan
	close(req.responseChan)
	return resp.key
}

func deleteUDTHandle(key int) {
	req := request{
		action:       del,
		key:          key,
		responseChan: make(chan request),
	}
	requestChan <- req
	<-req.responseChan
	close(req.responseChan)
}

// Startup initializes UDT library and creates session management infrastructure
// used to manage handles into the UDT library in a 'thread safe' manner. Note
// you must call Cleanup before calling Startup more than one time.
func Startup() (e error) {
	if requestChan != nil {
		return errors.New("Illegal Action. Startup can't be called twice unless Cleanup is called.")
	}
	requestChan = make(chan request)

	var udtResponse C.int
	udtResponse, e = C.startup()
	if e != nil {
		return e
	}
	if udtResponse != 0 {
		e = fmt.Errorf("UDT Startup failed with error %d", udtResponse)
	}

	sessions := map[int]int{}
	var nextKey int

	go func() {
		for {
			req := <-requestChan
			switch req.action {
			case put:
				req.key = nextKey
				nextKey++
				sessions[req.key] = req.handle
				req.responseChan <- req
			case get:
				req.handle, req.ok = sessions[req.key]
				req.responseChan <- req
			case del:
				delete(sessions, req.key)
				req.responseChan <- req
			case kill:
				close(requestChan)
				req.responseChan <- req
				return
			}
		}
	}()

	return nil
}

// Cleanup frees up resources allocated by Startup
func Cleanup() (e error) {
	req := request{
		action:       kill,
		responseChan: make(chan request),
	}
	requestChan <- req
	<-req.responseChan
	close(req.responseChan)
	requestChan = nil

	if i, e := C.cleanup(); e == nil {

		if i != 0 {
			e = fmt.Errorf("UDT Cleanup failed")
		}

	}
	return e
}

// Close frees up UDT socket
func Close(sessionKey int) {
	handle, err := getUDTHandle(sessionKey)
	if err != nil {
		deleteUDTHandle(sessionKey)
		C.udt_close(C.int(handle))
	}

}

// Listen opens a UDT socket for listening
func Listen(ipaddr string, port string) (sessionKey int, e error) {
	var result *C.struct_udt_result
	cipaddr := C.CString(ipaddr)
	defer C.free(unsafe.Pointer(cipaddr))
	cport := C.CString(port)
	defer C.free(unsafe.Pointer(cport))

	C.udt_listen(cipaddr, cport, &result)

	defer C.free(unsafe.Pointer(result))

	if result.errorMsg != nil {
		e = errors.New(C.GoString(result.errorMsg))
		C.free(unsafe.Pointer(result.errorMsg))
		return
	}

	sessionKey = saveUDTHandle(int(result.udtSocket))

	return

}

// Dial calls handles interface between Go and C. It returns a
// clientKey which is used to reference the connection socket
// returned from UDT C runtime.
func Dial(ipaddr string, port string) (clientKey int, e error) {
	var result *C.struct_udt_result
	cipaddr := C.CString(ipaddr)
	defer C.free(unsafe.Pointer(cipaddr))
	cport := C.CString(port)
	defer C.free(unsafe.Pointer(cport))

	C.udt_connect(cipaddr, cport, &result)
	defer C.free(unsafe.Pointer(result))

	if result.errorMsg != nil {
		e = errors.New(C.GoString(result.errorMsg))
		C.free(unsafe.Pointer(result.errorMsg))
		return
	}

	clientKey = saveUDTHandle(int(result.udtSocket))

	return
}

// Accept calls through to C UDT accept and takes care of freeing
// C allocated memory and copying data to Go managed memory
func Accept(serverKey int) (connectionKey int, addr string, e error) {

	var result *C.struct_udt_result
	var serverHnd int
	serverHnd, e = getUDTHandle(serverKey)
	if e != nil {
		return
	}

	C.udt_accept(C.int(serverHnd), &result)
	defer C.free(unsafe.Pointer(result))

	if result.errorMsg != nil {
		e = errors.New(C.GoString(result.errorMsg))
		C.free(unsafe.Pointer(result.errorMsg))
		return
	}

	addr = C.GoString(result.addrString)
	C.free(unsafe.Pointer(result.addrString))

	connectionKey = saveUDTHandle(int(result.udtSocket))

	return

}

func Read(connectionKey int, buffer []byte) (read int, e error) {
	var result *C.struct_udt_result

	var bytes_read C.int
	var connectionHnd int
	connectionHnd, e = getUDTHandle(connectionKey)

	if e != nil {
		return
	}

	C.udt_recv(C.int(connectionHnd), (*C.char)(unsafe.Pointer(&buffer[0])), C.int(len(buffer)), &bytes_read, &result)
	defer C.free(unsafe.Pointer(result))

	if result.errorMsg != nil {
		e = errors.New(C.GoString(result.errorMsg))
		C.free(unsafe.Pointer(result.errorMsg))
		return
	}

	read = int(bytes_read)

	fmt.Printf("Response '%s' bytes read %d\n", string(buffer), read)

	return read, e
}

// Write sends contents of buffer to network reciever
func Write(connectionKey int, buffer []byte) (written int, e error) {
	var result *C.struct_udt_result
	var connectionHnd int
	connectionHnd, e = getUDTHandle(connectionKey)

	if e != nil {
		return
	}

	len := C.int(len(buffer))
	var cwritten C.int

	C.udt_send(C.int(connectionHnd), (*C.char)(unsafe.Pointer(&buffer[0])), len, &cwritten, &result)
	defer C.free(unsafe.Pointer(result))

	if result.errorMsg != nil {
		e = errors.New(C.GoString(result.errorMsg))
		C.free(unsafe.Pointer(result.errorMsg))
	}

	return int(cwritten), e

}

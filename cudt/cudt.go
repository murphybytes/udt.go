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
	handle       unsafe.Pointer
	responseChan chan request
	ok           bool
}

var requestChan chan request

func getUDTHandle(key int) (p unsafe.Pointer, e error) {
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
		p = resp.handle
	}

	return
}

func saveUDTHandle(p unsafe.Pointer) int {
	req := request{
		action:       put,
		handle:       p,
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

	sessions := map[int]unsafe.Pointer{}
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
		C.udt_close(handle)
	}

}

// Listen opens a UDT socket for listening
func Listen(ipaddr string, port string) (sessionKey int, e error) {
	var result *C.struct_udt_result
	C.udt_listen(C.CString(ipaddr), C.CString(port), &result)
	defer C.free(unsafe.Pointer(result))

	if result.errorMsg != nil {
		e = errors.New(C.GoString(result.errorMsg))
		C.free(unsafe.Pointer(result.errorMsg))
		return
	}

	sessionKey = saveUDTHandle(result.udtPointer)
	return

}

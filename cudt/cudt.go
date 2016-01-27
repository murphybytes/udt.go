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

func Startup() (e error) {

	if i, e := C.startup(); e == nil {

		if i != 0 {
			e = fmt.Errorf("UDT Startup failed")
		}

	}
	return e
}

func Cleanup() (e error) {

	if i, e := C.cleanup(); e == nil {

		if i != 0 {
			e = fmt.Errorf("UDT Cleanup failed")
		}

	}
	return e
}

func Listen(ipaddr string, port string) (udtPointer unsafe.Pointer, e error) {
	var result *C.struct_udt_result
	C.udt_listen(C.CString(ipaddr), C.CString(port), &result)
	fmt.Println("got here")
	if result.errorMsg != nil {
		e = errors.New(C.GoString(result.errorMsg))
		C.free(unsafe.Pointer(result.errorMsg))

	}

	udtPointer = result.udtPointer
	fmt.Println("end")
	return

}

package cudt

// #cgo CXXFLAGS: -I${SRCDIR}/../vendor/udt/udt4/src
// #cgo LDFLAGS: -L${SRCDIR}/../vendor/udt/udt4/src -ludt -lstdc++ -lpthread -lm
// #include "cudt.h"
import "C"
import "fmt"

func startup() (e error) {

	if i, e := C.startup(); e == nil {

		if i != 0 {
			e = fmt.Errorf("UDT Startup failed")
		}

	}
	return e
}

func cleanup() (e error) {

	if i, e := C.cleanup(); e == nil {

		if i != 0 {
			e = fmt.Errorf("UDT Cleanup failed")
		}

	}
	return e
}

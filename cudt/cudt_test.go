package cudt

import (
	"fmt"
	"testing"
)

func TestInitializationAndShutdown(t *testing.T) {
	e := Startup()
	if e != nil {
		t.Error(e)
	}

	e = Cleanup()
	if e != nil {
		t.Error(e)
	}

}

func TestListen(t *testing.T) {
	e := Startup()
	if e != nil {
		t.Error(e)
	} else {
		defer Cleanup()
	}

	_, e = Listen("127.0.0.1", "9000")
	// fmt.Printf("Result -> %s\n", e.Error())
	fmt.Println("yep")
}

package cudt

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestInitializationAndShutdown(t *testing.T) {
	e := Startup()
	if e != nil {
		t.Fatal(e)
	}

	e = Cleanup()
	if e != nil {
		t.Fatal(e)
	}

}

func TestSessionManagment(t *testing.T) {
	e := Startup()
	if e != nil {
		t.Fatalf("Startup failed with %s", e)
	} else {
		defer Cleanup()
	}

	type dummy struct {
		d string
	}
	p1 := unsafe.Pointer(&dummy{d: "something"})
	fmt.Printf("P1 %q\n", p1)
	k1 := saveUDTHandle(p1)
	p2, err := getUDTHandle(k1)
	if err != nil {
		t.Fatal(err)
	}
	if p1 != p2 {
		t.Fatal("Pointers don't match")
	}
	wrongKey := k1 + 1
	p2, err = getUDTHandle(wrongKey)
	if err == nil {
		t.Fatal("Expected error, didn't get one")
	}
	if err.Error() != fmt.Sprintf("No session found for key %d", wrongKey) {
		t.Fatal("Error message not expected")
	}
	p3 := unsafe.Pointer(&dummy{d: "other"})
	fmt.Printf("P3 %q\n", p3)
	k2 := saveUDTHandle(p3)
	if k2 == k1 {
		t.Fatal("Keys should not be the same")
	}
	p4, _ := getUDTHandle(k2)
	if p4 != p3 {
		t.Fatal("Pointers should match")
	}
	if p4 == p1 {
		t.Fatal("Pointers shouldn't match")
	}

	deleteUDTHandle(k2)
	_, err = getUDTHandle(k2)
	if err == nil {
		t.Fatal("Expected error here because key should not be present")
	}

}

func TestListen(t *testing.T) {
	e := Startup()
	if e != nil {
		t.Error(e)
	} else {
		defer Cleanup()
	}

	key, e := Listen("127.0.0.1", "9000")
	if e != nil {
		t.Fatalf("Listen failed: %s", e.Error())
	}
	Close(key)

}

package cudt

import "testing"

func TestInitializationAndShutdown(t *testing.T) {
	e := startup()
	if e != nil {
		t.Error(e)
	}

	e = cleanup()
	if e != nil {
		t.Error(e)
	}
}

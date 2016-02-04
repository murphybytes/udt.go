package udt

import (
	"github.com/murphybytes/udt.go/cudt"
)

// Startup sets up resources used by UDT.
func Startup() error {
	return cudt.Startup()
}

// Cleanup frees resources used by UDT.
func Cleanup() error {
	return cudt.Cleanup()
}

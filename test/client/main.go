package main

import (
	"fmt"
	"os"

	"github.com/murphybytes/udt.go/cudt"
	"github.com/murphybytes/udt.go/udt"
)

func main() {
	cudt.Startup()
	defer cudt.Cleanup()
	fmt.Println("Starting client test.")
	conn, e := udt.Dial("127.0.0.1:9876")

	if e != nil {
		fmt.Printf("Dial fails with %s\n", e.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Dial succeeds.")
}

package main

import (
	"fmt"
	"os"

	"github.com/murphybytes/udt.go/udt"
)

func main() {
	udt.Startup()
	defer udt.Cleanup()
	fmt.Println("Starting server test.")

	l, e := udt.Listen("127.0.0.1:9876")
	if e != nil {
		fmt.Printf("Listen failed: %s\n", e.Error())
		os.Exit(1)
	}
	defer l.Close()

	conn, e := l.Accept()

	if e != nil {
		fmt.Printf("Accept failed: %s\n", e.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Success")
}

package main

import (
	"fmt"
	"os"

	"github.com/murphybytes/udt.go/udt"
)

func main() {
	udt.Startup()
	defer udt.Cleanup()
	fmt.Println("Starting client test.")
	conn, e := udt.Dial("127.0.0.1:9876")

	if e != nil {
		fmt.Printf("Dial fails with %s\n", e.Error())
		os.Exit(1)
	}

	defer conn.Close()
	var n int
	n, e = conn.Write([]byte("Hello World!"))

	if e != nil {
		fmt.Printf("Write fails with %s\n", e.Error())
		os.Exit(1)
	}
	fmt.Printf("Wrote %d bytes to sock\n", n)
	fmt.Println("Client succeeds. ")
}

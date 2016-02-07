package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/murphybytes/udt.go/test/helper"
	"github.com/murphybytes/udt.go/udt"
)

func main() {
	var cla *helper.CommandLineArgs
	cla = helper.NewCommandLineArgs()

	udt.Startup()
	defer udt.Cleanup()
	fmt.Println("Starting server test.")

	l, e := udt.Listen(cla.ServerAddress)
	if e != nil {
		fmt.Printf("Listen failed: %s\n", e.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, e := l.Accept()

		if e != nil {
			fmt.Printf("Accept failed: %s\n", e.Error())
			os.Exit(1)
		}

		go func() {
			defer conn.Close()

			reader := bufio.NewReader(conn)
			readbuf, e := reader.ReadBytes(helper.ETX)

			if e != nil {
				fmt.Printf("Server read failed with %s\n", e.Error())
				os.Exit(1)
			}

			writer := bufio.NewWriter(conn)
			_, e = writer.Write(readbuf)
			if e != nil {
				fmt.Printf("Server write failed %s\n", e.Error())
				os.Exit(1)
			}

			e = writer.Flush()
			if e != nil {
				fmt.Println(e.Error())
			}

		}()
	}

}

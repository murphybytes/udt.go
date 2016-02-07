package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/murphybytes/udt.go/test/helper"
	"github.com/murphybytes/udt.go/udt"
)

func main() {
	cla := helper.NewCommandLineArgs()
	udt.Startup()
	defer udt.Cleanup()

	conn, e := udt.Dial(cla.ServerAddress)

	if e != nil {
		fmt.Printf("Dial fails with %s\n", e.Error())
		os.Exit(1)
	}

	defer conn.Close()

	writer := bufio.NewWriter(conn)
	writebuff := []byte{}
	writebuff = append(writebuff, []byte(cla.TestBuffer)...)
	writebuff = append(writebuff, helper.ETX)

	_, e = writer.Write(writebuff)
	if e != nil {
		fmt.Printf("Client write fails %s\n", e.Error())
		os.Exit(1)
	}

	e = writer.Flush()

	reader := bufio.NewReader(conn)

	readbuff, e := reader.ReadBytes(helper.ETX)

	fmt.Println(string(readbuff))

	if string(writebuff) != string(readbuff) {
		fmt.Printf("GOT: %s\n", string(readbuff))
		fmt.Printf("EXPECTED: %s\n", string(writebuff))
		os.Exit(1)
	}

}

package helper

import "flag"

const ETX byte = 4

type CommandLineArgs struct {
	ServerAddress    string
	ServerBufferSize int
	ClientBufferSize int
	TestBuffer       string
}

func NewCommandLineArgs() *CommandLineArgs {
	cla := &CommandLineArgs{}
	flag.StringVar(&cla.ServerAddress, "ip", "127.0.0.1:9876", "IP Address of the server.")
	flag.StringVar(&cla.TestBuffer, "tb", "Hello World!", "Bytes used for test.")
	flag.IntVar(&cla.ServerBufferSize, "sb", 100, "Size of the server network buffer")
	flag.IntVar(&cla.ClientBufferSize, "cb", 100, "Size of the client network buffer")
	flag.Parse()
	return cla
}

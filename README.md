UDT.go A Golang wrapper for the UDT project
=================================================

UDP-based Data Transfer Protocol (UDT), is a high-performance data transfer protocol
designed for transferring large volumetric datasets over high-speed wide area networks.
Such settings are typically disadvantageous for the more common TCP protocol.
The original C++ UDT project can be found on SourceForge http://udt.sourceforge.net/

### Getting Started

First You'll need to compile the UDT C source code.  To do this run the build script
that is located in the root directory of the project.  You'll need g++ on your
system to successfully compile the UDT C library.  I've only tested the build on
x86_64 OSX so you may need to tweak the build script for your system.  If you
do get the build working on other systems please submit a pull request.  Once you
build the C source you can build UDT by setting your GOPATH and PATH, then building
github.com/murphybytes/udt.go/...

### Testing

There is a test script 'test.sh' in the root directory of the project.  Set your
GOPATH, then your PATH to $GOPATH/bin then run the test script from the project
root.  The test script will start a UDT server that will listen for text strings
from clients, the clients will send the string to the server, the server sends them
back to client performing a comparison.  The go test programs are in udt.go/test/client
udt.go/test/server

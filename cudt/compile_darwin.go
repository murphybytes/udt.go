package cudt

// #cgo CXXFLAGS: -I${SRCDIR}/../vendor/udt/src
// #cgo LDFLAGS: -L${SRCDIR}/../vendor/udt/src -ludt -lstdc++ -lpthread -lm
import "C"

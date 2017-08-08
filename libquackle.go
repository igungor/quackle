// XXX: Note that `go build` fails if the swig interface filename is the same
// as this filename. So we can't rename this file to quackle.go.

package quackle

// #cgo CXXFLAGS: -I${SRCDIR}/_include
// #cgo pkg-config: QtCore
// #cgo LDFLAGS: -L${SRCDIR}/_lib -lquackle -lquackleio -lQtCore
import "C"

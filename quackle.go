package quackle

// #cgo CXXFLAGS: -I${SRCDIR}/_include
// #cgo pkg-config: QtCore
// #cgo LDFLAGS: -L${SRCDIR}/_lib -lquackle -lquackleio -lQtCore
import "C"

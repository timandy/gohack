package gohack

import (
	"reflect"
	"unsafe"

	_ "github.com/timandy/gohack/g"
)

// getgp returns the pointer to the current runtime.g.
//
//go:linkname getgp github.com/timandy/gohack/g.getgp
func getgp() unsafe.Pointer

// getgt returns the type of runtime.g.
//
//go:linkname getgt github.com/timandy/gohack/g.getgt
func getgt() reflect.Type

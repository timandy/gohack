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

// getg0 returns the value of runtime.g0.
//
//go:linkname getg0 github.com/timandy/gohack/g.getg0
func getg0() any

// getgt returns the type of runtime.g.
//
//go:linkname getgt github.com/timandy/gohack/g.getgt
func getgt() reflect.Type

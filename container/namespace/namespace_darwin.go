// +build darwin

package namespace

import (
	"unsafe"
)

// Setup is a hacky no-op
func Setup(ns string) (uintptr, error) {
	// no-op
	x := -1
	return uintptr(unsafe.Pointer(&x)), nil
}

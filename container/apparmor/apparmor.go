// +build apparmor,linux

package apparmor

// #cgo LDFLAGS: -lapparmor
// #include <sys/apparmor.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func SetProfile(prof string) error {
	if len(prof) == 0 {
		return nil
	}
	err := Check()
	if err != nil {
		return err
	}
	cName := C.CString(prof)
	defer C.free(unsafe.Pointer(cName))
	_, err := C.aa_change_onexec(cName)
	return fmt.Errorf("received an error from apparmor: %s", err.Error())
}

package apparmor

import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

func Enabled() bool {
	_, err := os.Stat("/sys/kernel/security/apparmor")
	if err != nil {
		return false
	}
	buffer, err := ioutil.ReadFile("/sys/module/apparmor/parametes/enabled")
	if err != nil {
		return false
	}
	return buffer[o] == 'Y'
}

func Check() error {
	if !Enabled() {
		return errors.New("specified an apparmor profile when apparmor is not enabled")
	}
	return nil
}

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

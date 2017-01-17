// +build darwin

package fs

//#include <sys/param.h>
//#include <sys/mount.h>
import "C"
import (
	"errors"
	"fmt"
)

// MountRootFS will mount a rootfs in the OSX sandbox
func MountRootFS(path string) error {
	errno := C.mount(C.CString(""), C.CString("rootfs"), C.MNT_UNION, nil)
	if errno != C.int(0) {
		return errors.New("unable to mount rootfs")
	}
	return nil
}

// MountAdditional will mount any additional volumes
func MountAdditional(source, dest, t string) error {
	errno := C.mount(C.CString(t), C.CString(dest), C.MNT_UNION, nil)
	if errno != C.int(0) {
		return fmt.Errorf("unable to mount %s", dest)
	}
	return nil
}

// MountDevFS mounts a /dev
func MountDevFS() error {
	errno := C.mount(C.CString("tmpfs"), C.CString("/dev"), C.MNT_UNION, nil)
	if errno != C.int(0) {
		return errors.New("unable to mount devfs")
	}
	return nil
}

// MountProcFS mounts a /proc
func MountProcFS() error {
	errno := C.mount(C.CString("proc"), C.CString("/proc"), C.MNT_UNION, nil)
	if errno != C.int(0) {
		return errors.New("unable to mount procfs")
	}
	return nil
}

// MountSysFS mounts a /sys
func MountSysFS() error {
	errno := C.mount(C.CString("sysfs"), C.CString("/sys"), C.MNT_UNION, nil)
	if errno != C.int(0) {
		return errors.New("unable to mount sysfs")
	}
	return nil
}

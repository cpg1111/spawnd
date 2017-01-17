package fs

import (
	"syscall"
)

func MountRootFS(path string) error {
	return syscall.Mount(path, "rootfs", "", syscall.MS_BIND, "")
}

func MountAdditional(source, dest, t string) error {
	return syscall.Mount(source, dest, t, syscall.MS_BIND, "")
}

func MountDevFS() error {
	return syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_BIND, "")
}

func MountProcFS() error {
	return syscall.Mount("proc", "/proc", "proc", syscall.MS_BIND, "")
}

func MountSysFS() error {
	return syscall.Mount("sysfs", "/sys", "sysfs", syscall.MS_BIND, "")
}

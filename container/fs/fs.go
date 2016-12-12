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

package namespace

import (
	"fmt"
	"syscall"
)

func Setup(ns string) (uintptr, error) {
	switch ns {
	case "uts":
		return syscall.CLONE_NEWUTS, nil
	case "ipc":
		return syscall.CLONE_NEWIPC, nil
	case "user":
		return syscall.CLONE_NEWUSER, nil
	case "mount":
		return syscall.CLONE_NEWNS, nil
	case "net":
		return syscall.CLONE_NEWNET, nil
	// NOT SUPPORTED IN GO
	//case "cgroup":
	//	return syscall.CLONE_NEWCGROUP, nil
	default:
		return 0x0, fmt.Errorf("unsupported namespace %s", ns)
	}
}

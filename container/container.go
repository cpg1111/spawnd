package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...))
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runErr := cmd.Run()
	if runErr != nil {
		log.Fatal("ERROR:", runErr)
	}
}

func child() {
	mountErr := syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, "")
	if mountErr != nil {
		log.Fatal("ERROR:", mountErr)
	}
	mkdirErr := os.MkdirAll("rootfs/copy", 0700)
	if mkdirErr != nil {
		log.Fatal("ERROR:", mkdirErr)
	}
	pivotErr := syscall.PivotRoot("rootfs", "rootfs/copy")
	if pivotErr != nil {
		log.Fatal("ERROR:", pivotErr)
	}
	chdirErr := os.Chdir("/")
	if chdirErr != nil {
		log.Fatal("ERROR:", chdirErr)
	}
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runErr := cmd.Run()
	if runErr != nil {
		log.Fatal(runErr)
	}
}

func main() {
	switch os.Args[1] {
	case "exec":
		parent()
	case "child":
		child()
	default:
		panic("incorrect process")
	}
}

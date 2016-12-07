package pkg

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func Parent() {
	cmd := exec.Command("/proc/self/exe", "child")
	if len(os.Args) > 2 {
		cmd.Args = append(cmd.Args, os.Args[2:]...)
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runErr := cmd.Run()
	if runErr != nil {
		log.Fatal("ERROR:", runErr)
	}
}

func Child() {
	// TODO: chroot fs
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runErr := cmd.Run()
	if runErr != nil {
		log.Fatal(runErr)
	}
}

package pkg

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/cpg1111/spawnd/container/oci"
)

func Parent(conf *oci.Config) {
	cmd := exec.Command("/proc/self/exe", "child", os.Args[2])
	cloneFlags, nsErr := oci.SetupNamespaces(conf)
	if nsErr != nil {
		panic(nsErr)
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: oci.SetUser(conf),
		CloneFlags: cloneFlags,
		Setctty:    conf.Process.Terminal,
		Noctty:     !conf.Process.Terminal,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runErr := cmd.Run()
	if runErr != nil {
		log.Fatal("ERROR:", runErr)
	}
}

func Child(conf *oci.Config) {
	err := oci.SetupFS(conf)
	if err != nil {
		panic(err)
	}
	err = oci.SetCWD(conf)
	if err != nil {
		panic(err)
	}
	err = oci.SetEnv(conf)
	if err != nil {
		panic(err)
	}
	err = oci.SetCaps(conf)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(exec.LookPath(conf.Process.Args[0]))
	if len(conf.Process.Args) > 1 {
		cmd.Args = append(cmd.Args, conf.Process.Args[1:]...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runErr := cmd.Run()
	if runErr != nil {
		log.Fatal(runErr)
	}
}

// +build linux

package pkg

import (
	"context"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/cpg1111/spawnd/container/hooks"
	"github.com/cpg1111/spawnd/container/oci"
)

func Parent(conf oci.Config) {
	cmd := exec.Command("/proc/self/exe", "child", os.Args[2])
	cloneFlags, nsErr := oci.SetupNamespaces(conf)
	if nsErr != nil {
		panic(nsErr)
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: oci.SetUser(conf),
		Cloneflags: cloneFlags,
		Setctty:    conf.GetProcess().Terminal,
		Noctty:     !conf.GetProcess().Terminal,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	runErr := cmd.Run()
	if runErr != nil {
		log.Fatal("ERROR:", runErr)
	}
}

func setupChild(conf oci.Config) {
	err := oci.SetupFS(conf)
	if err != nil {
		panic(err)
	}
	err = oci.SetCWD(conf)
	if err != nil {
		panic(err)
	}
	err = oci.SetupEnv(conf)
	if err != nil {
		panic(err)
	}
	err = oci.SetCaps(conf)
	if err != nil {
		panic(err)
	}
	err = oci.SetRLimits(conf)
	if err != nil {
		panic(err)
	}
	setAdditional()
	//TODO: SELinux
	//TODO: noNewPrivileges
}

func execChild(conf oci.Config) {
	proc := conf.GetProcess()
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	binPath, err := exec.LookPath(proc.Args[0])
	if err != nil {
		panic(err)
	}
	cmd := exec.CommandContext(ctx, binPath)
	if len(proc.Args) > 1 {
		cmd.Args = append(cmd.Args, proc.Args[1:]...)
	}
	if len(proc.Env) > 0 {
		cmd.Env = append(cmd.Env, proc.Env...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	hookMgr := hooks.New(cmd)
	hookMgr.Run(conf)
}

func Child(conf oci.Config) {
	setupChild(conf)
	execChild(conf)
}

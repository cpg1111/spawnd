package pkg

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/cpg1111/spawnd/container/hooks"
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

func setupChild(conf *oci.Config) {
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
	err = oci.SetRLimits(conf)
	if err != nil {
		panic(err)
	}
	err = oci.SetupAppArmor(conf)
	if err != nil {
		panic(err)
	}
	//TODO: SELinux
	//TODO: noNewPrivileges
}

func execChild(conf *oci.Config) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, conf.Process.Timeout*time.Second)
	cmd := exec.CommandContext(exec.LookPath(conf.Process.Args[0]))
	if len(conf.Process.Args) > 1 {
		cmd.Args = append(cmd.Args, conf.Process.Args[1:]...)
	}
	if len(conf.Process.Env) > 0 {
		cmd.Env = append(cmd.Env, conf.Process.Env)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	hookMgr := hooks.New(conf)
	hookMgr.Run()
}

func Child(conf *oci.Config) {
	setupChild(conf)
	execChild(conf)
}

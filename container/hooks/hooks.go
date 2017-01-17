package hooks

import (
	"context"
	"os/exec"
	"time"

	"github.com/cpg1111/spawnd/container/oci"
)

type HooksMgr interface {
	Prestart(conf oci.Config) error
	Poststart(conf oci.Config) error
	Poststop(conf oci.Config) error
	Run(conf oci.Config)
}

type hooksMgr struct {
	HooksMgr
	MainCMD *exec.Cmd
}

func execHook(path string, args, env []string, timeout int) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Env = env
	return cmd.Run()
}

func execHooks(hooks []oci.Hook) error {
	for _, hook := range hooks {
		err := execHook(hook.Path, hook.Args, hook.Env, hook.Timeout)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h hooksMgr) Prestart(conf oci.Config) error {
	return execHooks(conf.GetHooks().PreStart)
}

func (h hooksMgr) Poststart(conf oci.Config) error {
	return execHooks(conf.GetHooks().PostStart)
}

func (h hooksMgr) Poststop(conf oci.Config) error {
	return execHooks(conf.GetHooks().PostStop)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (h hooksMgr) Run(conf oci.Config) {
	must(h.Prestart(conf))
	must(h.MainCMD.Start())
	must(h.Poststart(conf))
	must(h.MainCMD.Wait())
	must(h.Poststop(conf))
}

func New(main *exec.Cmd) HooksMgr {
	return hooksMgr{
		MainCMD: main,
	}
}

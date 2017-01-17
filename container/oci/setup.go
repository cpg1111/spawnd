package oci

import (
	"os"
	"syscall"

	"github.com/cpg1111/spawnd/container/fs"
)

func SetupFS(conf Config) error {
	var (
		hasMountedProc bool
		hasMountedSys  bool
		hasMountedDev  bool
	)
	rootPath := conf.GetRoot().Path
	osConf := conf.GetOS()
	err := fs.MountRootFS(rootPath)
	if err != nil {
		return err
	}
	err = syscall.Chroot(rootPath)
	if err != nil {
		return err
	}
	for _, m := range conf.GetMounts() {
		if m.Type == "proc" {
			hasMountedProc = true
		} else if m.Type == "sysfs" {
			hasMountedSys = true
		} else if m.Destination == "/dev" {
			hasMountedDev = true
		}
		err = fs.MountAdditional(m.Source, m.Destination, m.Type)
		if err != nil {
			return err
		}
	}
	if len(osConf.GetDevices()) > 0 {
		if !hasMountedProc {
			err = fs.MountProcFS()
			if err != nil {
				return err
			}
		}
		if !hasMountedSys {
			err = fs.MountSysFS()
			if err != nil {
				return err
			}
		}
		if !hasMountedDev {
			err = fs.MountDevFS()
			if err != nil {
				return err
			}
		}
		for _, d := range osConf.GetDevices() {
			err = fs.MountAdditional(d.Path, d.Path, d.Type)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SetupNamespaces(conf Config) (uintptr, error) {
	return conf.SetupNamespace()
}

func SetUser(conf Config) *syscall.Credential {
	proc := conf.GetProcess()
	return &syscall.Credential{
		Uid:    proc.User.UID,
		Gid:    proc.User.GID,
		Groups: proc.User.AdditionalGIDs,
	}
}

func SetCWD(conf Config) error {
	proc := conf.GetProcess()
	if len(proc.CWD) > 0 {
		return os.Chdir(proc.CWD)
	}
	return nil
}

func SetupEnv(conf Config) error {
	for _, e := range conf.GetProcess().Env {
		env, err := ParseENVStr(e)
		if err != nil {
			return err
		}
		for k := range env {
			os.Setenv(k, env[k])
		}
	}
	return nil
}

func SetCaps(conf Config) error {
	return conf.SetCaps()
}

func setRLimit(ty string, hard, soft int) error {
	resource, err := rlimitType(ty)
	if err != nil {
		return err
	}
	rlimit := &syscall.Rlimit{
		Cur: uint64(soft),
		Max: uint64(hard),
	}
	return syscall.Setrlimit(resource, rlimit)
}

func SetRLimits(conf Config) error {
	for _, rlim := range conf.GetProcess().RLimits {
		err := setRLimit(rlim.Type, rlim.Hard, rlim.Soft)
		if err != nil {
			return err
		}
	}
	return nil
}

//TODO: setup SELinux

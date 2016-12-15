package oci

import (
	"syscall"

	"github.com/syndtr/gocapability/capability"

	"github.com/cpg1111/spawnd/container/apparmor"
	"github.com/cpg1111/spawnd/container/fs"
	"github.com/cpg1111/spawnd/container/namespace"
)

func SetupFS(conf *Config) error {
	hasMountedDev := false
	err := fs.MountRootFS(conf.Root.Path)
	if err != nil {
		return err
	}
	err = syscall.Chroot(conf.Root.Path)
	if err != nil {
		return err
	}
	for _, m := range conf.Mounts {
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
	if len(conf.Linux.Devices) > 0 {
		if !hasMountedDev {
			err = fs.MountDevFS()
			if err != nil {
				return err
			}
		}
		for _, d := range conf.Linux.Devices {
			err = fs.MountAdditional(d.Path, d.Path, d.Type)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SetupNamespaces(conf *Config) (uintptr, error) {
	var err error
	flags := syscall.CLONE_NEWPID
	for _, n := range conf.Namespaces {
		newFlag, err := namespace.Setup(n.Type)
		if err != nil {
			return err
		}
		flags = flags | newFlag
	}
	return flags, err
}

func SetUser(conf *Config) *syscall.Credential {
	return &syscall.Credential{
		Uid:    conf.Process.User.UID,
		Gid:    conf.Process.User.GID,
		Groups: conf.Process.User.Groups,
	}
}

func SetCWD(conf *Config) error {
	if len(conf.Process.CWD) > 0 {
		return os.Chdir(conf.Process.CWD)
	}
	return nil
}

func SetupEnv(conf *Config) error {
	for _, e := range conf.Process.Env {
		err := ParseENVStr(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetCaps(conf *Config) error {
	c, err := capability.NewPid(0)
	if err != nil {
		return err
	}
	caps := conf.Process.Capabilities
	var capabilities []capability.Caps
	for i := range caps {
		capabilities = append(capabilities, CapStrToVal(caps[i]))
	}
	c.Set(capability.CAPS, capabilities...)
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

func SetRLimits(conf *Config) error {
	for _, rlim := range conf.Process.Rlimits {
		err := setRLimit(rlim.Type, rlim.Hard, rlim.Soft)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetupAppArmor(conf *Config) error {
	return apparmor.SetProfile(conf.Process.AppArmorProfile)
}

//TODO: setup SELinux

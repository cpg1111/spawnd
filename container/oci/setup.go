package oci

import (
	"fmt"
	"syscall"

	"github.com/syndtr/gocapability/capability"

	"github.com/cpg1111/spawnd/container/apparmor"
)

func mountRootFS(conf *Config) error {
	return syscall.Mount(conf.Root.Path, "rootfs", "", syscall.MS_BIND, "")
}

func mountAdditional(mountConf *mounts) error {
	return syscall.Mount(mountConf.Source, mountConf.Destination, mountConf.Type, syscall.MS_BIND, "")
}

func SetupFS(conf *Config) error {
	hasMountedDev := false
	err := mountRootFS(conf)
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
		err = mountAdditional(m)
		if err != nil {
			return err
		}
	}
	if len(conf.Linux.Devices) > 0 {
		if !hasMountedDev {
			err = mountDevFS(conf)
			if err != nil {
				return err
			}
		}
		for _, d := range conf.Linux.Devices {
			err = mountAdditional(m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SetupNamespaces(conf *Config) (uintptr, error) {
	flags := syscall.CLONE_NEWPID
	for _, n := range conf.Namespaces {
		switch n.Type {
		case "uts":
			flags = flags | syscall.CLONE_NEWUTS
			break
		case "ipc":
			flags = flags | syscall.CLONE_NEWIPC
			break
		case "user":
			flags = flags | syscall.CLONE_NEWUSER
			break
		case "mount":
			flags = flags | syscall.CLONE_NEWNS
			break
		case "net":
			flags = flags | syscall.CLONE_NEWNET
			break
		case "cgroup":
			flags = flags | syscall.CLONE_NEWCGROUP
			break
		default:
			return 0x0, fmt.Errorf("unsupported namespace %s", n.Type)
		}
	}
	return flags, nil
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

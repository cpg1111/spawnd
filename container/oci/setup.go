package oci

import (
	"fmt"
	"syscall"
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

func SetupUser(conf *Config) error {
	uid := 1
	if conf.Process.User.UID != nil {
		uid = conf.Process.User.UID
	}
	err := syscall.Setuid(uid)
	if err != nil {
		return err
	}
	gid := 1
	if conf.Process.User.GID != nil {
		gid = conf.Process.User.GID
	}
	if len(conf.Process.User.AdditonalGIDs) > 0 {
		err = syscall.Setgroups(append(conf.Process.User.AdditionalGIDs, gid))
		return err
	}
	return syscall.Setgid(gid)
}

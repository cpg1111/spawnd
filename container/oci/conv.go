package oci

import (
	"fmt"
	"syscall"

	"github.com/syndtr/gocapability/capability"
)

func ParseENVStr(envvar string) (res map[string]string, err error) {
	var (
		key string
		val string
	)
	begin := 0
	for i := range envvar {
		if envvar[i] == '=' {
			key = envvar[begin:i]
			val = envvar[i+1:]
			break
		}
	}
	if len(key) == 0 {
		return nil, fmt.Errorf("invalid env variable string '%s'", envvar)
	}
	res[key] = val
	return
}

func CapStrToVal(cap string) (capability.Cap, error) {
	switch cap {
	case "CAP_CHOWN":
		return capability.CAP_CHOWN, nil
	case "CAP_DAC_OVERRIDE":
		return capability.CAP_DAC_OVERRIDE, nil
	case "CAP_DAC_READ_SEARCH":
		return capability.CAP_DAC_READ_SEARCH, nil
	case "CAP_FOWNER":
		return capability.CAP_FOWNER, nil
	case "CAP_FSETID":
		return capability.CAP_FSETID, nil
	case "CAP_KILL":
		return capability.CAP_KILL, nil
	case "CAP_SETGID":
		return capability.CAP_SETGID, nil
	case "CAP_SETUID":
		return capability.CAP_SETUID, nil
	case "CAP_LINUX_IMMUTABLE":
		return capability.CAP_LINUX_IMMUTABLE, nil
	case "CAP_NET_BIND_SERVICE":
		return capability.CAP_NET_BIND_SERVICE, nil
	case "CAP_NET_BROADCAST":
		return capability.CAP_NET_BROADCAST, nil
	case "CAP_NET_ADMIN":
		return capability.CAP_NET_ADMIN, nil
	case "CAP_NET_RAW":
		return capability.CAP_NET_RAW, nil
	case "CAP_IPC_LOCK":
		return capability.CAP_IPC_LOCK, nil
	case "CAP_IPC_OWNER":
		return capability.CAP_IPC_OWNER, nil
	case "CAP_SYS_MODULE":
		return capability.CAP_SYS_MODULE, nil
	case "CAP_SYS_RAWIO":
		return capability.CAP_SYS_RAWIO, nil
	case "CAP_SYS_CHROOT":
		return capability.CAP_SYS_CHROOT, nil
	case "CAP_SYS_PTRACE":
		return capability.CAP_SYS_PTRACE, nil
	case "CAP_SYS_PACCT":
		return capability.CAP_SYS_PACCT, nil
	case "CAP_SYS_ADMIN":
		return capability.CAP_SYS_ADMIN, nil
	case "CAP_SYS_BOOT":
		return capability.CAP_SYS_BOOT, nil
	case "CAP_SYS_NICE":
		return capability.CAP_SYS_NICE, nil
	case "CAP_SYS_RESOURCE":
		return capability.CAP_SYS_RESOURCE, nil
	case "CAP_SYS_TIME":
		return capability.CAP_SYS_TIME, nil
	case "CAP_SYS_TTY_CONFIG":
		return capability.CAP_SYS_TTY_CONFIG, nil
	case "CAP_MKNOD":
		return capability.CAP_MKNOD, nil
	case "CAP_LEASE":
		return capability.CAP_LEASE, nil
	case "CAP_AUDIT_WRITE":
		return capability.CAP_AUDIT_WRITE, nil
	case "CAP_AUDIT_CONTROL":
		return capability.CAP_AUDIT_CONTROL, nil
	case "CAP_SETFCAP":
		return capability.CAP_SETFCAP, nil
	case "CAP_MAC_OVERRIDE":
		return capability.CAP_MAC_OVERRIDE, nil
	case "CAP_MAC_ADMIN":
		return capability.CAP_MAC_ADMIN, nil
	case "CAP_SYSLOG":
		return capability.CAP_SYSLOG, nil
	case "CAP_WAKE_ALARM":
		return capability.CAP_WAKE_ALARM, nil
	case "CAP_BLOCK_SUSPEND":
		return capability.CAP_BLOCK_SUSPEND, nil
	case "CAP_AUDIT_READ":
		return capability.CAP_AUDIT_READ, nil
	default:
		return -1, fmt.Errorf("invalid capability specified: '%s'", cap)
	}
}

func rlimitType(ty string) (int, error) {
	switch ty {
	case "RLIMIT_AS":
		return syscall.RLIMIT_AS, nil
	case "RLIMIT_CORE":
		return syscall.RLIMIT_CORE, nil
	case "RLIMIT_CPU":
		return syscall.RLIMIT_CPU, nil
	case "RLIMIT_DATA":
		return syscall.RLIMIT_DATA, nil
	case "RLIMIT_FSIZE":
		return syscall.RLIMIT_FSIZE, nil
	case "RLIMIT_NOFILE":
		return syscall.RLIMIT_NOFILE, nil
	case "RLIMIT_STACK":
		return syscall.RLIMIT_STACK, nil
	default:
		return -1, fmt.Errorf("invalid rlimit specified: %s", ty)
	}
}

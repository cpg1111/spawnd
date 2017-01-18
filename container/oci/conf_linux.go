package oci

import (
	"syscall"

	"github.com/syndtr/gocapability/capability"

	"github.com/cpg1111/spawnd/container/namespace"
)

type linux struct {
	OS
	Namespaces        []Namespace       `json:"namespaces",omitempty`
	UIDMapping        []mapping         `json:"uidMapping",omitempty`
	GIDMapping        []mapping         `json:"gidMapping",omitempty`
	Devices           []device          `json:"devices",omitempty`
	Sysctl            map[string]string `json:"sysctl"`
	CGRoupPath        string            `json:"cgroupPath"`
	Resources         cgroupResource    `json:"resources"`
	RootFSPropagation string            `json:"rootfsPropagation"`
	Seccomp           seccomp           `json:"seccomp"`
	MaskedPaths       []string          `json:"maskedPaths"`
	ReadonlyPaths     []string          `json:"readonlyPaths"`
	MountLabel        string            `json:"mountLabel"`
}

func (l linux) GetDevices() []device {
	return l.Devices
}

func (l linux) GetNamespaces() []Namespace {
	return l.Namespaces
}

type linuxConfig struct {
	Config
	OCIVersion string   `json:"ociversion"`
	Root       rootfs   `json:"root"`
	Mounts     []mount  `json:"mounts"`
	Process    process  `json:"process"`
	HostName   string   `json:"hostname"`
	Platform   platform `json:"platform"`
	Linux      linux    `json:"linux",omitempty`
	Hooks      hooks    `json:"hooks",omitempty`
}

func LoadConfig(path string) (Config, error) {
	conf := linuxConfig{}
	loaded, err := loadConfig(conf, path)
	return *loaded, err
}

func (l linuxConfig) GetRoot() rootfs {
	return l.Root
}

func (l linuxConfig) GetMounts() []mount {
	return l.Mounts
}

func (l linuxConfig) GetProcess() process {
	return l.Process
}

func (l linuxConfig) GetHostName() string {
	return l.HostName
}

func (l linuxConfig) GetPlatform() platform {
	return l.Platform
}

func (l linuxConfig) GetOS() OS {
	return l.Linux
}

func (l linuxConfig) GetHooks() hooks {
	return l.Hooks
}

func (l linuxConfig) SetCaps() error {
	c, err := capability.NewPid(0)
	if err != nil {
		return err
	}
	caps := l.GetProcess().Capabilities
	var capabilities []capability.Cap
	for i := range caps {
		cap, err := CapStrToVal(caps[i])
		if err != nil {
			return err
		}
		capabilities = append(capabilities, cap)
	}
	c.Set(capability.CAPS, capabilities...)
	return nil
}

func (l linuxConfig) SetupNamespaces() (uintptr, error) {
	flags := uintptr(syscall.CLONE_NEWPID)
	for _, n := range l.GetOS().GetNamespaces() {
		newFlag, err := namespace.Setup(n.Type)
		if err != nil {
			return flags, err
		}
		flags = flags | newFlag
	}
	return flags, nil
}

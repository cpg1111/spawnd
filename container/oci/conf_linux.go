package oci

import (
	"encoding/json"
	"os"
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
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	conf := &linuxConfig{}
	err = json.NewEncoder(file).Encode(conf)
	return *conf, err
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

package oci

import (
	"encoding/json"
	"os"
	"syscall"
)

type rootfs struct {
	Path     string `json:"path"`
	ReadOnly bool   `json:"readonly"`
}

type mounts struct {
	Destination string   `json:"destination"`
	Type        string   `json:"type"`
	Source      string   `json:"source"`
	Options     []string `json:"option"`
}

type consoleSize struct {
	Height uint `json:"height"`
	Width  uint `json:"width"`
}

type user struct {
	UID            int `json:"uid"`
	GID            int `json:"gid"`
	AdditionalGIDs `json:"additionalGids"`
}

type process struct {
	Terminal        bool             `json:"terminal"`
	ConsoleSize     consoleSize      `json:"consoleSize"`
	CWD             string           `json:"cwd"`
	Env             []string         `json:"env"`
	Args            []string         `json:"args"`
	Capabilities    []string         `json:"capabilities"`
	RLimits         []syscall.Rlimit `json:"rlimits"`
	AppArmorProfile string           `json:"apparmorprofile"`
	SELinuxLabel    string           `json:"selinuxLabel"`
	NoNewPrivileges bool             `json:"noNewPrivileges"`
	User            user             `json:"user"`
}

type platform struct {
	OS   string `json:"os"`
	ARCH string `json:"arch"`
}

type namespace struct {
	Type string `json:"type"`
	Path string `json:"path",omitempty`
}

type mapping struct {
	HostID      int `json:"hostID"`
	ContainerID int `json:"containerID"`
	Size        int `json:"size"`
}

type device struct {
	Type     string `json:"type"`
	Path     string `json:"path"`
	Major    int64  `json:"major"`
	Minor    int64  `json:"minor"`
	FileMode uint32 `json:"fileMode"`
	UID      uint32 `json:"uid"`
	GID      uint32 `json:"gid"`
}

type memory struct {
	Limit       int `json:"limit"`
	Reservation int `json:"reservation"`
	Swap        int `json:"swap"`
	Kernel      int `json:"kernel"`
	KernelTCP   int `json:"kernelTCP"`
	Swappiness  int `json:"swappiness"`
}

type cgroupdev struct {
	Allow  bool   `json:"allow"`
	Access string `json:"access"`
	Type   string `json:"type"`
	Major  int64  `json:"major"`
	Minor  int64  `json:"minor"`
}

type priority struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

type network struct {
	ClassID    int `json:"classID"`
	Priorities []priority
}

type pids struct {
	Limit int
}

type hpLimit struct {
	PageSize `json:"pageSize"`
	Limit    int64 `json:"limit"`
}

type cpu struct {
	Shares          int    `json:"shares"`
	Quota           int    `json:"quota"`
	Period          int    `json:"period"`
	RealTimeRuntime int    `json:"realtimeRuntime"`
	RealTimePeriod  int    `json"realtimePeriod"`
	CPUs            string `json:"cpus"`
	MEMs            string `json:"mems"`
}

type blkioDevice struct {
	Major      int `json:"major"`
	Minor      int `json:"minor"`
	Weight     int `json:"weight",omitempty`
	LeafWeight int `json:"leafWeight",omitempty`
}

type blockIO struct {
	BLKIOWeight                  int           `json:"blkioWeight"`
	BLKIOLeafWeight              int           `json:"blkioLeafWeight"`
	BLKIOWeightDevice            []blkioDevice `json:"blkioWeightDevice"`
	BLKIOThrottleReadBpsDevice   []blkioDevice `json:"blkioThrottleReadBpsDevice"`
	BLKIOThrottleWriteIOPSDevice []blkioDevice `json:"blkioThrottleWriteIOPSDevice"`
}

type cgroupResource struct {
	Memory           memory    `json:"memory"`
	Devices          cgroupdev `json:"devices"`
	Network          network   `json:"network"`
	PIDs             pids      `json:"pids"`
	HugePageLimits   []hpLimit `json:"hugepageLimits"`
	OOMScoreAdj      int       `json:"oomScoreAdj"`
	CPU              cpu       `json:"cpu"`
	DisableOOMKiller bool      `json:"disableOOMKiller"`
	BlockIO          blockIO   `json:"blockIO"`
}

type seccompSyscall struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

type seccomp struct {
	DefaultAction string           `json:"defaultAction"`
	Architectures []string         `json:"architectures"`
	Syscalls      []seccompSyscall `json:"syscalls"`
}

type linux struct {
	Namespaces        []namespace       `json:"namespaces",omitempty`
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

type Hook struct {
	Path   string   `json:"path"`
	Args   []string `json:"args"`
	Env    []string `json:"env"`
	Timout int      `json:"timeout"`
}

type hooks struct {
	PreStart  []hook `json:"prestart"`
	PostStart []hook `json:"poststart"`
	PostStop  []hook `json:"poststop"`
}

type Config interface{}

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

type bsdConfig struct {
	Config
	OCIVersion string   `json:"ociversion"`
	Root       rootfs   `json:"root"`
	Mounts     []mount  `json:"mounts"`
	Process    process  `json:"process"`
	HostName   string   `json:"hostname"`
	Platform   platform `json:"platform"`
	Hooks      hooks    `json:"hooks",omitempty`
}

type darwinConfig struct {
	Config
	OCIVersion string   `json:"ociversion"`
	Root       rootfs   `json:"root"`
	Mounts     []mount  `json:"mounts"`
	Process    process  `json:"process"`
	HostName   string   `json:"hostname"`
	Platform   platform `json:"platform"`
	Hooks      hooks    `json:"hooks",omitempty`
}

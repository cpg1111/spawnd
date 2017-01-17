package oci

import (
	"encoding/json"
	"os"
)

type rootfs struct {
	Path     string `json:"path"`
	ReadOnly bool   `json:"readonly"`
}

type mount struct {
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
	UID            uint32   `json:"uid"`
	GID            uint32   `json:"gid"`
	AdditionalGIDs []uint32 `json:"additionalGids"`
}

type rlimit struct {
	Type string `json:"type"`
	Hard int    `json:"hard"`
	Soft int    `json:"soft"`
}

type process struct {
	Terminal        bool        `json:"terminal"`
	ConsoleSize     consoleSize `json:"consoleSize"`
	CWD             string      `json:"cwd"`
	Env             []string    `json:"env"`
	Args            []string    `json:"args"`
	Capabilities    []string    `json:"capabilities"`
	RLimits         []rlimit    `json:"rlimits"`
	AppArmorProfile string      `json:"apparmorprofile"`
	SELinuxLabel    string      `json:"selinuxLabel"`
	NoNewPrivileges bool        `json:"noNewPrivileges"`
	User            user        `json:"user"`
}

type platform struct {
	OS   string `json:"os"`
	ARCH string `json:"arch"`
}

type Namespace struct {
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
	PageSize int64 `json:"pageSize"`
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

type Hook struct {
	Path    string   `json:"path"`
	Args    []string `json:"args"`
	Env     []string `json:"env"`
	Timeout int      `json:"timeout"`
}

type hooks struct {
	PreStart  []Hook `json:"prestart"`
	PostStart []Hook `json:"poststart"`
	PostStop  []Hook `json:"poststop"`
}

type OS interface {
	GetDevices() []device
	GetNamespaces() []Namespace
}

type Config interface {
	GetRoot() rootfs
	GetMounts() []mount
	GetProcess() process
	GetHostName() string
	GetPlatorm() platform
	GetOS() OS
	GetHooks() hooks
	SetCaps()
	SetupNamespaces()
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

func loadConfig(conf Config, path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	err = json.NewEncoder(file).Encode(&conf)
	return &conf, err
}

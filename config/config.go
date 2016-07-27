package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type UnixServer struct {
	Path  string
	User  string
	PWD   string
	Owner string
	Group string
	Mode  uint
}

type InetServer struct {
	Host string
	Port int
	User string
	PWD  string
}

type Server struct {
	Unix UnixServer
	Inet InetServer
}

type Daemon struct {
	User  string
	Group string
	Umask uint
}

type Log struct {
	STDOUTPath string
	STDERRPath string
}

type Container struct {
	CPULimit        int
	MemLimit        int
	UserNamespace   string
	NetNamespace    string
	UTSNamespace    string
	PIDNamespace    string
	MountNamespace  string
	IPCNamespace    string
	CgroupNamespace string
}

type Process struct {
	Name        string
	CMD         []string
	Priority    int
	NumProcs    int
	AutoRestart bool
	InContainer bool
	Container   Container
}

type Config struct {
	Server    Server
	Daemon    Daemon
	Logging   Log
	Processes []Process
}

func Load(path string) (*Config, error) {
	var conf Config
	confBytes, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}
	_, decodeErr := toml.Decode((string)(confBytes), &conf)
	if decodeErr != nil {
		return nil, decodeErr
	}
	return &conf, nil
}

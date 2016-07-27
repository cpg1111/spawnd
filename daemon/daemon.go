package daemon

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/cpg1111/spawnd/config"
)

type Proc interface {
	Start() (pid int, err error)
	Stop() error
	Restart() (pid int, err error)
	Reload() error
}

type Daemon struct {
	Process
	Processes []*Proc
}

func Init(conf *config.Config) (*Daemon, error) {
	self := NewProcess(&config.Process{
		Name:        "spawnd",
		CMD:         []string{"/proc/self/exe"},
		Priority:    1,
		NumProcs:    1,
		AutoRestart: false,
		InContainer: false,
	})
	var processes []*Proc
	for i := range conf.Processes {
		if conf.Processes[i].InContainer {
			// TODO NewContainer
		} else {
			append(processes, NewProcess(conf.Processes[i]))
		}
	}
	return &Daemon{
		Process:   self,
		Processes: processes,
	}
}

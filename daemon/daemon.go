package daemon

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/cpg1111/spawnd/config"
)

type Proc interface {
	GetPID() int
	GetName() string
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
	self.PID = os.Getpid()
	var processes []*Proc
	for i := range conf.Processes {
		if conf.Processes[i].InContainer {
			proc := NewContainer(conf.Processes[i])
			_, startErr := proc.Start()
			if startErr != nil {
				log.Println(startErr)
				continue
			}
			append(processes, proc)
		} else {
			proc := NewProcess(conf.Processes[i])
			_, startErr := proc.Start()
			if startErr != nil {
				log.Println(startErr)
				continue
			}
			append(processes, proc)
		}
	}
	return &Daemon{
		Process:   self,
		Processes: processes,
	}
}

func getProc(processes []*Proc, pid int) *Proc {
	if len(processes) == 1 {
		if pid == processes[0].GetPID() {
			return processes[0]
		}
		return nil
	}
	if pid < processes[len(processes)/2].GetPID() {
		return getProc(processes[0:len(processes)/2], pid)
	}
	if pid > processes[len(processes)/2].GetPID() {
		return getProc(processes[len(processes)/2:], pid)
	}
	if pid == processes[len(processes)/2].GetPID() {
		return processes[len(processes)/2]
	}
	return nil
}

func (d *Daemon) GetProc(pid int) *Proc {
	return getProc(d.Processes, pid)
}

func (d *Daemon) GetProcByName(name string) *Proc {
	for i := range d.Processes {
		if d.Processes[i].GetName() == name {
			return d.Processes[i]
		}
	}
	return nil
}

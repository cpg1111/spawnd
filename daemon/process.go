package daemon

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/cpg1111/spawnd/config"
)

type Process struct {
	Proc
	Name     string
	Command  []string
	Umask    uint
	Priority int
	UID      int
	GID      int
	PID      int
}

func NewProcess(conf *config.Process) *Process {
	return &Process{
		Name:     conf.Name,
		Command:  conf.CMD,
		Umask:    conf.Umask,
		Priority: conf.Priority,
	}
}

func (p *Process) Start() (pid int, err error) {
	bin := exec.LookPath(p.cmd[0])
	com := exec.Command(bin, p.cmd[1:]...)
	err = com.Start()
	pid = com.Process.Pid
	p.PID = pid
	return
}

func (p *Process) Stop() error {
	return syscall.Kill(p.PID, syscall.SIGTERM)
}

func (p *Process) Restart() (pid int, err error) {
	err = p.Stop()
	if err != nil {
		pid = -1
		return
	}
	pid, err = p.Start()
	p.Pid = pid
	return
}

func (p *Process) Reload() error {
	return syscall.Kill(p.Pid, syscall.SIGHUP)
}

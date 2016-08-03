package daemon

import (
	"os/exec"
	"syscall"

	"github.com/cpg1111/spawnd/config"
)

type Process struct {
	Proc
	Name     string
	Command  []string
	Priority int
	UID      int
	GID      int
	PID      int
}

func NewProcess(conf *config.Process) *Process {
	return &Process{
		Name:     conf.Name,
		Command:  conf.CMD,
		Priority: conf.Priority,
	}
}

func (p Process) GetPID() int {
	return p.PID
}

func (p *Process) SetPID(pid int) {
	p.PID = pid
}

func (p Process) GetName() string {
	return p.Name
}

func (p Process) Start() (pid int, err error) {
	bin, binErr := exec.LookPath(p.Command[0])
	if binErr != nil {
		return -1, binErr
	}
	cmd := exec.Command(bin, p.Command[1:]...)
	err = cmd.Start()
	pid = cmd.Process.Pid
	return
}

func (p Process) Stop() error {
	return syscall.Kill(p.PID, syscall.SIGTERM)
}

func (p Process) Restart() (pid int, err error) {
	err = p.Stop()
	if err != nil {
		pid = -1
		return
	}
	pid, err = p.Start()
	return
}

func (p Process) Reload() error {
	return syscall.Kill(p.PID, syscall.SIGHUP)
}

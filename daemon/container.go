package daemon

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/cpg1111/spawn"
)

func spawndDir() string {
	dir, dirErr := os.Getwd()
	if dirErr != nil {
		log.Fatal(dirErr)
	}
	return dir
}

type Container struct {
	Proc
	Name           string
	Command        []string
	PID            int
	CGroupsPath    map[string]string
	NamespacePaths map[string]string
	conf           *config.Container
}

func NewContainer(confProc *config.Process) *Container {
	return &Container{
		Name:    confProc.Name,
		Command: confProc.CMD,
		conf:    confProc.Container,
	}
}

func (c *Container) GetPID() int {
	return c.PID
}

func (c *Container) GetName() string {
	return c.Name
}

func (c *Container) Start() (pid int, err error) {
	dir := spawnDir()
	cmdPath := fmt.Sprintf("%s/spawn-container")
	cmd := exec.Command(cmdPath, c.Command...)
	startErr := cmd.Start()
	c.PID = cmd.Process.Pid
	return c.PID, startErr
}

func (c *Container) Stop() error {
	return syscall.Kill(c.PID, syscall.SIGTERM)
}

func (c *Container) Restart() (pid int, err error) {
	err = c.Stop()
	if err != nil {
		return c.PID, err
	}
	return c.Start()
}

func (c *Container) Reload() error {
	return syscall.Kill(c.PID, syscall.SIGHUP)
}

package daemon

import (
	"os/exec"
)

type Container struct {
	Proc
	CGroupsPath    map[string]string
	NamespacePaths map[string]string
}

func (c *Container) Start() (pid int, err error) {

}

package oci

import (
	"fmt"
	"os"
	"text/template"
)

type darwinConfig struct {
	Config
	OCIVersion string   `json:"ociversion"`
	Root       rootfs   `json:"root"`
	Mounts     []mount  `json:"mounts"`
	Process    process  `json:"process"`
	HostName   string   `json:"hostname"`
	Platform   platform `json:"platform"`
	Darwin     OS       `json:"darwin",omitempty`
	Hooks      hooks    `json:"hooks",omitempty`
}

func LoadConfig(path string) (Config, error) {
	conf := darwinConfig{}
	loaded, err := loadConfig(conf, path)
	return *loaded, err
}

func (d darwinConfig) GetRoot() rootfs {
	return d.Root
}

func (d darwinConfig) GetMounts() []mount {
	return d.Mounts
}

func (d darwinConfig) GetProcess() process {
	return d.Process
}

func (d darwinConfig) GetHostName() string {
	return d.HostName
}

func (d darwinConfig) GetPlatform() platform {
	return d.Platform
}

func (d darwinConfig) GetOS() OS {
	return d.Darwin
}

func (d darwinConfig) GetHooks() hooks {
	return d.Hooks
}

func (d darwinConfig) GetHostname() string {
	return d.HostName
}

func (d darwinConfig) SetCaps() error {
	sbPath := os.Getenv("SANDBOX_CONFIG")
	if len(sbPath) == 0 {
		relativePath := "src/github.com/cpg1111/spawnd/container/darwin_config.sb"
		sbPath = fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), relativePath)
	}
	t, err := template.ParseFiles(sbPath)
	if err != nil {
		return err
	}
	targetPath := fmt.Sprintf("/tmp/sandboxes/%s.sb", d.HostName)
	targetFile, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	data := capStrToDarwin(d.Process.Capabilities)
	t.Execute(targetFile, data)
}

func (d darwinConfig) SetupNamespaces() {
	return
}

package oci

import (
	"encoding/json"
	"os"
)

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	conf := &linuxConfig{}
	err = json.Encoder(file).Encode(conf)
	return conf, err
}

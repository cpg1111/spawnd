package config

import (
	"reflect"
	"testing"

	"github.com/fatih/structs"
)

var expected = &Config{
	Server: Server{
		Unix: UnixServer{
			Path:  "/var/run/spawn.sock",
			Owner: "root",
			Group: "root",
			Mode:  0755,
		},
	},
	Daemon: Daemon{
		User:  "root",
		Group: "root",
		Umask: 0644,
	},
	Logging: Log{
		STDOUTPath: "/var/log/spawnd/out/",
		STDERRPath: "/var/log/spawnd/err/",
	},
	Processes: []Process{
		Process{
			Name:        "ping_test",
			CMD:         []string{"ping", "www.google.com"},
			Priority:    1,
			NumProcs:    1,
			AutoRestart: true,
			InContainer: false,
		},
	},
}

func compare(expectedMap, testMap map[string]interface{}, t *testing.T) {
	for i := range expectedMap {
		switch i {
		case "Server":
			if testMap[i] == nil {

			}
		}
	}
}

func TestLoad(t *testing.T) {
	conf, confErr := Load("../test_conf.toml")
	if confErr != nil {
		t.Error(confErr)
	}
	expectedMap := structs.Map(*expected)
	confMap := structs.Map(*conf)
	compare(expectedMap, confMap, t)
}

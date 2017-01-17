package apparmor

import (
	"errors"
	"io/ioutil"
	"os"
)

func Enabled() bool {
	_, err := os.Stat("/sys/kernel/security/apparmor")
	if err != nil {
		return false
	}
	buffer, err := ioutil.ReadFile("/sys/module/apparmor/parametes/enabled")
	if err != nil {
		return false
	}
	return buffer[0] == 'Y'
}

func Check() error {
	if !Enabled() {
		return errors.New("specified an apparmor profile when apparmor is not enabled")
	}
	return nil
}

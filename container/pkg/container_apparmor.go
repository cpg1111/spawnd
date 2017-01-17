// +build apparmor,linux

package pkg

import (
    "github.com/cpg1111/spawnd/container/apparmor"
)

func setAdditional() {
    err := oci.SetupAppArmor(conf)
	if err != nil {
		panic(err)
	}
}

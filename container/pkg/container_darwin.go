package pkg

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cpg1111/spawnd/container/oci"
)

func Parent(conf oci.Config) {
	sbFilePath := fmt.Sprintf("/tmp/sandboxes/%s.sb", conf.GetHostName())
	sandboxBin, sandBinErr := exec.LookPath("sandbox-exec")
	if sandBinErr != nil {
		log.Fatal("ERROR:", sandBinErr)
	}
	runBin, runBinErr := exec.LookPath(os.Args[2])
	if runBinErr != nil {
		log.Fatal("ERROR:", runBinErr)
	}
	cmd := exec.Command(sandboxBin, "-f", sbFilePath, runBin)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func Child(conf oci.Config) {
	log.Fatal("ERROR: child func not supported on this OS")
}

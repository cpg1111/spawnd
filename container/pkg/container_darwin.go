package pkg

import (
	"log"
	"os"
	"os/exec"
)

func Parent() {
	sbFilePath := os.Getenv("SANDBOX_CONFIG")
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

func Child() {
	log.Fatal("ERROR: child func not supported on this OS")
}

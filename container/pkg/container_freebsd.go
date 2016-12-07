package pkg

// +build !darwin
// +build freebsd

import (
	jail "github.com/cpg1111/go-jail"
)

func concatArgs(args []string) string {
	var resultBytes []byte
	for i := range args {
		argBytes := ([]byte)(args[i])
		resultBytes = append(resultBytes, argBytes...)
	}
	return (string)(resultBytes)
}

func Parent() {
	cmd := concatArgs(os.Args[2:])
	jail.New(cmd)
}

func Child() {
	log.Fatal("ERROR: child func not supported in this OS")
}

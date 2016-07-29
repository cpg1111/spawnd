package main

import (
	"os"

	"github.com/cpg1111/spawnd/container/pkg"
)

func main() {
	switch os.Args[1] {
	case "exec":
		pkg.Parent()
	case "child":
		pkg.Child()
	default:
		panic("incorrect process")
	}
}

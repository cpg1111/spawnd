package main

import (
	"log"
	"os"

	"github.com/cpg1111/spawnd/container/pkg"
)

func main() {
	switch os.Args[1] {
	case "exec":
		log.Println("parent")
		pkg.Parent()
	case "child":
		log.Println("child")
		pkg.Child()
	default:
		log.Fatal("bad command")
	}
}

package main

import (
	"log"
	"os"

	"github.com/cpg1111/spawnd/container/oci"
	"github.com/cpg1111/spawnd/container/pkg"
)

func main() {
	conf, err := oci.LoadConfig(os.Args[2])
	if err != nil {
		panic(err)
	}
	switch os.Args[1] {
	case "exec":
		log.Println("parent")
		pkg.Parent(conf)
	case "child":
		log.Println("child")
		pkg.Child(conf)
	default:
		log.Fatal("bad command")
	}
}

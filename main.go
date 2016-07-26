package main

import (
	"flag"
	"log"

	"github.com/cpg1111/spawnd/config"
)

var (
	configPath = flag.String("config-path", "./test_conf.toml", "path to the config file, defaults to ./test_conf.toml")
)

func main() {
	flag.Parse()
	config, confErr := config.Load(*configPath)
	if confErr != nil {
		log.Fatal(confErr)
	}
	log.Println(config)
}

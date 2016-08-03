package main

import (
	"flag"
	"log"

	"github.com/cpg1111/spawnd/config"
	"github.com/cpg1111/spawnd/daemon"
	"github.com/cpg1111/spawnd/server"
)

var (
	configPath = flag.String("config-path", "./test_conf.toml", "path to the config file, defaults to ./test_conf.toml")
)

func main() {
	flag.Parse()
	conf, confErr := config.Load(*configPath)
	if confErr != nil {
		log.Fatal(confErr)
	}
	dae, initErr := daemon.Init(conf)
	if initErr != nil {
		log.Fatal(initErr)
	}
	var serverConf config.ConnServer
	if conf.Server.Unix.Path != "" {
		serverConf = conf.Server.Unix
	} else if conf.Server.Inet.Host != "" {
		serverConf = conf.Server.Inet
	} else {
		log.Fatal("no server config specified")
	}
	srv := server.New(serverConf, dae)
	log.Println("listening on", serverConf.Addr())
	runErr := srv.Run()
	if runErr != nil {
		log.Fatal(runErr)
	}
}

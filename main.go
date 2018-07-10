package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coreos/etcd/version"
	"github.com/gin-gonic/gin"
	"github.com/nanopony/web_etcd/conf"
	"github.com/nanopony/web_etcd/e3ch"
	"github.com/nanopony/web_etcd/routers"
)

const (
	PROGRAM_NAME    = "web_etcd"
	PROGRAM_VERSION = "0.0.2"
)

var configFilepath string

func init() {
	flag.StringVar(&configFilepath, "conf", "conf/config.default.ini", "config file path")
	rev := flag.Bool("rev", false, "print rev")
	flag.Parse()

	if *rev {
		fmt.Printf("[%s v%s]\n[etcd %s]\n",
			PROGRAM_NAME, PROGRAM_VERSION,
			version.Version,
		)
		os.Exit(0)
	}
}

func main() {
	config, err := conf.Init(configFilepath)
	if err != nil {
		panic(err)
	}

	client, err := e3ch.NewE3chClient(config)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.UseRawPath = true
	routers.InitRouters(router, config, client)

	router.Run(":" + config.Port)

}

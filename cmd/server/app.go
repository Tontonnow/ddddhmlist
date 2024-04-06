package main

import (
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server/trpc"
	"github.com/Tontonnow/ddddhmlist/website"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	name    = "ddddlist"
	version = "0.0.1"
)

func Init() {
	config.InitConfig()
	err := website.Init()
	if err != nil {
		return
	}
	trpc.Init()
}

func main() {
	log.Infof("Starting %s %s", name, version)
	Init()
}

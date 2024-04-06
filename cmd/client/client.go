package main

import (
	"github.com/Tontonnow/ddddhmlist/server"
	pb "github.com/Tontonnow/ddddhmlist/server/trpc"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	cfg, err := trpc.LoadConfig(trpc.ServerConfigPath)
	if err != nil {
		panic("load config fail: " + err.Error())
	}
	trpc.SetGlobalConfig(cfg)
	if err := trpc.Setup(cfg); err != nil {
		panic("setup plugin fail: " + err.Error())
	}
	proxy := pb.NewDDDDhmClientProxy(
		client.WithTarget("ip://127.0.0.1:8002"),
		client.WithProtocol("trpc"),
	)
	ctx := trpc.BackgroundContext()
	reply, err := proxy.DdddList(ctx, &server.Request{
		Url: "https://v.youku.com/v_show/id_XNTk1MjQxNjQzMg==.html",
	})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Debugf("simple  rpc   receive: %+v", reply)
}

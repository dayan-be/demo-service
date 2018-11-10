package main

import (
	"flag"
	"fmt"
	"github.com/dayan-be/demo-service/logic"
	"github.com/dayan-be/demo-service/global"
	"github.com/dayan-be/demo-service/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var (
	GitTag    = "2000.01.01.release"
	BuildTime = "2000-01-01T00:00:00+0800"
)

func main() {

	//显示版本号信息　
	version := flag.Bool("v", false, "version")
	flag.Parse()

	if *version {
		fmt.Println("Git Tag: " + GitTag)
		fmt.Println("Build Time: " + BuildTime)
		return
	}

	// log
	logrus.SetLevel(logrus.DebugLevel)

	service := micro.NewService(
		micro.Name(global.Config().Srv.SrvName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version(global.Config().Srv.Version),
		micro.Metadata(map[string]string{"ID": strconv.FormatUint(uint64(global.Config().Srv.SrvId), 10)}),
	)
	service.Init()

	demo.RegisterSayHandler(service.Server(), &logic.Handle{})

	if err := service.Run(); err != nil {
		logrus.Fatalf("service run error: %v", err)
	}
}
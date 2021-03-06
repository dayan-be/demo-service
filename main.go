package main

import (
	"github.com/dayan-be/demo-service/logic"
	"github.com/dayan-be/demo-service/proto/demo"
	"github.com/dayan-be/kit"
)


func main() {
	kit.Init()
	demo.RegisterSayHandler(kit.DefaultService.Server(),&logic.Handle{})

	kit.Run()
}
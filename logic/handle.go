package logic

import (
	"github.com/dayan-be/demo-service/proto/demo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Handle struct {

}

func (h *Handle) Hello(ctx context.Context, req *demo.HelloReq, rsp *demo.HelloRsp) error {
	logrus.Debugf("hello %s", req.Name)
	rsp.Msg = "my name is Jerry"
	return nil
}
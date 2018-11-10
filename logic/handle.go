package logic

import (
	"github.com/dayan-be/demo-service/proto"
	"golang.org/x/net/context"
)

type Handle struct {

}

func (h *Handle) Hello(ctx context.Context, req *demo.HelloReq, rsp *demo.HelloRsp) error {
	return nil
}
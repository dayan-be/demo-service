// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: demo.proto

/*
Package demo is a generated protocol buffer package.

It is generated from these files:
	demo.proto

It has these top-level messages:
	HelloReq
	HelloRsp
*/
package demo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Say service

type SayService interface {
	Hello(ctx context.Context, in *HelloReq, opts ...client.CallOption) (*HelloRsp, error)
}

type sayService struct {
	c    client.Client
	name string
}

func NewSayService(name string, c client.Client) SayService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "demo"
	}
	return &sayService{
		c:    c,
		name: name,
	}
}

func (c *sayService) Hello(ctx context.Context, in *HelloReq, opts ...client.CallOption) (*HelloRsp, error) {
	req := c.c.NewRequest(c.name, "Say.Hello", in)
	out := new(HelloRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayHandler interface {
	Hello(context.Context, *HelloReq, *HelloRsp) error
}

func RegisterSayHandler(s server.Server, hdlr SayHandler, opts ...server.HandlerOption) error {
	type say interface {
		Hello(ctx context.Context, in *HelloReq, out *HelloRsp) error
	}
	type Say struct {
		say
	}
	h := &sayHandler{hdlr}
	return s.Handle(s.NewHandler(&Say{h}, opts...))
}

type sayHandler struct {
	SayHandler
}

func (h *sayHandler) Hello(ctx context.Context, in *HelloReq, out *HelloRsp) error {
	return h.SayHandler.Hello(ctx, in, out)
}

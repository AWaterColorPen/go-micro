// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/examples/api/meta/proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/api/meta/proto/api.proto

It has these top-level messages:
	CallRequest
	CallResponse
	EmptyRequest
	EmptyResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for Example service

type ExampleService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type exampleService struct {
	c           client.Client
	serviceName string
}

func NewExampleService(serviceName string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "example"
	}
	return &exampleService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *exampleService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) {
	type example interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
	}
	type Example struct {
		example
	}
	h := &exampleHandler{hdlr}
	s.Handle(s.NewHandler(&Example{h}, opts...))
}

type exampleHandler struct {
	ExampleHandler
}

func (h *exampleHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.ExampleHandler.Call(ctx, in, out)
}

// Client API for Foo service

type FooService interface {
	Bar(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyResponse, error)
}

type fooService struct {
	c           client.Client
	serviceName string
}

func NewFooService(serviceName string, c client.Client) FooService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "foo"
	}
	return &fooService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *fooService) Bar(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Foo.Bar", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Foo service

type FooHandler interface {
	Bar(context.Context, *EmptyRequest, *EmptyResponse) error
}

func RegisterFooHandler(s server.Server, hdlr FooHandler, opts ...server.HandlerOption) {
	type foo interface {
		Bar(ctx context.Context, in *EmptyRequest, out *EmptyResponse) error
	}
	type Foo struct {
		foo
	}
	h := &fooHandler{hdlr}
	s.Handle(s.NewHandler(&Foo{h}, opts...))
}

type fooHandler struct {
	FooHandler
}

func (h *fooHandler) Bar(ctx context.Context, in *EmptyRequest, out *EmptyResponse) error {
	return h.FooHandler.Bar(ctx, in, out)
}

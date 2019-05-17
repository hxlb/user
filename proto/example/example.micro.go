// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/example/example.proto

package com_hxlb_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Example service

type ExampleService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Example_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Example_PingPongService, error)
}

type exampleService struct {
	c    client.Client
	name string
}

func NewExampleService(name string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "com.hxlb.srv.user"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Example.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Example_StreamService, error) {
	req := c.c.NewRequest(c.name, "Example.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &exampleServiceStream{stream}, nil
}

type Example_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type exampleServiceStream struct {
	stream client.Stream
}

func (x *exampleServiceStream) Close() error {
	return x.stream.Close()
}

func (x *exampleServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleService) PingPong(ctx context.Context, opts ...client.CallOption) (Example_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Example.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &exampleServicePingPong{stream}, nil
}

type Example_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type exampleServicePingPong struct {
	stream client.Stream
}

func (x *exampleServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *exampleServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *exampleServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Example_StreamStream) error
	PingPong(context.Context, Example_PingPongStream) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) error {
	type example interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Example struct {
		example
	}
	h := &exampleHandler{hdlr}
	return s.Handle(s.NewHandler(&Example{h}, opts...))
}

type exampleHandler struct {
	ExampleHandler
}

func (h *exampleHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.ExampleHandler.Call(ctx, in, out)
}

func (h *exampleHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ExampleHandler.Stream(ctx, m, &exampleStreamStream{stream})
}

type Example_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type exampleStreamStream struct {
	stream server.Stream
}

func (x *exampleStreamStream) Close() error {
	return x.stream.Close()
}

func (x *exampleStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *exampleHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.ExampleHandler.PingPong(ctx, &examplePingPongStream{stream})
}

type Example_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type examplePingPongStream struct {
	stream server.Stream
}

func (x *examplePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *examplePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *examplePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *examplePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *examplePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

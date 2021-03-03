// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/message/message.proto

package go_micro_srv_basic_message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Client API for Message service

type MessageService interface {
	Send(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*Response, error)
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...client.CallOption) (*Response, error)
	Init(ctx context.Context, in *InitRequest, opts ...client.CallOption) (*InitResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) Send(ctx context.Context, in *SendRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Message.Send", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Message.ChangeStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Init(ctx context.Context, in *InitRequest, opts ...client.CallOption) (*InitResponse, error) {
	req := c.c.NewRequest(c.name, "Message.Init", in)
	out := new(InitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Message.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageHandler interface {
	Send(context.Context, *SendRequest, *Response) error
	ChangeStatus(context.Context, *ChangeStatusRequest, *Response) error
	Init(context.Context, *InitRequest, *InitResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
}

func RegisterMessageHandler(s server.Server, hdlr MessageHandler, opts ...server.HandlerOption) error {
	type message interface {
		Send(ctx context.Context, in *SendRequest, out *Response) error
		ChangeStatus(ctx context.Context, in *ChangeStatusRequest, out *Response) error
		Init(ctx context.Context, in *InitRequest, out *InitResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type Message struct {
		message
	}
	h := &messageHandler{hdlr}
	return s.Handle(s.NewHandler(&Message{h}, opts...))
}

type messageHandler struct {
	MessageHandler
}

func (h *messageHandler) Send(ctx context.Context, in *SendRequest, out *Response) error {
	return h.MessageHandler.Send(ctx, in, out)
}

func (h *messageHandler) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, out *Response) error {
	return h.MessageHandler.ChangeStatus(ctx, in, out)
}

func (h *messageHandler) Init(ctx context.Context, in *InitRequest, out *InitResponse) error {
	return h.MessageHandler.Init(ctx, in, out)
}

func (h *messageHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.MessageHandler.Get(ctx, in, out)
}

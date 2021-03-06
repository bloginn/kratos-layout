// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type OpenHTTPServer interface {
	Hello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterOpenHTTPServer(s *http.Server, srv OpenHTTPServer) {
	r := s.Route("/")
	r.GET("/hello/{name}", _Open_Hello0_HTTP_Handler(srv))
}

func _Open_Hello0_HTTP_Handler(srv OpenHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.open.v1.Open/Hello")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Hello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type OpenHTTPClient interface {
	Hello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type OpenHTTPClientImpl struct {
	cc *http.Client
}

func NewOpenHTTPClient(client *http.Client) OpenHTTPClient {
	return &OpenHTTPClientImpl{client}
}

func (c *OpenHTTPClientImpl) Hello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/hello/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.open.v1.Open/Hello"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

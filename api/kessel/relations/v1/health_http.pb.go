// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             (unknown)
// source: kessel/relations/v1/health.proto

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

const OperationKesselHealthServiceGetLivez = "/kessel.relations.v1.KesselHealthService/GetLivez"
const OperationKesselHealthServiceGetReadyz = "/kessel.relations.v1.KesselHealthService/GetReadyz"

type KesselHealthServiceHTTPServer interface {
	GetLivez(context.Context, *GetLivezRequest) (*GetLivezResponse, error)
	GetReadyz(context.Context, *GetReadyzRequest) (*GetReadyzResponse, error)
}

func RegisterKesselHealthServiceHTTPServer(s *http.Server, srv KesselHealthServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/livez", _KesselHealthService_GetLivez0_HTTP_Handler(srv))
	r.GET("/readyz", _KesselHealthService_GetReadyz0_HTTP_Handler(srv))
}

func _KesselHealthService_GetLivez0_HTTP_Handler(srv KesselHealthServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLivezRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselHealthServiceGetLivez)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLivez(ctx, req.(*GetLivezRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLivezResponse)
		return ctx.Result(200, reply)
	}
}

func _KesselHealthService_GetReadyz0_HTTP_Handler(srv KesselHealthServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetReadyzRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselHealthServiceGetReadyz)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetReadyz(ctx, req.(*GetReadyzRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetReadyzResponse)
		return ctx.Result(200, reply)
	}
}

type KesselHealthServiceHTTPClient interface {
	GetLivez(ctx context.Context, req *GetLivezRequest, opts ...http.CallOption) (rsp *GetLivezResponse, err error)
	GetReadyz(ctx context.Context, req *GetReadyzRequest, opts ...http.CallOption) (rsp *GetReadyzResponse, err error)
}

type KesselHealthServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewKesselHealthServiceHTTPClient(client *http.Client) KesselHealthServiceHTTPClient {
	return &KesselHealthServiceHTTPClientImpl{client}
}

func (c *KesselHealthServiceHTTPClientImpl) GetLivez(ctx context.Context, in *GetLivezRequest, opts ...http.CallOption) (*GetLivezResponse, error) {
	var out GetLivezResponse
	pattern := "/livez"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKesselHealthServiceGetLivez))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *KesselHealthServiceHTTPClientImpl) GetReadyz(ctx context.Context, in *GetReadyzRequest, opts ...http.CallOption) (*GetReadyzResponse, error) {
	var out GetReadyzResponse
	pattern := "/readyz"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKesselHealthServiceGetReadyz))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.19.4
// source: server/prom/endpoint/endpoint.proto

package endpoint

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

const OperationEndpointAppendEndpoint = "/api.server.prom.endpoint.Endpoint/AppendEndpoint"
const OperationEndpointBatchEditEndpointStatus = "/api.server.prom.endpoint.Endpoint/BatchEditEndpointStatus"
const OperationEndpointDeleteEndpoint = "/api.server.prom.endpoint.Endpoint/DeleteEndpoint"
const OperationEndpointEditEndpoint = "/api.server.prom.endpoint.Endpoint/EditEndpoint"
const OperationEndpointGetEndpoint = "/api.server.prom.endpoint.Endpoint/GetEndpoint"
const OperationEndpointListEndpoint = "/api.server.prom.endpoint.Endpoint/ListEndpoint"
const OperationEndpointSelectEndpoint = "/api.server.prom.endpoint.Endpoint/SelectEndpoint"

type EndpointHTTPServer interface {
	// AppendEndpoint 添加监控端点
	AppendEndpoint(context.Context, *AppendEndpointRequest) (*AppendEndpointReply, error)
	// BatchEditEndpointStatus 批量修改端点状态
	BatchEditEndpointStatus(context.Context, *BatchEditEndpointStatusRequest) (*BatchEditEndpointStatusReply, error)
	// DeleteEndpoint 删除监控端点
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*DeleteEndpointReply, error)
	// EditEndpoint 编辑端点信息
	EditEndpoint(context.Context, *EditEndpointRequest) (*EditEndpointReply, error)
	// GetEndpoint 获取监控端点详情
	GetEndpoint(context.Context, *GetEndpointRequest) (*GetEndpointReply, error)
	// ListEndpoint 获取监控端点列表
	ListEndpoint(context.Context, *ListEndpointRequest) (*ListEndpointReply, error)
	// SelectEndpoint 获取监控端点下拉列表
	SelectEndpoint(context.Context, *SelectEndpointRequest) (*SelectEndpointReply, error)
}

func RegisterEndpointHTTPServer(s *http.Server, srv EndpointHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/endpoint/append", _Endpoint_AppendEndpoint0_HTTP_Handler(srv))
	r.POST("/api/v1/endpoint/delete", _Endpoint_DeleteEndpoint0_HTTP_Handler(srv))
	r.POST("/api/v1/endpoint/edit", _Endpoint_EditEndpoint0_HTTP_Handler(srv))
	r.POST("/api/v1/endpoint/batch/status", _Endpoint_BatchEditEndpointStatus0_HTTP_Handler(srv))
	r.POST("/api/v1/endpoint/detail", _Endpoint_GetEndpoint0_HTTP_Handler(srv))
	r.POST("/api/v1/endpoint/list", _Endpoint_ListEndpoint0_HTTP_Handler(srv))
	r.POST("/api/v1/endpoint/select", _Endpoint_SelectEndpoint0_HTTP_Handler(srv))
}

func _Endpoint_AppendEndpoint0_HTTP_Handler(srv EndpointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppendEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEndpointAppendEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppendEndpoint(ctx, req.(*AppendEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppendEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Endpoint_DeleteEndpoint0_HTTP_Handler(srv EndpointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEndpointDeleteEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteEndpoint(ctx, req.(*DeleteEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Endpoint_EditEndpoint0_HTTP_Handler(srv EndpointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EditEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEndpointEditEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EditEndpoint(ctx, req.(*EditEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EditEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Endpoint_BatchEditEndpointStatus0_HTTP_Handler(srv EndpointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BatchEditEndpointStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEndpointBatchEditEndpointStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchEditEndpointStatus(ctx, req.(*BatchEditEndpointStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BatchEditEndpointStatusReply)
		return ctx.Result(200, reply)
	}
}

func _Endpoint_GetEndpoint0_HTTP_Handler(srv EndpointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEndpointGetEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEndpoint(ctx, req.(*GetEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Endpoint_ListEndpoint0_HTTP_Handler(srv EndpointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEndpointListEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListEndpoint(ctx, req.(*ListEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListEndpointReply)
		return ctx.Result(200, reply)
	}
}

func _Endpoint_SelectEndpoint0_HTTP_Handler(srv EndpointHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SelectEndpointRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEndpointSelectEndpoint)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SelectEndpoint(ctx, req.(*SelectEndpointRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SelectEndpointReply)
		return ctx.Result(200, reply)
	}
}

type EndpointHTTPClient interface {
	AppendEndpoint(ctx context.Context, req *AppendEndpointRequest, opts ...http.CallOption) (rsp *AppendEndpointReply, err error)
	BatchEditEndpointStatus(ctx context.Context, req *BatchEditEndpointStatusRequest, opts ...http.CallOption) (rsp *BatchEditEndpointStatusReply, err error)
	DeleteEndpoint(ctx context.Context, req *DeleteEndpointRequest, opts ...http.CallOption) (rsp *DeleteEndpointReply, err error)
	EditEndpoint(ctx context.Context, req *EditEndpointRequest, opts ...http.CallOption) (rsp *EditEndpointReply, err error)
	GetEndpoint(ctx context.Context, req *GetEndpointRequest, opts ...http.CallOption) (rsp *GetEndpointReply, err error)
	ListEndpoint(ctx context.Context, req *ListEndpointRequest, opts ...http.CallOption) (rsp *ListEndpointReply, err error)
	SelectEndpoint(ctx context.Context, req *SelectEndpointRequest, opts ...http.CallOption) (rsp *SelectEndpointReply, err error)
}

type EndpointHTTPClientImpl struct {
	cc *http.Client
}

func NewEndpointHTTPClient(client *http.Client) EndpointHTTPClient {
	return &EndpointHTTPClientImpl{client}
}

func (c *EndpointHTTPClientImpl) AppendEndpoint(ctx context.Context, in *AppendEndpointRequest, opts ...http.CallOption) (*AppendEndpointReply, error) {
	var out AppendEndpointReply
	pattern := "/api/v1/endpoint/append"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEndpointAppendEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EndpointHTTPClientImpl) BatchEditEndpointStatus(ctx context.Context, in *BatchEditEndpointStatusRequest, opts ...http.CallOption) (*BatchEditEndpointStatusReply, error) {
	var out BatchEditEndpointStatusReply
	pattern := "/api/v1/endpoint/batch/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEndpointBatchEditEndpointStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EndpointHTTPClientImpl) DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...http.CallOption) (*DeleteEndpointReply, error) {
	var out DeleteEndpointReply
	pattern := "/api/v1/endpoint/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEndpointDeleteEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EndpointHTTPClientImpl) EditEndpoint(ctx context.Context, in *EditEndpointRequest, opts ...http.CallOption) (*EditEndpointReply, error) {
	var out EditEndpointReply
	pattern := "/api/v1/endpoint/edit"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEndpointEditEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EndpointHTTPClientImpl) GetEndpoint(ctx context.Context, in *GetEndpointRequest, opts ...http.CallOption) (*GetEndpointReply, error) {
	var out GetEndpointReply
	pattern := "/api/v1/endpoint/detail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEndpointGetEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EndpointHTTPClientImpl) ListEndpoint(ctx context.Context, in *ListEndpointRequest, opts ...http.CallOption) (*ListEndpointReply, error) {
	var out ListEndpointReply
	pattern := "/api/v1/endpoint/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEndpointListEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EndpointHTTPClientImpl) SelectEndpoint(ctx context.Context, in *SelectEndpointRequest, opts ...http.CallOption) (*SelectEndpointReply, error) {
	var out SelectEndpointReply
	pattern := "/api/v1/endpoint/select"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEndpointSelectEndpoint))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

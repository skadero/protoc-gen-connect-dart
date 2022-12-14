// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: haberdasher/haberdasher.proto

package haberdasherconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	haberdasher "github.com/skadero/protoc-gen-connect-dart/example/go_server/rpc/haberdasher"
	pinger "github.com/skadero/protoc-gen-connect-dart/example/go_server/rpc/pinger"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// HaberdasherName is the fully-qualified name of the Haberdasher service.
	HaberdasherName = "haberdasher.Haberdasher"
)

// HaberdasherClient is a client for the haberdasher.Haberdasher service.
type HaberdasherClient interface {
	// MakeHat produces a hat of mysterious, randomly-selected color!
	MakeHat(context.Context, *connect_go.Request[haberdasher.Size]) (*connect_go.Response[haberdasher.Hat], error)
	// Ping let's see if  the haberdasher is ready
	Ping(context.Context, *connect_go.Request[pinger.PingRequest]) (*connect_go.Response[pinger.PingResponse], error)
	// InvalidArg always returns an invalid argument error
	InvalidArg(context.Context, *connect_go.Request[pinger.PingRequest]) (*connect_go.Response[pinger.PingResponse], error)
}

// NewHaberdasherClient constructs a client for the haberdasher.Haberdasher service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHaberdasherClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HaberdasherClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &haberdasherClient{
		makeHat: connect_go.NewClient[haberdasher.Size, haberdasher.Hat](
			httpClient,
			baseURL+"/haberdasher.Haberdasher/MakeHat",
			opts...,
		),
		ping: connect_go.NewClient[pinger.PingRequest, pinger.PingResponse](
			httpClient,
			baseURL+"/haberdasher.Haberdasher/Ping",
			opts...,
		),
		invalidArg: connect_go.NewClient[pinger.PingRequest, pinger.PingResponse](
			httpClient,
			baseURL+"/haberdasher.Haberdasher/InvalidArg",
			opts...,
		),
	}
}

// haberdasherClient implements HaberdasherClient.
type haberdasherClient struct {
	makeHat    *connect_go.Client[haberdasher.Size, haberdasher.Hat]
	ping       *connect_go.Client[pinger.PingRequest, pinger.PingResponse]
	invalidArg *connect_go.Client[pinger.PingRequest, pinger.PingResponse]
}

// MakeHat calls haberdasher.Haberdasher.MakeHat.
func (c *haberdasherClient) MakeHat(ctx context.Context, req *connect_go.Request[haberdasher.Size]) (*connect_go.Response[haberdasher.Hat], error) {
	return c.makeHat.CallUnary(ctx, req)
}

// Ping calls haberdasher.Haberdasher.Ping.
func (c *haberdasherClient) Ping(ctx context.Context, req *connect_go.Request[pinger.PingRequest]) (*connect_go.Response[pinger.PingResponse], error) {
	return c.ping.CallUnary(ctx, req)
}

// InvalidArg calls haberdasher.Haberdasher.InvalidArg.
func (c *haberdasherClient) InvalidArg(ctx context.Context, req *connect_go.Request[pinger.PingRequest]) (*connect_go.Response[pinger.PingResponse], error) {
	return c.invalidArg.CallUnary(ctx, req)
}

// HaberdasherHandler is an implementation of the haberdasher.Haberdasher service.
type HaberdasherHandler interface {
	// MakeHat produces a hat of mysterious, randomly-selected color!
	MakeHat(context.Context, *connect_go.Request[haberdasher.Size]) (*connect_go.Response[haberdasher.Hat], error)
	// Ping let's see if  the haberdasher is ready
	Ping(context.Context, *connect_go.Request[pinger.PingRequest]) (*connect_go.Response[pinger.PingResponse], error)
	// InvalidArg always returns an invalid argument error
	InvalidArg(context.Context, *connect_go.Request[pinger.PingRequest]) (*connect_go.Response[pinger.PingResponse], error)
}

// NewHaberdasherHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHaberdasherHandler(svc HaberdasherHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/haberdasher.Haberdasher/MakeHat", connect_go.NewUnaryHandler(
		"/haberdasher.Haberdasher/MakeHat",
		svc.MakeHat,
		opts...,
	))
	mux.Handle("/haberdasher.Haberdasher/Ping", connect_go.NewUnaryHandler(
		"/haberdasher.Haberdasher/Ping",
		svc.Ping,
		opts...,
	))
	mux.Handle("/haberdasher.Haberdasher/InvalidArg", connect_go.NewUnaryHandler(
		"/haberdasher.Haberdasher/InvalidArg",
		svc.InvalidArg,
		opts...,
	))
	return "/haberdasher.Haberdasher/", mux
}

// UnimplementedHaberdasherHandler returns CodeUnimplemented from all methods.
type UnimplementedHaberdasherHandler struct{}

func (UnimplementedHaberdasherHandler) MakeHat(context.Context, *connect_go.Request[haberdasher.Size]) (*connect_go.Response[haberdasher.Hat], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("haberdasher.Haberdasher.MakeHat is not implemented"))
}

func (UnimplementedHaberdasherHandler) Ping(context.Context, *connect_go.Request[pinger.PingRequest]) (*connect_go.Response[pinger.PingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("haberdasher.Haberdasher.Ping is not implemented"))
}

func (UnimplementedHaberdasherHandler) InvalidArg(context.Context, *connect_go.Request[pinger.PingRequest]) (*connect_go.Response[pinger.PingResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("haberdasher.Haberdasher.InvalidArg is not implemented"))
}

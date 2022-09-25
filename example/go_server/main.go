package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/bufbuild/connect-go"
	haberdasher "github.com/skadero/protoc-gen-connect-dart/example/go_server/rpc/haberdasher"
	"github.com/skadero/protoc-gen-connect-dart/example/go_server/rpc/haberdasher/haberdasherconnect"
	"github.com/skadero/protoc-gen-connect-dart/example/go_server/rpc/pinger"
)

type haberdasherServer struct{}

func (s *haberdasherServer) MakeHat(ctx context.Context, req *connect.Request[haberdasher.Size]) (*connect.Response[haberdasher.Hat], error) {
	log.Println("Make Hat")
	res := connect.NewResponse(&haberdasher.Hat{Size: req.Msg.Inches, Color: "green"})
	return res, nil
}

func (s *haberdasherServer) Ping(context.Context, *connect.Request[pinger.PingRequest]) (*connect.Response[pinger.PingResponse], error) {
	log.Println("Ping")
	return connect.NewResponse(&pinger.PingResponse{Answer: "pong"}), nil
}

func (s *haberdasherServer) InvalidArg(context.Context, *connect.Request[pinger.PingRequest]) (*connect.Response[pinger.PingResponse], error) {
	log.Println("Invalid arg error")
	return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid ping"))
}

func main() {
	haber := &haberdasherServer{}
	mux := http.NewServeMux()
	path, handler := haberdasherconnect.NewHaberdasherHandler(haber)
	mux.Handle(path, handler)
	log.Println("Listening...")
	http.ListenAndServe(
		"localhost:9080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

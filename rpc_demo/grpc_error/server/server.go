package main

import (
	"context"
	"hello_go/rpc_demo/grpc_test/proto"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	*proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	time.Sleep(5 * time.Second)
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, status.Errorf(codes.NotFound, "not find: %s", request.Name)
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic("filed to listen: " + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("filed to statr grpc: " + err.Error())
	}
}

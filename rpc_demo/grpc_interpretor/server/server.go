package main

import (
	"context"
	"fmt"
	"hello_go/rpc_demo/grpc_test/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	*proto.UnimplementedGreeterServer
}


func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error){
		fmt.Println("recived an new request")
		res, err := handler(ctx, req)
		return res, err

	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
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

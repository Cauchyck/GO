package main

import (
	"context"
	"fmt"
	"hello_go/rpc_demo/grpc_test/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("recived an new request")

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {

			fmt.Println("get metadata error")
			return resp, status.Error(codes.Unauthenticated, "No Token")
		}


		var (
			appid string
			appkey string
		)
		if var1, ok := md["appid"]; ok{
			appid = var1[0]
		}
		if var1, ok := md["appkey"]; ok{
			appkey = var1[0]
		}

		if appid != "bobby"{
			return resp, status.Error(codes.Unauthenticated, "appid  error")
		}
		if appkey != "imooc"{
			return resp, status.Error(codes.Unauthenticated, "appkey error")
		}


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

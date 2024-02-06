package main

import (
	"context"
	"fmt"
	"hello_go/rpc_demo/grpc_test/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	*proto.UnimplementedGreeterServer
}


func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}
	for key, val := range md{
		fmt.Println(key, val)
	}
	if name, ok := md["name"];   ok{
		fmt.Println(name)
		for i, e := range name{
			fmt.Println(i, e)
		}
	}
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
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

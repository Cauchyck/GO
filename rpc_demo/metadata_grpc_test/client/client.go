package main

import (
	"context"
	"fmt"
	"hello_go/rpc_demo/grpc_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() { 
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	md := metadata.New(map[string]string{
		"name": "bobby",
		"passwd": "imooc",
	})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "bobby",
		Url: "google.com",
		G:   proto.Gender_FEMALE,
		Mp: map[string]string{
			"name": "bobby",
		},
		AddTime: timestamppb.New(time.Now()),})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)

}

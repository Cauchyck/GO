package main

import (
	"context"
	"fmt"
	"hello_go/rpc_demo/grpc_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() { 
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	ctx , _:= context.WithTimeout(context.Background(), time.Second*3) 
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "bobby",
		Url: "google.com",
		G:   proto.Gender_FEMALE,
		Mp: map[string]string{
			"name": "bobby",
		},
		AddTime: timestamppb.New(time.Now()),})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("faild")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	fmt.Println(r.Message)

}

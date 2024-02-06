package main

import (
	"context"
	"fmt"
	"hello_go/rpc_demo/grpc_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	interceptor :=  func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("time : %s", time.Since(start))

		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure(), opt)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "bobby",
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

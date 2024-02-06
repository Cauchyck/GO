package main

import (
	"context"
	"fmt"
	"hello_go/rpc_demo/grpc_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type customCredential struct{

}
func (c customCredential)GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error){
	// RequireTransportSecurity indicates whether the credentials requires
	// transport security.
	return map[string]string{
			"appid":   "bobby",
			"appkey": "imooc",
	}, nil
}
	
func (c customCredential)RequireTransportSecurity() bool{
	return false
}
func main() {
	// interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 	start := time.Now()

	// 	md := metadata.New(map[string]string{
	// 		"appid":   "bobby",
	// 		"appkey": "imooc",
	// 	})
	// 	ctx = metadata.NewOutgoingContext(context.Background(), md)

	// 	err := invoker(ctx, method, req, reply, cc, opts...)
	// 	fmt.Printf("time : %s\n", time.Since(start))

	// 	return err
	// }
	opt := grpc.WithPerRPCCredentials(customCredential{})

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
		AddTime: timestamppb.New(time.Now())})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)

}

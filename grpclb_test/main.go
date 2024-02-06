package main

import (
	"context"
	"fmt"
	"hello_go/grpclb_test/proto"
	"log"

	_ "github.com/mbobakov/grpc-consul-resolver"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"consul://127.0.0.1:8500/user_srv?wait=14s",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()



	userClient := proto.NewUserClient(conn)
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn: 1,
		PSize: 2,
	})
	if err != nil {
		panic(err)
	}

	for index, data := range rsp.Data{
		fmt.Println(index, data)
	}
}

package main

import (
	"context"
	"fmt"
	"hello_go/rpc_demo/stream_grpc_test/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5678", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "protp_StreamReqData"})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)

	}

	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("proto.StreamData %d", i),
		})

		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("recive from server: " + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allStr.Send(&proto.StreamReqData{Data: "message from client"})
			time.Sleep(time.Second)

		}
	}()
	wg.Wait()

}

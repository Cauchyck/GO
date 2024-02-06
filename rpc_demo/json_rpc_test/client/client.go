package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 建立连接
	conn, err := net.Dial("tcp", "localhost:1234")

	if err != nil {
		panic(err)
	}

	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("failed")
	}
	fmt.Println(reply)

}

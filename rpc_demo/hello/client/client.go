package main

import (
	"fmt"
	"net/rpc"
)


func main() {
	// 建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		panic(err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("failed")
	}
	fmt.Println(reply)

}
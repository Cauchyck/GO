package main

import (
	"fmt"
	clientproxy "hello_go/rpc_demo/new_hello/client_proxy"
)

func main() {
	// 建立连接
	client := clientproxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply string
	err := client.Hello("booby", &reply)
	if err != nil {
		panic("Hello failed")
	}
	fmt.Println(reply)

}

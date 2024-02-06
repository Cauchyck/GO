package main

import (
	"hello_go/rpc_demo/new_hello/handler"
	serverproxy "hello_go/rpc_demo/new_hello/server_proxy"
	"net"
	"net/rpc"
)


func main(){
	// 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 注册处理逻辑handler
	_ =    serverproxy.RegisterHelloService(&handler.HelloServer{})
	// 启动服务
	for{
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}

}
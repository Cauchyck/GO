package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer struct{

}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "Hello, "+request
	return nil
}

func main(){
	// 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloServer{})
	// 启动服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}

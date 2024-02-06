package serverproxy

import (
	"hello_go/rpc_demo/new_hello/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}
func RegisterHelloService(sev HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, sev)
}

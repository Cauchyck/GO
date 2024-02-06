package clientproxy

import (
	"hello_go/rpc_demo/new_hello/handler"
	"net/rpc"
)

type HelloServiceStub struct{
	*rpc.Client
}

func NewHelloServiceClient (protol, address string) HelloServiceStub{
	conn, err := rpc.Dial(protol, address)
	if err != nil{
		panic("connect error!")
	}

	return HelloServiceStub{conn}

}

func (c *HelloServiceStub)Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}

	return nil

}
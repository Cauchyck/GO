package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServer struct {
}

func (s *HelloServer) Hello(request string, reply *string) error {
	*reply = "Hello, " + request
	return nil
}

func main() {
	// 实例化一个server
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	// 注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloServer{})
	// 启动服务
	http.ListenAndServe(":1234", nil)

}

package main

import (
	"fmt"
	"hello_go/rpc_demo/stream_grpc_test/proto"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":5678"

type Server struct {
	*proto.UnimplementedGreeterServer
}

func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})

		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *Server) PutStream(cliSrt proto.Greeter_PutStreamServer) error {
	for {
		a, err := cliSrt.Recv();
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}
	return nil


}

func (s *Server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		defer wg.Done()
		for{
			data, _ := allStr.Recv()
			fmt.Println("recive from client: " + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allStr.Send(&proto.StreamResData{Data: "message from server"})
			time.Sleep(time.Second)
			
		}
	}()
	wg.Wait()
	return nil
}

func main() {

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic("filed to listen: " + err.Error())
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		panic("filed to statr grpc: " + err.Error())
	}
}

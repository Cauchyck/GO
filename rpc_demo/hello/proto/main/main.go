package main

import (
	"encoding/json"
	"fmt"
	helloworld "hello_go/rpc_demo/hello/proto"

	"google.golang.org/protobuf/proto"
)

type Hello struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Courses []string `json:"coursers"`
}

func main(){
	req := helloworld.HelloRequest{
		Name: "bobby",
		Age: 18,
		Courses: []string{"go", "gin"},
	}
	rsp, _ :=proto.Marshal(&req)
	fmt.Println(len(rsp))

	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name)

	jsonStruct := Hello{
		Name: "bobby",
		Age: 18,
		Courses: []string{"go", "gin"},
	}
	jsonRsp, _ := json.Marshal(jsonStruct)
	fmt.Println(len(jsonRsp))



}
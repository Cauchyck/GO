package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func HandleError(err error){
	fmt.Println("error: ", err)
	os.Exit(-1)
}

func main() {
  fmt.Println("OSS Go SDK Version: ", oss.Version)

  client, err := oss.New("http://oss-cn-shanghai.aliyuncs.com", "LTAI5t7fj52mUUWSmLZRtZ59", "U92z3L37nABteJjgiDSD0x4PepsKu9")
    if err != nil {
        HandleError(err)
    }
    
    bucket, err := client.Bucket("mxshop-go")
    if err != nil {
        HandleError(err)
    }
    
    err = bucket.PutObjectFromFile("goods/test.png", "oss_test/OM3xxP4B3l.png")
    if err != nil {
        HandleError(err)
    }

}
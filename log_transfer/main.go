package main

import (
	"fmt"
	"hello_go/log_transfer/es"
	"hello_go/log_transfer/kafka"
	"hello_go/log_transfer/model"

	"github.com/go-ini/ini"
)

func main() {
	var cfg = new(model.Config)
	err := ini.MapTo(cfg, "config/logtransfer.ini")
	if err != nil {
		fmt.Printf("load config failed, err: %v \n", err)
		panic(err)
	}

	fmt.Println("Load config success")

	// 初始化Es
	err = es.Init(cfg.ESConf.Index, cfg.ESConf.Address, cfg.ESConf.GoroutineNum, cfg.ESConf.MaxSize)
	if err != nil {
		fmt.Printf("Connect es failed, err: %v \n", err)
		panic(err)
	}
	fmt.Println("Connect es success")

	// 初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Printf("Connect kafka failed, err: %v \n", err)
		panic(err)
	}

	fmt.Println("Connect kafka success")

	select {}

}

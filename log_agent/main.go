package main

import (
	"fmt"
	"hello_go/log_agent/common"
	"hello_go/log_agent/etcd"
	"hello_go/log_agent/kafka"
	tailfile "hello_go/log_agent/tail_file"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

// 日志收集客户端
// 类似的开源项目：filebeat

// 收集指定目录下的日志文件，发送到kafka中

// 整个log agent的配置结构体
type Config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
	EtcdConfig	`ini:"etcd"`
}
type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
	chanSize int64	`ini:"chan_size"`
}
type EtcdConfig struct{
	Address string `ini:"address"`
	CollectKey string `ini:"collect_key"`
	
}
type CollectConfig struct {
	LogFilePath string `ini:"logfile_path"`
}

func run(){

	select {

	}

}

func main() {
	// -1
	ip, err := common.GetOutboundIP()
	if err != nil {
		logrus.Errorf("get ip failed, err: %v", err)
		return
	}

	var configObj = new(Config)
	// 1. 读配置文件： ”go ini“
	// cfg, err := ini.Load("log_agent/config.ini")
	// if err != nil {
	// 	logrus.Error("Load config fialed, err: %v", err)
	// 	return
	// }
	// kafkaAddr := cfg.Section("kafka").Key("address").String()
	// fmt.Println(kafkaAddr)
	err = ini.MapTo(configObj, "./config.ini")
	if err != nil {
		logrus.Errorf("Load config fialed, err: %v", err)
		return
	}
	fmt.Printf("%#v\n", configObj)

	// 2. 连接kafka
	err = kafka.Init([]string{configObj.KafkaConfig.Address}, configObj.KafkaConfig.chanSize)
	if err != nil {
		logrus.Errorf("Load config fialed, err: %v", err)
		return
	}
	logrus.Info("Init kafka success!!!")

	// 初始化etcd连接， 从ectd中拉取要收集日志的配置项
	// 初始化
	err = etcd.Init([]string{configObj.EtcdConfig.Address})
	if err != nil {
		logrus.Errorf("init etcd fialed, err: %v", err)
		return
	}
	logrus.Info("Init etcd success!!!")

	// 拉取配置项
	collectKey := fmt.Sprintf(configObj.EtcdConfig.CollectKey, ip)
	// fmt.Println("configObj.EtcdConfig.CollectKey:", configObj.EtcdConfig.CollectKey)
	allConf, err := etcd.GetConf(collectKey)
	if err != nil {
		logrus.Errorf("get conf from etcd fialed, err: %v", err)
		return
	}
	fmt.Println(allConf)
	// 监测configObj.EtcdConfig.CollectKey对应值的变化
	go etcd.WatchConf(collectKey)
	//3. 初始化tail
	err = tailfile.Init(allConf)
	if err != nil {
		logrus.Errorf("init etcd fialed, err: %v", err)
		return
	}
	logrus.Info("Init tail success!!!")

	// 4. 将日志通过sarama发往kafka
	// logfile --> TailObj --> log --> Client --> kafka
	run()


}

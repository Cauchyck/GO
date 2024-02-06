package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

// kafka和nsq的区别
// nsq: 更多做消息队列
// kafka: 比较重量级的兼顾存储和消息队列

// kafka client demo

 func main(){
	// 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true //确认

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err: ", err)
		return
	}
	defer client.Close()

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("test log")

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("producer closed, err: ", err)
		return
	}
	fmt.Printf("pid: %v, offset: %v \n", pid, offset)

 }
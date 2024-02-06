package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

// kafka消费者

func main(){
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition, err:%v\n", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList{
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d, err:%v\n", partition, err)
			return
		} 
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer){
			for msg := range pc.Messages(){
				fmt.Printf("Partition:%d Offset: %d Key: %v Value: %s", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}
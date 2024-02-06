package kafka

import (
	"encoding/json"
	"fmt"
	"hello_go/log_transfer/es"

	"github.com/IBM/sarama"
)

// 初始化kafka连接

func Init(address []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list of partition, err:%v\n", err)
		return
	}
	// fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d, err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset: %d Key: %v Value: %s", msg.Partition, msg.Offset, msg.Key, msg.Value)
				var m1 map[string]interface{}
				err = json.Unmarshal(msg.Value, &m1)
				if err != nil {
					fmt.Printf("unmarshal msg failed, err: %v", err)
					continue
				}
				es.PutLogData(m1)
			}
		}(pc) 
	}
	return
}

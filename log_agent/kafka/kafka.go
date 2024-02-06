package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

var (
	client sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage
)

// 初始化全局的kafka Client
func Init(address []string, chanSize int64)(err error){
	// 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true //确认

	// 连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		logrus.Error("producer closed, err: ", err)
		return
	}
	// 初始化MsgChan
	msgChan = make(chan *sarama.ProducerMessage, chanSize)
	// defer client.Close()
	// 起一个后台的goroutine从msgchan中读数据
	go sendMsg()
	return 
}

// 从msgChan中读取msg发送到kafka
func sendMsg(){
	for {
		select{
		case msg := <- msgChan:
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				logrus.Warning("send msg failed, err:", err)
				return
			}
			logrus.Infof("send msg to kafka success. pid:%v offset:%v", pid, offset)
		}
	}

}

func ToMsgChan(msg *sarama.ProducerMessage){
	msgChan <- msg

}
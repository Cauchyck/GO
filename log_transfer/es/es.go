package es

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type ESClient struct {
	client      *elastic.Client
	index       string
	logDataChan chan interface{}
}

type Person struct {
	Name    string
	Age     int
	Married bool
}

var (
	esClient *ESClient
)

// 将日志数据写入ES

func Init(index, address string, goroutineNum, maxSize int) (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + address))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")
	esClient = &ESClient{
		client:      client,
		index:       index,
		logDataChan: make(chan interface{}, maxSize),
	}
	// p1 := Person{Name:"guan", Age:18, Married: true}
	for i := 0; i < goroutineNum; i++ {
		go SendToES()
	}

	return
}

func SendToES() {
	for m1 := range esClient.logDataChan {
		put1, err := esClient.client.Index().Index(esClient.index).BodyJson(m1).Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Indexed user %s to index %s, type %s \n", put1.Id, put1.Index, put1.Type)

	}

}
func PutLogData(msg interface{}) {
	esClient.logDataChan <- msg
}

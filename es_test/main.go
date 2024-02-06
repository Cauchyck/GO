package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

type Account struct {
	AccountNumber int32  `json:"account_number"`
	FirstName     string `json:"firstname"`
}

const goodsMapping = `
{
	"mappings": {
	  "properties": {
		"mygoods": {
		  "properties": {
			"name": {
			  "type": "text",
			  "analyzer": "ik_max_word"
			},
			"id": {
			  "type": "integer"
			}
		  }
		}
	  }
	}
}
`

func main() {

	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	q := elastic.NewMatchQuery("address", "street")
	// src, err := q.Source()
	// if err != nil {
	// 	panic(err)
	// }

	// data, err := json.Marshal(src)
	// got := string(data)
	// fmt.Println(got)

	result, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	total := result.Hits.TotalHits.Value
	fmt.Println("搜索数量：", total)
	for _, value := range result.Hits.Hits {
		account := Account{}
		json.Unmarshal(value.Source, &account)
		// if jsonData, err := value.Source.MarshalJSON(); err == nil {
		// 	fmt.Println(string(jsonData))
		// } else {
		// 	panic(err)
		// }
		fmt.Println(account)
	}

	// account := Account{
	// 	AccountNumber: 12234,
	// 	FirstName: "bobby",
	// }
	// put, err := client.Index().Index("myuesr").BodyJson(account).Do(context.Background())
	// if err != nil{
	// 	panic(err)
	// }

	// fmt.Println(put.Id, put.Index, put.Type)
	createIndex, err := client.CreateIndex("mygoods").BodyString(goodsMapping).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !createIndex.Acknowledged{}
}

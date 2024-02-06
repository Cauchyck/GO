package main

import (
	"encoding/json"
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/sirupsen/logrus"
)

// connect
const (
	MyDB          = "mydb"
	username      = "admin"
	password      = ""
	MyMeasurement = "cpu_usage"
)

func connInflux() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		logrus.Error(err)
	}
	return cli
}

func queryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "mydb",
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// Insert
func WritesPoints(cli client.Client) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		logrus.Fatal(err)
	}

	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   20.1,
		"system": 43.3,
		"user":   86.6,
	}
	pt, err := client.NewPoint(
		"cpu_usage",
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	bp.AddPoint(pt)
	if err := cli.Write(bp); err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("Insert success")
}

func main() {
	conn := connInflux()
	fmt.Println(conn)

	WritesPoints(conn)

	// 获取数据并展示
	qs := fmt.Sprintf("select * from %s LIMIT %d", MyMeasurement, 10)
	res, err := queryDB(conn, qs)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println("===========", res)
	for i, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			logrus.Fatal(err)
		}
		value := row[2].(json.Number)
		logrus.Printf("[%2d] %s: %s ", i, t.Format(time.Stamp), value)
	}
}

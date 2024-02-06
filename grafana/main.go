package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/sirupsen/logrus"

	client "github.com/influxdata/influxdb1-client/v2"
)
var (
	cli client.Client

)
func initConnInflux()(err error){
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://127.0.0.1:8086",
		Username: "admin",
		Password: "",
	})
	return
}

func writesPoints(percent int64){
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: "monitor",
		Precision: "s",
	})
	if err != nil {
		logrus.Fatal(err)
	}

	tags := map[string]string{"cpu": "cpu0"}
	fields := map[string]interface{}{
		"cpu_percent": percent,
	}

	pt, err := client.NewPoint("cpu_percent", tags, fields, time.Now())
	if err != nil{
		logrus.Fatal(err)
	}
	bp.AddPoint(pt)

	err = cli.Write(bp)
	if err != nil{
		logrus.Fatal(err)
	}
	fmt.Println("insert success")
}
func getCpuInfo(){
	//CPU Usage
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("CPU Usage percent: %v \n", percent)
	// write into influxdb
	writesPoints(int64(percent[0]))
}

func main(){
	err := initConnInflux()
	if err != nil{
		fmt.Printf("Connect to influxdb failedm err: %v", err)
		return
	}
	for{
		getCpuInfo()
		time.Sleep(time.Second)
	}
	

}
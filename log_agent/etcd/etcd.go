package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"hello_go/log_agent/common"
	tailfile "hello_go/log_agent/tail_file"
	"time"

	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

var(
	client *clientv3.Client


)

func Init(address []string)(err error){
	client, err = clientv3.New(clientv3.Config{
		Endpoints: address,
		DialTimeout: time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd faild, err: %v", err)
		return
	}
	return

}

// 拉取日志项
func GetConf(key string)(collectEntryList []common.ClloectEntry, err error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		logrus.Errorf("get conf from etcd by key: %s faild, err: %v", key, err)
		return
	}
	if len(resp.Kvs) == 0 {
		logrus.Warningf("get len: 0 conf from etcd by key: %s", key)
	}
	ret := resp.Kvs[0]
	err = json.Unmarshal(ret.Value, &collectEntryList)
	if err != nil {
		logrus.Errorf("json unmarshal failed, err: %v", err)
		return
	}

	return 
}

// 监控configObj.EtcdConfig.CollectKey配置变化
func WatchConf(collectKey string){
	for{
		watchChan := client.Watch(context.Background(), collectKey)

		for wresp := range watchChan{

			for _, evt := range wresp.Events{
				fmt.Printf("type:%s key:%s value: %s \n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				var newConf []common.ClloectEntry
				if evt.Type == clientv3.EventTypeDelete{
					logrus.Warning("FBI warning: etcd delete the key!!!")
					tailfile.SendNewConf(newConf)
					continue
				}
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil{
					logrus.Errorf("json unmarshal new conf failed, err: %v", err)
					continue
				} 
				tailfile.SendNewConf(newConf)
			}
		}
	
	}
}
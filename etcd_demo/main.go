package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//代码连接etcd

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		DialTimeout: time.Second*5,
	})

	if err != nil{
		fmt.Printf("connect to etcd failed, err: %v", err)
		return
	}
	defer cli.Close()
	// Put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "collect_log_58.199.161.0_conf", `[{"path":"/Volumes/KIOXIARC20/GO/tailf_demo/logfile.log","topic":"db_log"}]`)
	if err != nil {
		fmt.Printf("Put key to etcd failed, err: %v", err)
		return
	}
	cancel()
	// Get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	gr, err := cli.Get(ctx, "collect_log_conf")
	if err != nil {
		fmt.Printf("Get key from etcd failed, err: %v", err)
		return
	}
	for _, ev := range gr.Kvs{
		fmt.Printf("key: %s, value: %s\n", ev.Key, ev.Value)
	}
	
	// Watch 监测etcd
	watchChan := cli.Watch(context.Background(), "key")
	
	for wresp := range watchChan{
		for _, evt := range wresp.Events{
			fmt.Printf("type: %s, key: %s, value: %s", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}
	



}

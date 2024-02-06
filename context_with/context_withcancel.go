package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 全局变量方式
// var eixt bool  = false
// 使用channel的方式实现


var wg sync.WaitGroup

func worker(ctx context.Context){
	defer wg.Done()
	LABEL:
	for {
		select{
		case <- ctx.Done():
			break LABEL
		default:
			fmt.Println("worker...")
			time.Sleep(time.Second)
		}


	}
	// 如何接收外部命令退出

}

func main() {
	// make和new的区别
	// 共同点：初始化内存
	// new()返回指针，多用于初始化基本数据类型
	// make用于初始化slice, map, channel
	// var exitChan = make(chan bool , 1)
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go worker(ctx)

	// 如何优雅的实现结束子goroutine

	time.Sleep(5 * time.Second)
	// eixt = true
	// exitChan <- true 
	cancel()

	wg.Wait()


	fmt.Println("========over===========")
}
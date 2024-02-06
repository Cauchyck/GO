package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)


func main(){
	
	fileName := `./logfile.log`

	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
	}
	// 打开文件开始读取数据
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Printf("tail %s failed, err: %v \n", fileName, err)
		return
	}

	var (		
		msg *tail.Line
		ok bool
	)

	for {
		msg, ok = <- tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename: %s \n",tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg: ", msg.Text)
	}



}
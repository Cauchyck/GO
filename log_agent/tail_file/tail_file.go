package tailfile

import (
	"context"
	"hello_go/log_agent/kafka"
	"time"

	"github.com/IBM/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

var (
	TailObj *tail.Tail
)

type tailTask struct {
	path   string
	topic  string
	tObj   *tail.Tail
	ctx    context.Context
	cancel context.CancelFunc
}

// tail 初始化
func newTailTask(path, topic string) *tailTask {
	ctx, cancel := context.WithCancel(context.Background())
	tt := &tailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
	}
	return tt
}

func (t *tailTask) run() {
	logrus.Infof("collect for path: %s is running...", t.path)
	for {
		select {
		case <-t.ctx.Done(): // 调用ctx.cancel则返回
			logrus.Infof("path: %s is stopped", t.path)
			return
		case line, ok := <-t.tObj.Lines:
			if !ok {
				logrus.Warningf("tail file close reopen, filename: %s \n", t.path)
				time.Sleep(time.Second)
				continue
			}

			// fmt.Println("msg: ", msg.Text)
			// 利用通道将同步的代码改为异步
			// 将读出的一行日志包装成kafka的msg类型，丢到通道中
			msg := &sarama.ProducerMessage{}
			msg.Topic = t.topic
			msg.Value = sarama.StringEncoder(line.Text)
			kafka.ToMsgChan(msg)
		}

	}
}

func (t *tailTask) Init() (err error) {

	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t.tObj, err = tail.TailFile(t.path, config)
	return
}

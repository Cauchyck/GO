package tailfile

import (
	"hello_go/log_agent/common"
	"github.com/sirupsen/logrus"
)

// 管理tail

type tailTaskMgr struct {
	tailTaskMap      map[string]*tailTask
	collectEntryList []common.ClloectEntry
	confChan         chan []common.ClloectEntry
}

var(
	ttMgr *tailTaskMgr
)

func Init(allConf []common.ClloectEntry) (err error) {
	// allConf包含若干日志收集项
	// 针对每个日志收集项创建一个对应的tailObj
	ttMgr = &tailTaskMgr{
		tailTaskMap: make(map[string]*tailTask, 20),
		collectEntryList: allConf,
		confChan: make(chan []common.ClloectEntry),

	}

	for _, conf := range allConf {  
		tt := newTailTask(conf.Path, conf.Topic)
		err = tt.Init()
		if err != nil {
			logrus.Errorf("Crate tailObj for path %s failed, err: %v \n", conf.Path, err)
			continue
		} 
		ttMgr.tailTaskMap[tt.path] = tt
		logrus.Infof("Create a tail task for path: %s", conf.Path)
		go tt.run()
	}
	go ttMgr.watch()

	return
}

func (t *tailTaskMgr)watch(){
	for{
		newConf := <-t.confChan
		logrus.Infof("Get new conf from etcd, comf: %v", newConf)
	
		for _, conf := range newConf  {  
			if t.isExist(conf){
				continue
			}
	
			tt := newTailTask(conf.Path, conf.Topic)
			err := tt.Init()
			if err != nil {
				logrus.Errorf("Crate tailObj for path %s failed, err: %v \n", conf.Path, err)
				continue
			} 
			t.tailTaskMap[tt.path] = tt
			logrus.Infof("Create a tail task for path: %s", conf.Path)
			go tt.run()
		}
		// 找出原来有，现在没有的tailtask停掉
		for key, task := range t.tailTaskMap{
			var found bool
			for _, conf := range newConf{
				if key == conf.Path{
					found = true
					break
				}
			}
			if !found{
				logrus.Infof("The task collect path %s need to stop!", key)
				delete(t.tailTaskMap,key)
				task.cancel()
			}
		}
	}
	
}

func (t *tailTaskMgr)isExist(conf common.ClloectEntry) bool{
	_, ok := t.tailTaskMap[conf.Path]
	return ok
}

func SendNewConf(newConf []common.ClloectEntry){
	ttMgr.confChan <- newConf
}
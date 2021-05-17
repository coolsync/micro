package taillog

import (
	"fmt"
	"logagent/etcd"
	"time"
)

var (
	tskMgr *tailLogMgr
)

// multiple tailtask manager
type tailLogMgr struct {
	logEntrys   []*etcd.LogEntry
	tailTaskMap map[string]*TailTask  // tailtask map
	newConfChan chan []*etcd.LogEntry // 不断从 etch watch 读取data
}

func Init(logEntrys []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntrys:   logEntrys,                      // 把当前的日志收集项配置信息保存起来
		tailTaskMap: make(map[string]*TailTask, 16), // 将 multiple tail task 配置形成 list
		newConfChan: make(chan []*etcd.LogEntry),    // 无缓冲 channel
	}

	for _, logEntry := range logEntrys {
		// 初始化的时候起了多少个tailtask 都要记下来，为了后续判断方便
		tailObj := NewTailTask(logEntry.Topic, logEntry.Path)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tailTaskMap[mk] = tailObj
	}

	go tskMgr.run() // 不断 从 etcd watch 读取 data, 修改 tail tsk map
}

func (tm *tailLogMgr) run() {
	for {
		select {
		case newConfs := <-tskMgr.newConfChan:
			for _, conf := range newConfs {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := tskMgr.tailTaskMap[mk]
				if ok {
					// 如果在 tail tsk map 存在， 不操作
					continue
				} else {
					// 如果在 tail tsk map 不存在， 添加操作
					tailObj := NewTailTask(conf.Topic, conf.Path)
					tskMgr.tailTaskMap[mk] = tailObj

				}
				// 比较 newConfs 的 task 与 logEntrys 已存在 task, 没有则删除
				for _, c1 := range tm.logEntrys  {
					isDelete := true
					for _, c2 := range newConfs{
						if c1.Path == c2.Path && c1.Topic == c2.Topic {
							isDelete = false
							continue
						}
					}
					if isDelete {	
								
					}
				}
			}
			// 配置变化
			// 配置删除
			fmt.Printf("New conf from etcd watch: %v\n", newConfs)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 对外开放一个 api， 将 etch watch 读取data 通过 开放 channel 写入到本地chan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}

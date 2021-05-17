package taillog

import (
	"logagent/etcd"
)

var (
	tailMgr *taillogMgr
)

type taillogMgr struct {
	logEntrys []*etcd.LogEntry
}

func Init(logEntrys []*etcd.LogEntry) {
	tailMgr = &taillogMgr{
		logEntrys: logEntrys, // 把当前的日志收集项配置信息保存起来
	}

	for _, logEntry := range logEntrys {
		NewTailTask(logEntry.Topic, logEntry.Path)
	}
}

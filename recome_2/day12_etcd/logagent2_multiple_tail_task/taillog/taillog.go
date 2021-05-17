package taillog

import (
	"fmt"
	"logagent/kafka"

	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

// TailTask: a log file collect task
type TailTask struct {
	topic    string
	path     string
	instance *tail.Tail
}

// New tail task
func NewTailTask(topic string, path string) (tailObj *tail.Tail) {
	tsk := &TailTask{
		topic: topic,
		path:  path,
	}

	tsk.init()

	return
}

// init single tail task
func (t *TailTask) init() (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}

	go t.run() // after collected log, direct send to kafka

	return
}

// line by line send to kafka
func (t *TailTask) run() {
	// for {
	// 	select {
	// 	case line := <-t.instance.Lines:
	// 		// kafka.SendToKafka(t.topic, line.Text) // slow, func call func
	// 		kafka.SendToChan(t.topic, line.Text)
	// 	}
	// }

	for line := range t.instance.Lines {
		kafka.SendToChan(t.topic, line.Text)
	}
}

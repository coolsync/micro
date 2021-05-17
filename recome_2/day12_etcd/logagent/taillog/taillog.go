package taillog

import (
	"context"
	"fmt"
	"logagent/kafka"

	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

// TailTask: a log file collect task
type TailTask struct {
	topic      string
	path       string
	instance   *tail.Tail
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// New tail task
func NewTailTask(topic string, path string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{ // 千万千万别写成 :=
		topic:      topic,
		path:       path,
		ctx:        ctx,
		cancelFunc: cancel,
	}

	tailObj.init() //  根据路径去打开对应的日志

	return
}

// init single tail task
func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		// return
	}

	// 当 goruntine run func 退出时， goroutine 结束
	go t.run() // after collected log, direct send to kafka
	// return
}

// line by line send to kafka
func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task %s_%s end!\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			// kafka.SendToKafka(t.topic, line.Text) // slow, func call func
			fmt.Printf("From %s get data %s success\n", t.path, line.Text)
			kafka.SendToChan(t.topic, line.Text)
		}
	}

	// for line := range t.instance.Lines {
	// 	kafka.SendToChan(t.topic, line.Text)
	// }
}

package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

// 日志收集配置项
type LogEntry struct {
	Path  string `json:"path"`  // log store path
	Topic string `json:"topic"` // 日志要发往 哪个 kafka topic
}

// init etcd cli
func Init(addr string, timeout time.Duration) (err error) {
	if addr == "" {
		fmt.Println("dail ip addr is null")
		return
	}

	// create etcd client
	cli, err = clientv3.New(clientv3.Config{
		Endpoints: []string{addr},
		// Endpoints:   addrs,
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

// Form Etcd 根据key获取配置项
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		// fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%v\n", err)
			return
		}
	}
	return
}

// watch conf file change
func WatchConf(key string, newConfChan chan<- []*LogEntry) {
	// watch 派一个哨兵 watch "topic" 这个key的变化
	wch := cli.Watch(context.Background(), key)

	// 从 channel 尝试取值 (监视的信息)
	for wresp := range wch { // <-chan WatchResponse
		for _, evt := range wresp.Events {
			// fmt.Printf("Type: %v, Key: %v, Value: %v\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
			fmt.Printf("Type: %s, Key: %s, Value: %s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)

			var newConf []*LogEntry
			// 判断 是否 是删除操作， delete operate 没有 value
			if evt.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("Unmarshal evt kv value failed, err: %v\n", err)
					continue
				}
			}

			// 将 获取的 data send to tailMgr newConfChan
			fmt.Printf("Get new conf: %v, send to tailMgr\n", newConf)

			newConfChan <- newConf
		}
	}

}

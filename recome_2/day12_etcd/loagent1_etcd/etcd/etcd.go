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

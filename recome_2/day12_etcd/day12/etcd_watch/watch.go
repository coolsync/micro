package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	// create etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// watch 派一个哨兵 watch "topic" 这个key的变化
	wch := cli.Watch(context.Background(), "topic")

	// 从 channel 尝试取值 (监视的信息)
	for wresp := range wch { // <-chan WatchResponse
		for _, evt := range wresp.Events {
			fmt.Printf("Type: %s, Key: %s, Value: %s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}
}

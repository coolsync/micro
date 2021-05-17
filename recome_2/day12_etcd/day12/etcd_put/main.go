package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
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

	// get ip
	ipStr, err := GetOutBoundIP()
	if err != nil {
		log.Fatal(err)
	}

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	key := fmt.Sprintf("/logagent/%s/collect_config", ipStr)
	// conf 2 kafka topic
	value := `[{"path":"/home/dart/Documents/log_tmp/log1/nginx.log","topic":"web_log"},{"path":"/home/dart/Documents/log_tmp/log2/redis.log","topic":"redis_log"}]`

	// conf 3 kafka topic and self path
	// value := `[{"path":"/home/dart/Documents/log_tmp/log1/nginx.log","topic":"web_log"},{"path":"/home/dart/Documents/log_tmp/log2/redis.log","topic":"redis_log"},{"path":"/home/dart/Documents/log_tmp/log2/mysql.log","topic":"mysql_log"}]`

	fmt.Println(value)
	_, err = cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Printf("net dial err: %v\n", err)
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]

	return
}

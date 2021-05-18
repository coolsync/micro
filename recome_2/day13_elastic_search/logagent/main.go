package main

import (
	"fmt"
	"logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"logagent/utils"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

// logAgent入口程序

var (
	cfg = new(conf.AppConf)
)

func main() {
	// Get Local out bound ip
	ipStr, err := utils.GetOutBoundIP()
	if err != nil {
		fmt.Printf("get ip failed: %v\n", err)
		return
	}

	// 1. 加载配置文件
	err = ini.MapTo(cfg, "conf/conf.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}

	cfg.KafkaConf.Address = ipStr + ":9092" // local addr give to conf addr
	fmt.Println(cfg.KafkaConf.Address)

	// 2. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	// 3. 初始化etcd连接
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")

	// 每个 logagent duo to etcd key 都有独自的配置， 可以使用 pull ip 方式， 也可以通过 业务线name 方式
	etcdKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	// 3.1 From etcd get log collection conf info
	logEntrys, err := etcd.GetConf(etcdKey)
	if err != nil {
		fmt.Printf("etcd.GetConf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get etcd conf info success: %v\n", logEntrys)

	for index, value := range logEntrys {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	// 4. 打开日志文件 准备收集日志
	taillog.Init(logEntrys)

	newConfChan := taillog.NewConfChan() // chan 在 taillog 初始化， 所以在 etcd.WatchConf 上面

	// 派一个哨兵去监视日志收集项的变化 （有变化及时通知log agent, 实现热加载配置）
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdKey, newConfChan) // 通知 上面 tail mgr chan
	wg.Wait()

	// select {}
}

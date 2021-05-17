package main

import (
	"fmt"
	"logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"time"

	"gopkg.in/ini.v1"
)

// logAgent入口程序

var (
	cfg = new(conf.AppConf)
)

func main() {
	// 1. 加载配置文件
	err := ini.MapTo(cfg, "conf/conf.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}

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

	// 3.1 From etcd get log collection conf info
	logEntrys, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("etcd.GetConf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get etcd conf info success: %v\n", logEntrys)

	// 3.2 派一个哨兵去监视日志收集项的变化 （有变化及时通知log agent, 实现热加载配置）
	for index, value := range logEntrys {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	// 4. 打开日志文件 准备收集日志
	taillog.Init(logEntrys)

}

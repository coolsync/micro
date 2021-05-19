package main

import (
	"fmt"
	"logtrans/conf"
	"logtrans/es"
	"logtrans/kafka"

	"gopkg.in/ini.v1"
)

func main() {
	// 1. Load conf file
	var cfg conf.AppConf
	err := ini.MapTo(&cfg, "conf/conf.ini")
	if err != nil {
		fmt.Printf("init conf failed, err: %v\n", err)
		return
	}
	fmt.Printf("conf file load ok: %v\n", cfg)

	// 2. init es
	err = es.Init(cfg.EsConf.Address)
	if err != nil {
		fmt.Printf("init es client failed, err: %v\n", err)
		return
	}
	fmt.Println("init es client success")

	// 3. init kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer client failed, err: %v\n", err)
		return
	}
	fmt.Println("init kafka consumer client success")

	// 4. get kafka log data
	// 5. send to es
	select {}
}

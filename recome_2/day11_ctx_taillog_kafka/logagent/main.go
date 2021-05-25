package main

import (
	"fmt"
	"log"
	"logagent/conf"
	"logagent/kafka"
	"logagent/taillog"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	cfg = new(conf.AppConf) // init yml app api
)

func run() {
	for {
		select {
		case line := <-taillog.LogDataChan(): // once read line, send to msg
			// fmt.Printf("tail file close reopen, filename:%s\n", tailObj.Filename)
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
			fmt.Println(line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// parse conf file
	ymlData, err := os.ReadFile("conf/conf.yml")
	if err != nil {
		log.Fatalf("read yml file failed, err:%v\n", err)
	}

	err = yaml.Unmarshal(ymlData, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 1. init kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("init kafka success!")
	
	// 2. init taillog
	err = taillog.Init(cfg.TailConf.FileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("init taillog success!")

	// 3. read log file send to kafka
	run()

	// select {}
}

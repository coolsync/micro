package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client      *elastic.Client
	LogDataChan chan *LogData
)

// init es client
func Init(addr string, chanSize, numsGo int) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))

	if err != nil {
		// fmt.Printf("")
		return
	}
	LogDataChan = make(chan *LogData, chanSize)
	// start nums goroutine
	for i := 0; i <= numsGo; i++ {
		go sendToES()
	}
	return
}

// recv kafka send to chan data
func SendToESChan(msg *LogData) {
	LogDataChan <- msg
}

// get kafka log send to es db
func sendToES() {
	for {
		select {
		case msg := <-LogDataChan:
			// insert record, PUT req, chain operation
			put1, err := client.Index().Index(msg.Topic).Type("xxx").BodyJson(msg).Do(context.Background())
			if err != nil {
				// Handle error
				fmt.Printf("send to es failed, err:%v\n", err)
				continue
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}

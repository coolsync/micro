package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic"
)

var (
	client *elastic.Client
)

// init es client
func Init(addr string) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))

	if err != nil {
		// fmt.Printf("")
		return
	}
	return
}

// get kafka log send to es db
func SendToES(indexStr string, data interface{}) (err error) {
	// insert record, PUT req, chain operation
	put1, err := client.Index().
		Index(indexStr).
		Type("xxx").
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Printf("send to es failed, err:%v\n", err)
		return
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return
}

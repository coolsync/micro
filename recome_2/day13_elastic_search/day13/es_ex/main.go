package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Worked bool   `json:"worked"`
}

func main() {
	// init connection, get cli obj
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.0.108:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")

	// insert record, PUT req
	p1 := Person{Name: "rion", Age: 22, Worked: false}
	// chain operation,
	put1, err := client.Index().
		Index("user").
		Type("person").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

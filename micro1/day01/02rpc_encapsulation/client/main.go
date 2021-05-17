package main

import (
	"day01/02rpc_encapsulation/design"
	"log"
)

func main() {
	// 1. use rpc protocol connect server
	// cli, err := design.InitClient("8081")
	cli, err := design.NewClient("8081")
	if err != nil {
		log.Fatalf("rpc dial failed, err: %v\n", err)
	}
	// create a var, recive server side send value, Outgoing parameters
	var reply string

	// 2. call remote methodj, param 2: Incoming parameters
	err = cli.HelloWorld("mark", &reply)
	if err != nil {
		log.Fatalf("rpc call remote method failed, err: %v\n", err)
	}
	// print server send info
	log.Println(reply)
}

// func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error

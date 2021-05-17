package main

import (
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	// 1. use rpc protocol connect server
	// cli, err := rpc.Dial("tcp", "localhost:8081")
	cli, err := jsonrpc.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("rpc dial failed, err: %v\n", err)
	}
	defer cli.Close()

	// create a var, recive server side send value, Outgoing parameters
	var reply string

	// 2. call remote methodj, param 2: Incoming parameters
	err = cli.Call("hello.HelloWorld", "bob", &reply)
	if err != nil {
		log.Fatalf("rpc dial failed, err: %v\n", err)
	}

	// print server send info
	log.Println(reply)
}

// func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error

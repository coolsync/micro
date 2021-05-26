package main

import (
	"day01/01rpc/design"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main01() {
	// rpc connect server
	// conn, err := rpc.Dial("tcp", "localhost:8081")
	conn, err := jsonrpc.Dial("tcp", "localhost:8081")

	if err != nil {
		log.Fatal("rpc.Dial: ", err)
	}
	defer conn.Close()

	// rpc call 远程方法
	var reply string // 传出值

	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		log.Fatal("conn.Call: ", err)
	}

	fmt.Println(reply)
}

func main() {
	// rpc connect server
	cli := design.NewClient("localhost:8081")

	// rpc call remote method
	var reply string // outGoing param, recv server return value

	// call method
	err := cli.HelloWorld("haha", &reply)
	if err != nil {
		log.Fatal("myclient.HelloWorld err: ", err)
	}

	fmt.Println("result: ", reply)
}

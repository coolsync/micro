package main

import (
	"day01/01rpc/design"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

// create object
type World struct {
}

// bind method to obj
func (*World) HelloWorld(name string, resp *string) error {
	*resp = name + " hello!"
	// *resp = "nihao !"
	return nil // {"id":0,"result":"李白 hello","error":null}
	// return errors.New("unknown err!") // {"id":0,"result":null,"error":"unknown err!"}, 结果返回给client 为nil
}

func main() {
	// 1. 注册rpc, 绑定 对象方法
	// err := rpc.RegisterName("hello", new(World))
	// if err != nil {
	// 	log.Fatal("rpc.RegisterName: ", err)
	// }

	design.RegisterService(new(World))

	// 2. set listener
	liser, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("net.Listen: ", err)
	}
	defer liser.Close()

	// 3. connect client
	fmt.Println("start listen ... ")
	conn, err := liser.Accept()
	if err != nil {
		log.Fatal("Accept(): ", err)
	}
	fmt.Println("conn ok ... ")

	defer conn.Close()

	// 4. 绑定rpc 到 tcp
	// rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
}

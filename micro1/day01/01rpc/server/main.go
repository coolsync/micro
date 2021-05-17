package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// define struct
type world struct {
}

// create receiver method
func (w *world) HelloWorld(name string, resp *string) error {
	*resp = "hello, " + name
	return nil
}

func main() {
	// 1. Register serve, specify serve name, register obj method
	err := rpc.RegisterName("hello", new(world))
	if err != nil {
		log.Fatalf("rpc register failed, err: %v", err)
	}

	// 2. Set listener
	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("net listen failed, err: %v", err)
	}

	// 3. Listen client connect
	log.Println("lis ok!")
	conn, err := lis.Accept()
	if err != nil {
		log.Fatalf("lis accept failed, err: %v", err)
	}
	defer conn.Close()
	log.Println("conn cli ok!")

	// 4. rpc bind tcp socket
	// rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)

}

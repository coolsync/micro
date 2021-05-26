package design

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MyInterface interface {
	HelloWorld(string, *string) error
}

func RegisterService(m MyInterface) {
	err := rpc.RegisterName("hello", m)
	if err != nil {
		log.Fatal("rpc.RegisterName: ", err)
	}
}

// ----------- client
type Myclient struct {
	c *rpc.Client
}

// init client
func NewClient(addr string) *Myclient {
	conn, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return &Myclient{c: conn}
}

//  invokes the named method
func (mc *Myclient) HelloWorld(a string, b *string) error {
	return mc.c.Call("hello.HelloWorld", a, b)
}

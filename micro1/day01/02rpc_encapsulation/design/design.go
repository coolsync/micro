package design

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Server-side encapsulation
// 1 define interface
type MyInterface interface {
	HelloWorld(string, *string) error
}

// 2 encapsulate register service method
func RegisterService(i MyInterface) error {
	err := rpc.RegisterName("hello", i)
	if err != nil {
		return err
	}
	return nil
}

// Client-side encapsulation
type MyClient struct {
	cli *rpc.Client
}

// cli.Call("hello.HelloWorld", "bob", &reply)

func (m *MyClient) HelloWorld(a string, b *string) error {
	return m.cli.Call("hello.HelloWorld", a, &b)
}

// Init client
func InitClient(addr string) (*MyClient, error) {
	cli, err := jsonrpc.Dial("tcp", "localhost:"+addr)
	if err != nil {
		return nil, err
	}
	// defer cli.Close()
	m := &MyClient{cli: cli}
	return m, nil
	// return &MyClient{cli: cli}, nil
}

func (m *MyClient) InitClient(addr string) (*MyClient, error) {
	cli, err := jsonrpc.Dial("tcp", "localhost:"+addr)
	if err != nil {
		return nil, err
	}

	return &MyClient{cli: cli}, nil
}

func NewClient(s string) (*MyClient, error) {
	m, err := InitClient(s)
	if err != nil {
		return nil, err
	}
	return m, nil
}

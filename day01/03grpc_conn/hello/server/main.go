package main

import (
	"context"
	pb "day01/03grpc_conn/proto/hello" // 引入编译生成的包
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:9001"
)

// 定义helloService并实现约定的接口
type helloService struct {
	pb.UnimplementedHelloServer
}

// HelloService Hello服务
// var HelloService = helloService{}

// type HelloService interface {
// 	SayHello(context.Context, *pb.HelloRequest) (*pb.HelloResponse, error)
// }

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	log.Println("server ok!")
	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	// pb.RegisterHelloServer(s, HelloService)
	pb.RegisterHelloServer(s, helloService{})

	grpclog.Infoln("Listen on " + Address)
	s.Serve(listen)
}

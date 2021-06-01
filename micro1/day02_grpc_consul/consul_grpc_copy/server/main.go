package main

import (
	"context"
	"errors"
	"implconsul/pb"
	"log"
	"net"

	consul_api "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserInfoServer
}

const (
	port = ":8800"
)

func (s *server) GetUserInfo(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	name := in.GetName()
	if name == "mark" {
		return &pb.UserResponse{
			Id:      1,
			Name:    name,
			Age:     32,
			Address: []string{"addr1", "addr2"},
		}, nil
	}
	return nil, errors.New("user no exist") // error strings should not end with punctuation or a newline (ST1005)
}

func main() {
	// +++++++++++ consul service register ++++++++++++
	// 1. init consul conf
	consul_conf := consul_api.DefaultConfig()
	// consul_conf.Address = "http://127.0.0.1:8800"

	// 2. create client consul_api
	consul_client, err := consul_api.NewClient(consul_conf)
	if err != nil {
		log.Fatal(err)
	}

	// 3. tell consul, Conf info of the service to be registered
	// svr,err := connect.NewService("bj38", consul_client)
	reg := consul_api.AgentServiceRegistration{
		ID:      "bj38",
		Tags:    []string{"grpc", "consul"},
		Name:    "Grpc and Consul",
		Address: "localhost",
		Port:    8800,
		Check: &consul_api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "localhost:8800",
			Timeout:  "1s",
			Interval: "5s",
		},
	}

	// 4. register grpc service to consul
	consul_client.Agent().ServiceRegister(&reg)

	// +++++++++++ grpc remote call ++++++++++++
	// 1. create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	// 2. instance grpc server
	s := grpc.NewServer()

	// 3. register service method
	// s.RegisterService(&server{}) // err
	pb.RegisterUserInfoServer(s, &server{})

	log.Println("server start ...")
	// 4. run serve
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

	// log.Println("server start ...")
}

package main

import (
	"context"
	"d02/01grpc/pb"
	"fmt"
	"log"
	"net"

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
	return nil, fmt.Errorf("user no exist!")
}

func main() {
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

	// 4. run serve
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

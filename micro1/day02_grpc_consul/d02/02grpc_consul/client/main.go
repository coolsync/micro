package main

import (
	"context"
	"d02/02grpc_consul/pb"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const (
	// address     = "localhost:8800"
	defaultName = "mark"
)

func main() {
	// +++++++++++ consul service find ++++++++++++
	// init consul conf
	consul_conf := api.DefaultConfig()

	// create client
	consul_client, err := api.NewClient(consul_conf)
	if err != nil {
		log.Fatal(err)
	}

	// from consul get health service
	// func (h *Health) Service(service, tag string, passingOnly bool, q *QueryOptions) ([]*ServiceEntry, *QueryMeta, error)
	services, _, err := consul_client.Health().Service("Grpc and Consul", "grpc", true, nil)
	if err != nil {
		log.Fatal(err)
	}

	// simple 负载均衡
	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	// +++++++++++ grpc remote call ++++++++++++
	// 1. dial grpc srv
	// cliConn, err := grpc.Dial("localhost:8800" grpc.WithInsecure(), grpc.WithBlock())
	cliConn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer cliConn.Close()

	// 2. from pb instance grpc client
	cli := pb.NewUserInfoClient(cliConn)

	// 3. assembly req data
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 4. use common method get resp data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	respData, err := cli.GetUserInfo(ctx, &pb.UserRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}
	// 7. print
	fmt.Printf("resp: %v\n", respData)
}

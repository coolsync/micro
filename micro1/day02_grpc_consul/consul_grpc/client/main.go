package main

import (
	"context"
	"fmt"
	"implconsul/consulexe"
	"implconsul/pb"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	// address     = "localhost:8800"
	defaultName = "mark"
)

func main() {
	// +++++++++++ consul service find ++++++++++++
	addrs, err := consulexe.ConsulFindService()
	if err != nil {
		log.Fatal(err)
	}

	addr := addrs[0] // get first node
	fmt.Println(addr)

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

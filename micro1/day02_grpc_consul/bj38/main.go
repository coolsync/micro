package main

import (
	"bj38/handler"
	pb "bj38/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("bj38"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterBj38Handler(srv.Server(), new(handler.Bj38))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

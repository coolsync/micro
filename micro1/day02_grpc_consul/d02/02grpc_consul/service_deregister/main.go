package main

import (
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	// init consul conf
	consul_conf := api.DefaultConfig()
	
	// create consul client
	consul_client, err := api.NewClient(consul_conf)
	if err != nil {
		log.Fatal(err)
	}

	// deregister service
	_ = consul_client.Agent().ServiceDeregister("bj38")
}

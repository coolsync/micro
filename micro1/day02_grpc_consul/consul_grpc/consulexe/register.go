package consulexe

import (
	"log"

	consul_api "github.com/hashicorp/consul/api"
)

func ConsulRegister() {
	// +++++++++++ consul service register ++++++++++++
	// 1. init consul conf
	consul_conf := consul_api.DefaultConfig()
	// fmt.Println(consul_conf.Address)	// 127.0.0.1:8500
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
}

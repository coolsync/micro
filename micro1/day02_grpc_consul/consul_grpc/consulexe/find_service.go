package consulexe

import (
	"errors"
	"strconv"

	consul_api "github.com/hashicorp/consul/api"
)

func ConsulFindService() ([]string, error) {
	// +++++++++++ consul service find ++++++++++++
	// init consul conf
	consul_conf := consul_api.DefaultConfig()

	// create client
	consul_client, err := consul_api.NewClient(consul_conf)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	// from consul get health service
	// func (h *Health) Service(service, tag string, passingOnly bool, q *QueryOptions) ([]*ServiceEntry, *QueryMeta, error)
	services, _, err := consul_client.Health().Service("Grpc and Consul", "grpc", true, nil)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	if len(services) == 0 {
		return nil, errors.New("no services")
	}

	// simple 负载均衡
	// addrs := make([]string, 1) // [""]
	var addrs []string
	// addr := ""
	for _, service := range services {
		addr := service.Service.Address + ":" + strconv.Itoa(service.Service.Port)
		if addr != "" {
			addrs = append(addrs, addr)
		}
	}
	// addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)
	return addrs, nil
}

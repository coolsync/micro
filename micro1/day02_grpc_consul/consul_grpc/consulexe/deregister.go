package consulexe

import (
	"github.com/hashicorp/consul/api"
)

func ConsulDeregister(service_name string) error {
	// init consul conf
	consul_conf := api.DefaultConfig()

	// create consul client
	consul_client, err := api.NewClient(consul_conf)
	if err != nil {
		// log.Fatal(err)
		return err
	}

	// deregister service
	err = consul_client.Agent().ServiceDeregister(service_name)
	if err != nil {
		return err
	}
	return nil
}

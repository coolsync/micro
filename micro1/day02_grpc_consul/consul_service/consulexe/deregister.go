package consulexe

import (
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

// 取消consul注册的服务
func ConsulDeRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "172.16.242.129:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	client.Agent().ServiceDeregister("111")
}

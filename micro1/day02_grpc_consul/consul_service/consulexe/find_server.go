package consulexe

import (
	"fmt"
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

// 从consul中发现服务
func ConsulFindServer() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "172.16.242.129:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// 获取所有service
	services, _ := client.Agent().Services()
	for _, value := range services {
		fmt.Println(value.Address)
		fmt.Println(value.Port)
	}

	fmt.Println("=================================")
	// 获取指定service
	service, _, err := client.Agent().Service("111", nil)
	if err == nil {
		fmt.Println(service.Address)
		fmt.Println(service.Port)
	}

}

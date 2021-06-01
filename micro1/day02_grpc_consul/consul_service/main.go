package main

import (
	"implconsul/consulexe"

	"github.com/gin-gonic/gin"

	"fmt"
	"log"
	"net/http"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	r := gin.Default()

	// consul健康检查回调函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// 注册服务到consul
	consulexe.ConsulRegister()

	// 从consul中发现服务
	consulexe.ConsulFindServer()

	// 取消consul注册的服务
	consulexe.ConsulDeRegister()

	http.ListenAndServe(":8081", r)
}

func ConsulKVTest() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = "172.16.242.129:8500"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// KV, put值
	values := "test"
	key := "go-consul-test/172.16.242.129:8100"
	client.KV().Put(&consulapi.KVPair{Key: key, Flags: 0, Value: []byte(values)}, nil)

	// KV get值
	data, _, _ := client.KV().Get(key, nil)
	fmt.Println(string(data.Value))

	// KV list
	datas, _, _ := client.KV().List("go", nil)
	for _, value := range datas {
		fmt.Println(value)
	}
	keys, _, _ := client.KV().Keys("go", "", nil)
	fmt.Println(keys)
}

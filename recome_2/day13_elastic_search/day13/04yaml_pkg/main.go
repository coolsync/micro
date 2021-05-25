package main

import (
	"day13/04yaml_pkg/conf"
	"fmt"
	"log"
)

func main() {
	yamlConf, err := conf.NewYamlConf("conf/conf.yml")
	if err != nil {
		log.Fatalf("init conf file failed, err: %v\n", err)
	}

	fmt.Println("yamlConf: ", yamlConf)

	// get conf data
	fmt.Printf("mysql, host:%s, port: %d, dbname: %s\n", yamlConf.MysqlConf.Host, yamlConf.MysqlConf.Port, yamlConf.MysqlConf.DBName)

	fmt.Printf("cache, enable: %t,list: %v\n", yamlConf.CacheConf.Enable, yamlConf.CacheConf.List)
}

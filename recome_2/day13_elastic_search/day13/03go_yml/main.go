package main

import (
	"day13/03go_yml/conf"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// new conf struct
	yamlConf := new(conf.YamlConf)

	// read yaml file, get []data
	yamlBytes, err := os.ReadFile("03go_yml/conf/conf.yml")
	if err != nil {
		log.Fatalf("read yaml file err: %v\n", err)
	}
	log.Println("read yamlBytes ok: ", string(yamlBytes))

	// unmarshal file data to struct
	err = yaml.Unmarshal(yamlBytes, yamlConf)
	if err != nil {
		log.Fatalf("Unmarshal file to struct err: %v\n", err)
	}

	// bind map
	confMap := make(map[string]interface{})

	err = yaml.Unmarshal(yamlBytes, confMap)
	if err != nil {
		log.Fatalf("Unmarshal file to map err: %v\n", err)
	}

	fmt.Println("confMap: ", confMap)

	for k, v := range confMap {
		_, is_str := v.(string)
		_, is_uint32 := v.(uint32)
		_, is_bool := v.(bool)

		if is_str {
			fmt.Printf("key: %s, value(string):%v\n", k, v)
		} else if is_uint32 {
			fmt.Printf("key: %s, value(uint32):%v\n", k, v)
		} else if is_bool {
			fmt.Printf("key: %s, value(bool):%v\n", k, v)
		} else {
			fmt.Printf("key: %s, value:%v\n", k, v)
		}
	}

	// 使用map之前已经知道file结构是什么， 可以提取自己想要的
	// panic: interface conversion: interface {} is map[string]interface {}, not map[interface {}]interface {}
	mySqlValInterface := confMap["mysql"]
	fmt.Printf("%T\n", mySqlValInterface)
	mysqlValue, ok := mySqlValInterface.(interface{})
	fmt.Println(mysqlValue)
	if ok {
		// extract
		// if mysqlMap, ok := mysqlValue.(map[interface{}]interface{}); ok {	// "gopkg.in/yaml.v2"
		if mysqlMap, ok := mysqlValue.(map[string]interface{}); ok { // "gopkg.in/yaml.v3"
			fmt.Printf("mysql parse: host: %s, port: %d\n", mysqlMap["host"], mysqlMap["port"])
		} else {
			fmt.Println("no mysql info")
		}
	} else {
		fmt.Println("mysql value is not interface")
	}

	bytes, err := yaml.Marshal(&confMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("confMap dump: \n\n%v\n", string(bytes))
}

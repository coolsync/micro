package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type YamlConf struct {
	AppConf   `yaml:"app"`
	CacheConf `yaml:"cache"`
	MysqlConf `yaml:"mysql"`
}

type AppConf struct {
	Name string `yaml:"name"`
	// ConfFile string `yaml:"conffile"`
}

type CacheConf struct {
	Enable bool     `yaml:"enable"`
	List   []string `yaml:"list,flow"`
}
type MysqlConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint32 `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

func NewYamlConf(confFile string) (*YamlConf, error) {

	// new conf struct
	yamlConf := new(YamlConf)

	// read yaml file, get []data
	yamlBytes, err := os.ReadFile(confFile)
	if err != nil {
		fmt.Printf("read yaml file err: %v\n", err)
		return nil, err
	}
	// log.Println("read yamlBytes ok: ", string(yamlBytes))

	// unmarshal file data to struct
	err = yaml.Unmarshal(yamlBytes, yamlConf)
	if err != nil {
		fmt.Printf("Unmarshal file to struct err: %v\n", err)
		return nil, err
	}
	// fmt.Println("yamlConf: ", yamlConf)
	return yamlConf, nil
}

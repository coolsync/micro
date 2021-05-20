package conf

type YamlConf struct {
	AppConf   `yaml:"app"`
	CacheConf `yaml:"cache"`
	MysqlConf `yaml:"mysql"`
}

type AppConf struct {
	Name string `yaml:"name"`
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

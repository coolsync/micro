package conf

type AppConf struct {
	KafkaConf `yaml:"kafka"`
	TailConf  `yaml:"taillog"`
}

type KafkaConf struct {
	Address string `yaml:"address"`
	Topic   string `yaml:"topic"`
}

type TailConf struct {
	FileName string `yaml:"filename"`
}

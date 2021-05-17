package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

type KafkaConf struct {
	Address string `ini:"address"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"collect_log_key"`
	Timeout int    `ini:"timeout"`
}

// type TaillogConf struct {
// 	FileName string `ini:"filename"`
// }

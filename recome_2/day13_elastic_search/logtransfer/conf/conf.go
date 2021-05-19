package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EsConf    `ini:"es"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EsConf struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chan_size"`
	Nums     int    `int:"nums"`
}

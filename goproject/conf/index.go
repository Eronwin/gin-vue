package conf

import (
	"io/ioutil"
	"log"
)

type Config struct {
	DB string
}

func LoadConf() *Config {
	source, errRead := ioutil.ReadFile("./conf/conf.yml")
	if errRead != nil {
		log.Println("[读取yml失败]", errRead.Error())
	}
	config := Config{}
	err := yaml.Unmarshal(source, &config)
	if err != nil {
		log.Println("[初始化配置失败]", err.Error())
	}
	return &config
}

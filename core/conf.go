package core

import (
	"fmt"
	"go-blog/config"
	"go-blog/global"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// 读yaml文件配置信息
func InitConf() {
	const ConfigFile = "/home/mr4/gitSrc/go-blog/setting.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	if err = yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}

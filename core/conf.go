package core

import (
	"fmt"
	"go-blog/config"
	"go-blog/global"
	"io/fs"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const ConfigFile = "/home/mr4/gitSrc/go-blog/setting.yaml"

// 读yaml文件配置信息
func InitConf() {

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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}

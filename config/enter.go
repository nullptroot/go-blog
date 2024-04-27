package config

import "fmt"

// 这里的` `是用来标识序列化和反序列化的文件格式和对应的字段
type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

package autoconfig

import (
	"bbs-forgo/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//解析yml文件
type BaseInfo struct {
	Port string     `yaml:"port"`
	Ip   string     `yaml:"ip"`
	Host string     `yaml:"host"`
	Base BaseEntity `yaml:"base"`
}

type BaseEntity struct {
	Redis    RedisData    `yaml:"redis"`
	Database DatabaseData `yaml:"database"`
}

type RedisData struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DataBase string `yaml:"dataBase"`
	Timeout  string `yaml:"timeout"`
}

type DatabaseData struct {
	DBType   string `yaml:"dbtype"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (c *BaseInfo) GetConf(filename string) *BaseInfo {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.GetLogger().Error(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.GetLogger().Error(err.Error())
	}
	return c
}

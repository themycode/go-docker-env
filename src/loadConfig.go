package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// 读取配置文件
func Str(st string) string {
	return strings.ToUpper(st)
}

type MainConfig struct {
	Port    string
	Address string
}

func LoadConfig(path string) *MainConfig {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed", string(buf), err)
	}

	fmt.Println("mainConfig:", mainConfig.Address, mainConfig.Port)

	return mainConfig
}

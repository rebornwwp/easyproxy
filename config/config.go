package config

import (
	"encoding/json"
	"github.com/rebornwwp/easyproxy/structure"
	"io/ioutil"
	"log"
)

type Config struct {
	Service      string              `json:"service"`
	Host         string              `json:"host"`
	Port         uint16              `json:"port"`
	WebPort      uint16              `json:"webport"`
	Strategy     string              `json:"strategy"`
	HeartBeat    int                 `json:"heartbeat"`
	MaxProcessor int                 `json:"maxprocesser"`
	Backends     []structure.Backend `json:"backends"`
}

// NewConfig 从一个json文件获取配置信息
func NewConfig(filename string) (*Config, error) {
	var config Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("load config file failed.", err)
	} else {
		err = json.Unmarshal(file, &config)
		if err != nil {
			log.Println("decode json file failed", err)
		}
	}
	return &config, err
}

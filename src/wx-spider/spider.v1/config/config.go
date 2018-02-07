package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	WxIdList []map[string]string   `json:"wxIdList"`
}

func New(path string) (config *Config, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	println("config: ", string(data))
	config = &Config{}
	err = json.Unmarshal(data, config)
	return
}

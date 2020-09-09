package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var defaultConfigPath = "./conf/config.json"

var Config Configuration

type Configuration struct {
	ApiBaseURL    string `json:"api_base_url"`
	BangDBSpec    string `json:"bang_db_spec"`
	NoticeDBSpec  string `json:"notice_db_spec"`
	RedisAddr     string `json:"redis_addr"`
	RedisPassword string `json:"redis_password"`
	RedisDB       int    `json:"redis_db"`
}

func InitConfigs() error {
	data, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		return fmt.Errorf("init config error, err=%v", err)
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		return nil
	}
	return nil
}

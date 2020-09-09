package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	defaultConfigPath = "conf/conf.json"
)

var (
	Config Configuration
)

type Configuration struct {
	AppID         string   `json:"AppID"`
	AppSecret     string   `json:"AppSecret"`
	AppUserId     string   `json:"AppUserId"`
	PermanentCode string   `json:"PermanentCode"`
	Token         string   `josn:"token"`
	URL           string   `json:"url"`
	ONESBaseURL   string   `json:"ones_base_url"`
	BAUsers       []string `json:"ba_users"`
	CorpID        string   `json:"corp_id"`
	CorpSecret    string   `json:"corp_secret"`
	AgentID       int64    `json:"agent_id"`
	WechatBaseURL string   `json:"weichat_base_url"`
}

func InitConfigs() error {
	data, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		return fmt.Errorf("init config error: %v", err)
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		return nil
	}
	return nil
}

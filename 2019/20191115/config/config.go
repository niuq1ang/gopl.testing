package config

import (
	"encoding/json"
	"io/ioutil"
)

const DefaultConfigPath = "./conf/config.json"

var Config Configuration

type Configuration struct {
	ProjectDBSpec string `json:"project_db_spec"`
	WikiDBSpec    string `json:"wiki_db_spec"`
}

func LoadConfigs(filepath string) (err error) {
	if filepath == "" {
		filepath = DefaultConfigPath
	}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	return json.Unmarshal(data, &Config)
}

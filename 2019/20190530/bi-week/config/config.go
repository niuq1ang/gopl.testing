package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const defaultConfigPath = "conf/conf.json"

var config map[string]interface{}

func LoadConfigs(configPath string) (err error) {
	if configPath == "" {
		configPath = defaultConfigPath
	}
	file, err := os.Open(configPath)
	if err != nil {
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return
}

func String(key string, defaultValue string) (value string) {
	value, _ = getString(key, defaultValue)
	return
}

func getString(key string, defaultValue string) (value string, found bool) {
	v, found := config[key]
	if !found {
		return defaultValue, false
	}
	if value, ok := v.(string); ok {
		return value, ok
	} else {
		return defaultValue, false
	}
}

func StringOrError(key string) (value string, err error) {
	value, found := getString(key, "")
	if !found {
		err = fmt.Errorf("%s is not configured", key)
	}
	return
}

func StringOrPanic(key string) string {
	v, err := StringOrError(key)
	if err != nil {
		panic(err)
	}
	return v
}

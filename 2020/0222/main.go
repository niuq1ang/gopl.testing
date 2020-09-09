package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Test struct {
	Config string
	Array  []string
	Simple string
}
var loginType []string

func main() {
	allFlags := pflag.NewFlagSet("all", pflag.ExitOnError)
	allFlags.StringArrayVar(&loginType, "login_type", []string{}, "")
	fmt.Println(loginType)
}

func initConfig(cmd *cobra.Command) error {
	viper.BindPFlags(cmd.Flags())
	viper.SetConfigName("config")
	// viper.SetConfigType()
	viper.AddConfigPath(".")
	if cfgFile != "" {
		content, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			return err
		}
		fmt.Println(string(content))
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired, %v", err)
		}
		log.Fatal("Config file was found but another error was produced, %v", err)
	}

	err := viper.Unmarshal(&options)
	if err != nil {
		log.Fatal("unmarshall config error, %v", err)
	}

	if err := retrieveConfigJson(); err != nil {
		log.Fatal("retrieve config.json err:%v", err)
	}
	return nil
}

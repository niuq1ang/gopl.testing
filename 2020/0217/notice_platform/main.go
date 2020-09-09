package main

import (
	_ "github.com/bangwork/ones-plugin/notice_platform/app"
	"github.com/bangwork/ones-plugin/notice_platform/app/router"
)

func main() {
	// config, err := node.StartPluginSrv("./conf/configuration.json", 2000, 5)
	// if err != nil {
	// 	panic(err)
	// }
	// router.Run(config.NodeAddress)
	router.Run("127.0.0.1:8088")
}

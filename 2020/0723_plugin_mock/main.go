package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type User struct {
	Name string
	Age  int
}

var pluginSecretMap map[string]string

func main() {
	run()
	initPluginSecretMap()
}

func initPluginSecretMap() {
	pluginSecretMap = make(map[string]string)
}

func setPluginSecretMap(key, value gjson.Result) bool {

	return true
}

func getRouterList(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", []byte(pluginJson))
	return
}

func reciveSecretKeys(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	gjson.Get(string(data), "").ForEach(setPluginSecretMap)

	fmt.Printf("secrets: %s\n", data)
	c.String(http.StatusOK, "%s", "ok")

}

func run() {
	router := gin.Default()

	pluginHost := router.Group("/plugin")
	pluginHost.GET("/router_list", getRouterList)
	pluginHost.POST("/set_secrets", reciveSecretKeys)

	router.Run("0.0.0.0:9005")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Run() error {
	r := SetupRouter()
	return r.Run(":8999")
}
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/wechat/callback", handlerGrafanaMsg)

	return router
}

func handlerGrafanaMsg(context *gin.Context) {
	bodyBytes, err := context.GetRawData()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(string(bodyBytes))
}

func main() {
	Run()

}

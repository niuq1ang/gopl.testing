package main

import (
	"io/ioutil"
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	Run("0.0.0.0:8999")
}

func Run(addr string) {
	app := iris.Default()
	app.Use(recover.New())

	app.Post("/handler/crm/message", Accept)

	app.Run(iris.Addr(addr))
}

func Accept(c iris.Context) {
	crmData, _ := ioutil.ReadAll(c.Request().Body)
	log.Println(c.GetReferrer())
	log.Println(c.RemoteAddr())
	log.Printf("RequestBody: %v\n", string(crmData))
	// if err := api.HandlerCRMRequest(crmData); err != nil {
	// 	fmt.Println(err)
	// }
	c.WriteString("ok！！！！")
}

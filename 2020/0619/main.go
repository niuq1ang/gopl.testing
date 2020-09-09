package main

import (
	"io/ioutil"
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	Run()
}

func Run() {
	app := iris.Default()
	app.Use(recover.New())

	app.Post("/fxxk/approval/flow", Accept)

	app.Run(iris.Addr("0.0.0.0:8998"))
}

func Accept(c iris.Context) {
	crmData, _ := ioutil.ReadAll(c.Request().Body)
	log.Println("RemoteAddr: ", c.RemoteAddr())
	log.Printf("RequestBody: %v\n", string(crmData))

	c.WriteString("ok")
}

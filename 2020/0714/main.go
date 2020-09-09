package main

import (
	"fmt"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	teamUUID := strings.Split("/team/:teamUUID/plugin/upload_opk", "/")[2]
	fmt.Println(teamUUID)
	// Run("0.0.0.0:8888")
}

func Run(addr string) {
	app := iris.Default()
	app.Use(recover.New())

	app.Post("/test", Accept)
	app.Get("/test", Accept)
	app.Run(iris.Addr(addr))
}

func Accept(c iris.Context) {
	fmt.Printf("path=%s", c.Path())
	fmt.Printf("path=%s", string(c.Path()))
}

package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	Run()
}

type Book struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	NamePinyin string `json:"name_pinyin"`
}

type Books struct {
	List []Book `json:"books"`
}

// 实现几个Web接口

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Plugin got <%s> <%s> <%s>", ctx, ctx.Method(), ctx.Path())
	ctx.Next()

}

func Run() {
	api := iris.Default()
	api.Use(myMiddleware)

	// new api
	api.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"msg": "pong"})
	})

	// replace existed api
	api.Get("/echo", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "welcome from plugin"})
	})

	// hook api (after)
	book := api.Party("/book")
	book.Get("/list", func(ctx iris.Context) {
		books := []Book{Book{
			UUID:       "001",
			Name:       "万里十五年",
			NamePinyin: "wanlishiwunian",
		},
			Book{
				UUID:       "002",
				Name:       "月亮与六便士",
				NamePinyin: "yueliangyuliubianshi",
			}}
		ctx.JSON(books)
	})

	book.Post("/add", func(ctx iris.Context) {
		var books Book
		if err := ctx.ReadJSON(&books); err != nil {
			fmt.Println(err)
			ctx.JSON(iris.Map{"result": "param error"})
			return
		}
		fmt.Printf("Got books: %+v\n", books)
		ctx.JSON(iris.Map{"status": 0})
	})

	// http://localhost:8088.
	api.Run(iris.Addr(":8088"))
}

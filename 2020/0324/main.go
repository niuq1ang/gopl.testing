package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bangwork/mygo/2020/0324/api"
	"github.com/bangwork/mygo/2020/0324/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	if err := config.InitConfigs(); err != nil {
		panic(err)
	}
	if err := api.InitWecahtToken(); err != nil {
		panic(err)
	}
	// api.QueryProjectApprovalByID("5e79c1f343c6930001b6f793")
	// fmt.Println(api.AddCustomerField("[测试添加属性值]"))
	// fmt.Println(api.CreateWiki("[么得感情]测试标题"))
	// api.InitUserMap()
	// fmt.Println(api.QueryDemandDetailByProjectID("5e79c1f343c6930001b6f793"))
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
	log.Printf("RequestBody: %v\n", string(crmData))
	if err := api.HandlerCRMRequest(crmData); err != nil {
		fmt.Println(err)
	}
	c.WriteString("ok")
}

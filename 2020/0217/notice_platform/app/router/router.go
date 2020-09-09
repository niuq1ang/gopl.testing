package router

import (
	. "github.com/bangwork/ones-plugin/notice_platform/app/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func Run(addr string) {
	app := iris.Default()
	app.Use(recover.New())

	team := app.Party("/team/{teamUUID}/webhook")
	team.Post("/update", UpdateWebhook)
	team.Get("/list/issue", ListWebHookIssue)
	team.Get("/list", ListWebHookList)
	team.Get("/list/uuid/{uuid}", ListWebHookDetail)
	team.Get("/delete/uuid/{uuid}", DeleteWebhook)
	team.Get("/test", TestWebhook)
	team.Get("/verify/name", VerifyWebhookName)

	app.Run(iris.Addr(addr))
}

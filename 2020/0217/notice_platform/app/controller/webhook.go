package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	commonerrors "github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/webhook"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/ngaut/log"
)

func getTeamUUID(c iris.Context) string { return c.Params().Get("teamUUID") }
func getUUID(c iris.Context) string     { return c.Params().Get("uuid") }

func UpdateWebhook(c iris.Context) {
	teamUUID := getTeamUUID(c)
	obj := new(webhook.Webhook)
	if err := BindJSONWithReason(c, obj); err != nil {
		return
	}
	err := webhook.UpsertWebHook(teamUUID, obj)
	RenderJSON(c, err, nil)

}
func DeleteWebhook(c iris.Context) {
	teamUUID := getTeamUUID(c)
	uuid := getUUID(c)
	err := webhook.DeleteWebHook(teamUUID, uuid)
	RenderJSON(c, err, nil)

}
func VerifyWebhookName(c iris.Context) {
	teamUUID := getTeamUUID(c)
	name := c.URLParam("name")
	err := webhook.VerifyWebhookName(teamUUID, name)
	RenderJSON(c, err, nil)
}

func ListWebHookDetail(c iris.Context) {
	teamUUID := getTeamUUID(c)
	uuid := getUUID(c)
	webhookType, err := c.URLParamInt64("webhook_type")
	if err != nil {
		RenderJSON(c, errors.InvalidParameterError("", "webhook_type", "not in domain defined"), nil)
		return
	}
	obj, err := webhook.ListWebhookDetail(teamUUID, webhookType, uuid)
	RenderJSON(c, err, obj)
}

func ListWebHookIssue(c iris.Context) {
	teamUUID := getTeamUUID(c)
	obj, err := webhook.ListIssues(teamUUID)
	RenderJSON(c, err, obj)
}

func ListWebHookList(c iris.Context) {
	teamUUID := getTeamUUID(c)
	obj, err := webhook.ListWebhooks(teamUUID)
	RenderJSON(c, err, obj)
}

func TestWebhook(c iris.Context) {
	url := c.URLParam("url")
	err := webhook.WebhookTest(url)
	RenderJSON(c, err, nil)
}

func BindJSONWithReason(c iris.Context, obj interface{}) (err error) {
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
			renderBindJSONError(c, err)
		}
	}()

	if err = c.ReadJSON(obj); err != nil {
		renderBindJSONError(c, err)
		return err
	}

	return nil
}

func renderBindJSONError(c iris.Context, err error) {
	code := commonerrors.Code(commonerrors.Malformed, commonerrors.JSON)
	err = commonerrors.Wrapf(err, code, "error at parsing json: "+err.Error())
	RenderJSON(c, err, nil)
}

func RenderJSON(c iris.Context, result error, obj interface{}) {
	errp := buildErrPayloadAndLog(result, true)
	c.ContentType(context.ContentJSONHeaderValue)
	c.StatusCode(errp.HttpStatus)

	var err error
	if obj == nil {
		res, _ := json.Marshal(errp)
		_, err = c.Write(res)
	} else {
		res, _ := json.Marshal(obj)
		_, err = c.Write(res)
	}

	if err != nil {
		code := commonerrors.Code(commonerrors.Malformed, commonerrors.JSON)
		err = commonerrors.Wrapf(err, code, "render json failed")
		log.Error(err)
	}
	c.Next()
}

func buildErrPayloadAndLog(err error, shouldLog bool) (errp *commonerrors.ErrPayload) {
	errp = commonerrors.NewErrPayload(err)
	if shouldLog {
		// 根据状态码打印日志
		if errp.HttpStatus < 400 {
			// 不需要打印日志
		} else if errp.HttpStatus >= 500 && errp.HttpStatus < 600 {
			// 服务端错误
			log.Error(err)
			// 对客户端隐藏详细信息
			errp.Code = commonerrors.ServerError
			errp.HttpStatus = http.StatusInternalServerError
			errp.Desc = ""
			errp.Values = nil
		} else {
			// 客户端错误 & 自定义错误，Warn
			log.Warn(err)
		}
	}
	return
}

package webhook

import (
	"fmt"
	"strings"
	"time"

	"github.com/bangwork/ones-plugin/notice_platform/app/utils"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"gopkg.in/gorp.v1"
)

const (
	FilterTypeTaskFilter = 1

	WebHookIsBatchEnable = 1
	WebHookIsBatchUnable = 2

	WebHookStatusValid      = 1
	WebHookStatusInvalid    = 2
	WebhookStatusUnavilable = 3
	WebHookStatusDelete     = 4
	webHookColumns          = "uuid, team_uuid, name, url, is_batch, batch_count, status, create_time, update_time"
)

type WebHook struct {
	UUID       string `db:"uuid"`
	TeamUUID   string `db:"team_uuid"`
	Name       string `db:"name"`
	URL        string `db:"url"`
	IsBatch    int64  `db:"is_batch"`
	BatchCount int64  `db:"batch_count"`
	Status     int64  `db:"status"`
	CreateTime int64  `db:"create_time"`
	UpdateTime int64  `db:"update_time"`
}

func ListAllWebHooks(src gorp.SqlExecutor) ([]*WebHook, error) {
	var webHookList []*WebHook
	sql := fmt.Sprintf("SELECT %s FROM webhook;", webHookColumns)
	_, err := src.Select(&webHookList, sql)
	if err != nil {
		return nil, errors.Sql(err)
	}
	return webHookList, nil
}

func VerifyWebhookName(src gorp.SqlExecutor, teamUUID, name string) (*WebHook, error) {
	w := new(WebHook)
	sql := fmt.Sprintf("SELECT %s FROM webhook WHERE team_uuid=? AND name=?;", webHookColumns)
	rows, err := src.Select(&w, sql, teamUUID, name)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		return rows[0].(*WebHook), nil
	}
	return nil, nil
}

func DeleteWebhookByUUID(src gorp.SqlExecutor, uuid string) error {
	sql := fmt.Sprintf("DELETE FROM webhook where uuid=?;")
	_, err := src.Exec(sql, uuid)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func ListWebHooksByTeamUUID(src gorp.SqlExecutor, teamUUID string) ([]*WebHook, error) {
	var webHookList []*WebHook
	sql := fmt.Sprintf("SELECT %s FROM webhook where team_uuid=?;", webHookColumns)
	_, err := src.Select(&webHookList, sql, teamUUID)
	if err != nil {
		return nil, errors.Sql(err)
	}
	return webHookList, nil
}

func GetWebhookByUUID(src gorp.SqlExecutor, uuid string) (*WebHook, error) {
	w := new(WebHook)
	sql := fmt.Sprintf("SELECT %s FROM webhook WHERE uuid=?;", webHookColumns)
	rows, err := src.Select(&w, sql, uuid)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		return rows[0].(*WebHook), nil
	}
	return nil, nil
}

func InsertIntoWebhook(src gorp.SqlExecutor, item *WebHook) error {
	if item == nil {
		return nil
	}
	sql := fmt.Sprintf("INSERT INTO webhook(%s) VALUES (?,?,?,?,?,?,?,?,?);", webHookColumns)
	_, err := src.Exec(sql, item.UUID, item.TeamUUID, item.Name, item.URL, item.IsBatch, item.BatchCount, item.Status, item.CreateTime, item.UpdateTime)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func UpdateWebhook(src gorp.SqlExecutor, item *WebHook) error {
	if item == nil {
		return nil
	}
	var build strings.Builder
	var args []interface{}
	build.WriteString("UPDATE webhook SET ")
	if !utils.IsEmptyString(item.Name) {
		build.WriteString("`name`=?, ")
		args = append(args, item.Name)
	}

	if !utils.IsEmptyString(item.URL) {
		build.WriteString("`url`=?, ")
		args = append(args, item.URL)
	}

	if item.IsBatch != 0 {
		build.WriteString("`is_batch`=?, ")
		args = append(args, item.IsBatch)
	}

	if item.BatchCount != 0 {
		build.WriteString("`batch_count`=?, ")
		args = append(args, item.BatchCount)
	}

	if item.Status != 0 {
		build.WriteString("`status`=?, ")
		args = append(args, item.Status)
	}

	build.WriteString(fmt.Sprintf("`update_time`=? WHERE uuid='%s';", item.UUID))
	args = append(args, time.Now().Unix())
	_, err := src.Exec(build.String(), args...)
	if err != nil {
		return err
	}
	return nil
}

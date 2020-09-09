package tunnel

import (
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"gopkg.in/gorp.v1"
)

const (
	TunnelStatusValid   = 1
	TunnelStatusInValid = 2

	LinkTypeWebHook  = 1
	LinkTypeWechat   = 2
	LinkTypeDingding = 3
	tunnelColumns    = "uuid, team_uuid, link_type, link_uuid, filter_type, filter_uuid"
)

type Tunnel struct {
	UUID       string `db:"uuid"`
	TeamUUID   string `db:"team_uuid"`
	LinkType   int    `db:"link_type"`
	LinkUUID   string `db:"link_uuid"`
	FilterType int    `db:"filter_type"`
	FilterUUID string `db:"filter_uuid"`
}

func ListAllTunnels(src gorp.SqlExecutor) ([]*Tunnel, error) {
	var tunnelList []*Tunnel
	sql := fmt.Sprintf("SELECT %s FROM tunnel", tunnelColumns)
	_, err := src.Select(&tunnelList, sql)
	if err != nil {
		return nil, errors.Sql(err)
	}
	return tunnelList, nil
}

func DeleteTunnelByUUID(src gorp.SqlExecutor, uuid string) error {
	sql := fmt.Sprintf("DELETE FROM tunnel where uuid=?;")
	_, err := src.Exec(sql, uuid)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func ListTunnelByWebhookUUID(src gorp.SqlExecutor, uuid string) (*Tunnel, error) {
	var tunnelList []*Tunnel
	sql := fmt.Sprintf("SELECT %s FROM tunnel WHERE link_uuid='%s';", tunnelColumns, uuid)
	_, err := src.Select(&tunnelList, sql)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(tunnelList) > 0 {
		return tunnelList[0], nil
	}
	return nil, nil
}

func InsertIntoTunnel(src gorp.SqlExecutor, item *Tunnel) error {
	if item == nil {
		return nil
	}
	fieldsNumber := 6
	sql := fmt.Sprintf("INSERT INTO tunnel(%s) VALUES %s;", tunnelColumns, utils.SqlInMultiInsertValues(fieldsNumber, 1))
	_, err := src.Exec(sql, item.UUID, item.TeamUUID, item.LinkType, item.LinkUUID, item.FilterType, item.FilterUUID)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func InsertIntoTunnels(src gorp.SqlExecutor, items []*Tunnel) error {
	length := len(items)
	if length == 0 {
		return nil
	}
	fieldsNumber := 6
	sql := fmt.Sprintf("INSERT INTO tunnel(%s) VALUES %s;", tunnelColumns, utils.SqlInMultiInsertValues(fieldsNumber, length))
	args := make([]interface{}, fieldsNumber*length)
	for i, item := range items {
		args[i*fieldsNumber+0] = item.UUID
		args[i*fieldsNumber+1] = item.TeamUUID
		args[i*fieldsNumber+2] = item.LinkType
		args[i*fieldsNumber+3] = item.LinkUUID
		args[i*fieldsNumber+4] = item.FilterType
		args[i*fieldsNumber+5] = item.FilterUUID
	}
	_, err := src.Exec(sql, args...)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

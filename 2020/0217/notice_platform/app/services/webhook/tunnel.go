package webhook

import (
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	filterModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/filter"
	tunnrlModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/tunnel"
	webhookModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/webhook"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/filter"
)

type AbstractTunnel struct {
	*tunnrlModel.Tunnel
	*webhookModel.WebHook
	*filterModel.TaskFilter
}

var TunnelMap map[string]*AbstractTunnel

func InitTunnelMap() error {
	tunnelList, err := tunnrlModel.ListAllTunnels(db.NoticeDBMap)
	if err != nil {
		return err
	}
	TunnelMap = make(map[string]*AbstractTunnel, len(tunnelList))
	for _, t := range tunnelList {
		w, ok := WebHookMap[t.LinkUUID]
		if !ok {
			continue
		}
		f, ok := filter.TaskFilterMap[t.FilterUUID]
		if !ok {
			continue
		}
		TunnelMap[t.UUID] = &AbstractTunnel{
			Tunnel:     t,
			WebHook:    w,
			TaskFilter: f,
		}
	}
	return nil
}

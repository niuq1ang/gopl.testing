package webhook

import (
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	noticeModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/notice"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/record"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/tunnel"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/webhook"
)

func DoSend(n *noticeModel.Notice, tunnelType int, tunnelUUID string) {
	switch tunnelType {
	case tunnel.LinkTypeWebHook:
		w, ok := WebHookMap[tunnelUUID]
		if ok && w.Status == webhook.WebHookStatusValid {
			executeWebHook(w, n)
		} else {
			record.UpdateRecordBySelNo(db.NoticeDBMap, record.RecordStatusFailure, n.SeqNo)
		}
	case tunnel.LinkTypeWechat:
	case tunnel.LinkTypeDingding:
	default:
	}
}

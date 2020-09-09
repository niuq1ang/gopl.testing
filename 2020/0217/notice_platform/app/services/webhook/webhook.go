package webhook

import (
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	noticeModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/notice"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/webhook"
)

var (
	client     map[string]*Client
	WebHookMap map[string]*webhook.WebHook
)

func InitWebHookMap() error {
	WebHookMap = make(map[string]*webhook.WebHook, 0)
	webHooks, err := webhook.ListAllWebHooks(db.NoticeDBMap)
	if err != nil {
		return err
	}
	for _, webHook := range webHooks {
		WebHookMap[webHook.UUID] = webHook
	}
	return nil
}

func executeWebHook(w *webhook.WebHook, n *noticeModel.Notice) {
	if len(client) == 0 {
		client = make(map[string]*Client, 0)
	}
	c, ok := client[w.UUID]
	if !ok {
		c = NewClient(w)
		client[c.UUID] = c
	}

	go c.AddNotice(n)
}

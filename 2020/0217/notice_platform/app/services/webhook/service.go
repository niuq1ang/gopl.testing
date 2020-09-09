package webhook

import (
	"fmt"
	"time"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	filterModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/filter"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/tunnel"
	tunnelModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/tunnel"
	webhookModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/webhook"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/filter"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"gopkg.in/gorp.v1"
)

const (
	WebhookModel      = "webhook"
	WebhookNameMaxLen = 32
	WebhookURLMaxLen  = 1024
)

type WebhookList struct {
	WebhookList []*Webhook `json:"webhook_list"`
}

type WebhookIssues struct {
	IssueTypeList  []*filter.FieldChooseEvent `json:"issue_type_list"`
	IssueFieldList []*filter.FieldChooseEvent `json:"issue_field_list"`
}

type Webhook struct {
	UUID           string                     `json:"uuid"`
	WebhookType    int64                      `json:"webhook_type"`
	Name           string                     `json:"name"`
	URL            string                     `json:"url"`
	Status         int64                      `json:"status"`
	IsBatch        int64                      `json:"is_batch"`
	BatchCount     int64                      `json:"batch_count"`
	IssueTypeList  []*filter.FieldChooseEvent `json:"issue_type_list,omitempty"`
	IssueFieldList []*filter.FieldChooseEvent `json:"issue_field_list,omitempty"`
}

func VerifyWebhookName(teamUUID, name string) error {
	if utils.IsEmptyString(name) || len(name) > WebhookNameMaxLen {
		return errors.InvalidParameterError(WebhookModel, "name", "is empty or exceed maximum length limit")
	}
	result, err := webhookModel.VerifyWebhookName(db.NoticeDBMap, teamUUID, name)
	if err != nil {
		return err
	}
	if result != nil {
		return errors.VerificationFailureError("duplicate name")
	}
	return nil
}

func WebhookTest(url string) error {
	if utils.IsEmptyString(url) || len(url) > WebhookURLMaxLen {
		return errors.InvalidParameterError(WebhookModel, "url", "is empty or exceed maximum length limit")
	}
	err := PostHeartbeat(url)
	if err != nil {
		return errors.VerificationFailureError(err.Error())
	}
	return nil
}

func DeleteWebHook(teamUUID, uuid string) error {
	t, err := tunnelModel.ListTunnelByWebhookUUID(db.NoticeDBMap, uuid)
	if err != nil {
		return err
	}
	if t == nil {
		return errors.VerificationFailureError("uuid not exists")
	}
	err = db.DBMTransact(db.NoticeDBMap, func(tx *gorp.Transaction) error {
		if err := tunnelModel.DeleteTunnelByUUID(tx, t.UUID); err != nil {
			return err
		}
		if err := webhookModel.DeleteWebhookByUUID(tx, t.LinkUUID); err != nil {
			return err
		}
		if err := filterModel.DeleteTaskFilterByUUID(tx, t.FilterUUID); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	delete(filter.TaskFilterMap, t.FilterUUID)
	delete(WebHookMap, t.LinkUUID)
	delete(TunnelMap, t.UUID)
	return nil
}

func UpsertWebHook(teamUUID string, w *Webhook) error {
	if err := w.validate(); err != nil {
		return err
	}

	switch w.WebhookType {
	case tunnel.LinkTypeWebHook:
		if w.UUID == "" { // create
			if err := VerifyWebhookName(teamUUID, w.Name); err != nil {
				return err
			}
			if err := w.Save(teamUUID); err != nil {
				return errors.Trace(err)
			}
		} else { // update
			if !utils.IsEmptyString(w.Name) {
				result, err := webhookModel.VerifyWebhookName(db.NoticeDBMap, teamUUID, w.Name)
				if err != nil {
					return err
				}
				if result != nil && result.UUID != w.UUID {
					return errors.VerificationFailureError("duplicate name")
				}
			}
			if err := w.Update(teamUUID); err != nil {
				return errors.Trace(err)
			}
		}
	default:
		return errors.InvalidParameterError(WebhookModel, "webhook_type", "not in domain defined")
	}
	return nil
}

func ListWebhooks(teamUUID string) (*WebhookList, error) {
	webhookList, err := webhookModel.ListWebHooksByTeamUUID(db.NoticeDBMap, teamUUID)
	if err != nil {
		return nil, err
	}
	list := new(WebhookList)
	for _, w := range webhookList {
		wh := &Webhook{
			UUID:   w.UUID,
			Name:   w.Name,
			URL:    w.URL,
			Status: w.Status,
		}
		list.WebhookList = append(list.WebhookList, wh)
	}
	return list, nil
}

func ListIssues(teamUUID string) (*WebhookIssues, error) {
	issueField := filter.ListAllFieldEvent()
	issueType, err := filter.ListAllIssueTypeEvent(teamUUID)
	if err != nil {
		return nil, err
	}
	wh := &WebhookIssues{
		IssueTypeList:  issueType,
		IssueFieldList: issueField,
	}
	return wh, nil
}

func ListWebhookDetail(teamUUID string, webhookType int64, uuid string) (*Webhook, error) {
	switch webhookType {
	case tunnel.LinkTypeWebHook:
		w, err := webhookModel.GetWebhookByUUID(db.NoticeDBMap, uuid)
		if err != nil {
			return nil, err
		}
		if w == nil {
			return &Webhook{}, nil
		}
		t, err := tunnelModel.ListTunnelByWebhookUUID(db.NoticeDBMap, w.UUID)
		if err != nil {
			return nil, err
		}
		f, err := filterModel.ListTaskFilterByUUID(db.NoticeDBMap, t.FilterUUID)
		if err != nil {
			return nil, err
		}
		issueField := filter.Int64ToSelectedField(f.FieldFilter)
		issueType, err := filter.Int64ToSelectedIssueType(teamUUID, f.IssueTypeFilter)
		if err != nil {
			return nil, err
		}
		wh := &Webhook{
			UUID:           w.UUID,
			WebhookType:    tunnel.LinkTypeWebHook,
			Name:           w.Name,
			URL:            w.URL,
			Status:         w.Status,
			IsBatch:        w.IsBatch,
			BatchCount:     w.BatchCount,
			IssueTypeList:  issueType,
			IssueFieldList: issueField,
		}
		return wh, nil
	default:
		return nil, errors.InvalidParameterError(WebhookModel, "webhook_type", "not in domain defined")
	}
}

func (w *Webhook) Update(teamUUID string) error {
	err := db.DBMTransact(db.NoticeDBMap, func(tx *gorp.Transaction) error {
		t, err := tunnelModel.ListTunnelByWebhookUUID(tx, w.UUID)
		if err != nil {
			return err
		}
		if t == nil {
			return fmt.Errorf("internal error")
		}
		wh := &webhookModel.WebHook{
			UUID:       w.UUID,
			URL:        w.URL,
			Name:       w.Name,
			IsBatch:    w.IsBatch,
			BatchCount: w.BatchCount,
			Status:     w.Status,
		}
		if err := webhookModel.UpdateWebhook(tx, wh); err != nil {
			return err
		}
		if len(w.IssueFieldList)+len(w.IssueTypeList) > 0 {
			issueTypeInt64, err := filter.SelectedIssueTypeToInt64(teamUUID, w.IssueTypeList)
			if err != nil {
				return err
			}
			issueFieldInt64 := filter.SelectedFieldToInt64(w.IssueFieldList)
			if err := filterModel.UpdateFilter(tx, t.FilterUUID, issueFieldInt64, issueTypeInt64); err != nil {
				return err
			}
		}
		err = UpdateCache(wh.UUID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateCacheByWebhook(w *webhookModel.WebHook) error {
	err := db.DBMTransact(db.NoticeDBMap, func(tx *gorp.Transaction) error {
		if err := webhookModel.UpdateWebhook(tx, w); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return UpdateCache(w.UUID)
}

// UpdateCache
func UpdateCache(uuid string) error {
	t, err := tunnelModel.ListTunnelByWebhookUUID(db.NoticeDBMap, uuid)
	if err != nil {
		return err
	}
	if t == nil {
		return fmt.Errorf("internal error")
	}
	w, err := webhookModel.GetWebhookByUUID(db.NoticeDBMap, uuid)
	if err != nil {
		return err
	}
	if w == nil {
		return fmt.Errorf("internal error")
	}
	f, err := filterModel.ListTaskFilterByUUID(db.NoticeDBMap, t.FilterUUID)
	if err != nil {
		return err
	}
	if f == nil {
		return fmt.Errorf("internal error")
	}
	// update cache
	filter.TaskFilterMap[f.UUID] = f
	WebHookMap[w.UUID] = w
	TunnelMap[t.UUID] = &AbstractTunnel{
		Tunnel:     t,
		WebHook:    w,
		TaskFilter: f,
	}
	return nil
}
func (w *Webhook) Save(teamUUID string) error {
	webhookUUID, filterUUID := utils.Random62String(8), utils.Random62String(8)
	wh := &webhookModel.WebHook{
		UUID:       webhookUUID,
		TeamUUID:   teamUUID,
		Name:       w.Name,
		URL:        w.URL,
		IsBatch:    w.IsBatch,
		BatchCount: w.BatchCount,
		Status:     w.Status,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	issueInt64, err := filter.SelectedIssueTypeToInt64(teamUUID, w.IssueTypeList)
	if err != nil {
		return err
	}
	f := &filterModel.TaskFilter{
		UUID:            filterUUID,
		TeamUUID:        teamUUID,
		FieldFilter:     filter.SelectedFieldToInt64(w.IssueFieldList),
		IssueTypeFilter: issueInt64,
	}

	t := &tunnelModel.Tunnel{
		UUID:       utils.Random62String(8),
		TeamUUID:   teamUUID,
		LinkType:   tunnelModel.LinkTypeWebHook,
		LinkUUID:   webhookUUID,
		FilterType: webhookModel.FilterTypeTaskFilter,
		FilterUUID: filterUUID,
	}
	err = db.DBMTransact(db.NoticeDBMap, func(tx *gorp.Transaction) error {
		if err := tunnelModel.InsertIntoTunnel(tx, t); err != nil {
			return err
		}
		if err := webhookModel.InsertIntoWebhook(tx, wh); err != nil {
			return err
		}
		if err := filterModel.InsertIntoFilter(tx, f); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	//  add to cache
	TunnelMap[t.UUID] = &AbstractTunnel{
		Tunnel:     t,
		WebHook:    wh,
		TaskFilter: f,
	}
	return nil
}

func (w *Webhook) validate() error {
	if utils.IsEmptyString(w.UUID) { // create
		if utils.IsEmptyString(w.Name) || len(w.Name) > WebhookNameMaxLen {
			return errors.InvalidParameterError(WebhookModel, "name", "is empty or exceed maximum length limit")
		}
		if utils.IsEmptyString(w.URL) || len(w.URL) > WebhookURLMaxLen || !utils.IsURL(w.URL) {
			return errors.InvalidParameterError(WebhookModel, "url", "is empty or exceed maximum length limit")
		}
		if w.Status < 1 || w.Status > 3 {
			return errors.InvalidParameterError(WebhookModel, "status", "not in domain defined")
		}
		if w.IsBatch == 0 {
			w.IsBatch = 2
		}
		if w.IsBatch < 1 || w.IsBatch > 2 {
			return errors.InvalidParameterError(WebhookModel, "is_batch", "not in domain defined")
		}
		if w.IsBatch == 1 {
			if w.BatchCount < 0 || w.BatchCount > 200 {
				return errors.InvalidParameterError(WebhookModel, "batch_count", "not in domain defined")
			}
		}
		if len(w.IssueFieldList) < 1 {
			return errors.InvalidParameterError(WebhookModel, "issue_field_list", "is empty")
		}
		if len(w.IssueTypeList) < 1 {
			return errors.InvalidParameterError(WebhookModel, "issue_type_list", "is empty")
		}
	} else { // update
		if len(w.UUID) != 8 {
			return errors.InvalidParameterError(WebhookModel, "uuid", "not match with fixed length")
		}
		if len(w.Name) > WebhookNameMaxLen {
			return errors.InvalidParameterError(WebhookModel, "name", "is empty or exceed maximum length limit")
		}
		if !utils.IsEmptyString(w.URL) && (len(w.URL) > WebhookURLMaxLen || !utils.IsURL(w.URL)) {
			return errors.InvalidParameterError(WebhookModel, "url", "is empty or exceed maximum length limit")
		}
		if w.Status < 0 || w.Status > 4 {
			return errors.InvalidParameterError(WebhookModel, "status", "not in domain defined")
		}

		if w.IsBatch < 0 || w.IsBatch > 2 {
			return errors.InvalidParameterError(WebhookModel, "is_batch", "not in domain defined")
		}
		if w.IsBatch == 1 {
			if w.BatchCount < 0 || w.BatchCount > 200 {
				return errors.InvalidParameterError(WebhookModel, "batch_count", "not in domain defined")
			}
		}
	}
	return nil
}

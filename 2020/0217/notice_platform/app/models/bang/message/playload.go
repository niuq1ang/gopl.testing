package message

import (
	"encoding/json"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/batch"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/resource"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
)

type MessagePayload struct {
	UUID                  string                    `json:"uuid"`
	TeamUUID              string                    `json:"team_uuid"`
	RefType               string                    `json:"ref_type"`
	RefID                 string                    `json:"ref_id"`
	Type                  string                    `json:"type"`
	FromUUID              string                    `json:"from"`
	FromName              string                    `json:"from_name,omitempty"`
	ToUUID                string                    `json:"to"`
	ToName                string                    `json:"to_name,omitempty"`
	SendTime              int64                     `json:"send_time"`
	Text                  string                    `json:"text,omitempty"`
	Resource              *resource.ResourcePayload `json:"resource,omitempty"`
	SubjectType           string                    `json:"subject_type,omitempty"`
	SubjectID             string                    `json:"subject_id,omitempty"`
	SubjectName           string                    `json:"subject_name,omitempty"`
	Action                string                    `json:"action,omitempty"`
	ObjectType            string                    `json:"object_type,omitempty"`
	ObjectID              string                    `json:"object_id,omitempty"`
	ObjectName            string                    `json:"object_name,omitempty"`
	ObjectAttr            string                    `json:"object_attr,omitempty"`
	OldValue              string                    `json:"old_value,omitempty"`
	NewValue              string                    `json:"new_value,omitempty"`
	Ext                   interface{}               `json:"ext,omitempty"`
	IsCanShowRichtextDiff bool                      `json:"is_can_show_richtext_diff"`
}

type UUIDNameExt struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
type TaskFieldExt struct {
	FieldUUID         string        `json:"field_uuid"`
	FieldName         string        `json:"field_name,omitempty"`
	FieldType         int           `json:"field_type"`
	OldValue          string        `json:"old_value"`
	NewValue          string        `json:"new_value"`
	OldOption         *UUIDNameExt  `json:"old_option,omitempty"`
	NewOption         *UUIDNameExt  `json:"new_option,omitempty"`
	NewMultiOption    []UUIDNameExt `json:"new_multi_option,omitempty"`
	OldMultiOption    []UUIDNameExt `json:"old_multi_option,omitempty"`
	BatchAction       batch.Action  `json:"batch_action,omitempty"`
	ParentMessageUUID string        `json:"parent_message_uuid,omitempty"` // 在F1231迭代中，需要当前message是基于某一条message的，便于前端生成「父子消息」的文案
	TriggerTaskUUID   string        `json:"trigger_task_uuid,omitempty"`   //F1286, task statusCategory 可以由其它 task 的后置动作更改
	TriggerTaskTitle  string        `json:"trigger_task_title,omitempty"`  //F1286, 原 task 的信息, ex: 「#234 任务summary」
}

func NewMessagePayload(m *Message, r *resource.FileResource) (payload *MessagePayload, err error) {
	p := MessagePayload{
		UUID:     m.UUID,
		TeamUUID: m.TeamUUID,
		RefID:    m.ReferenceID,
		FromUUID: m.FromUUID,
		ToUUID:   m.ToUUID,
		Text:     m.Text,
		SendTime: m.SendTime,
	}

	label, err := utils.LabelForEntityType(m.ReferenceType)
	if err != nil {
		return
	}
	p.RefType = label

	label, err = LabelForMessageType(m.Type)
	if err != nil {
		return
	}
	p.Type = label

	if r != nil {
		var err error
		p.Resource, err = resource.NewResourcePayload(r)
		if err != nil {
			return nil, err
		}
	}

	if m.SubjectType > 0 {
		label, err = utils.LabelForEntityType(m.SubjectType)
		if err != nil {
			return
		}
		p.SubjectType = label
		p.SubjectID = m.SubjectID
	}

	if m.Action > 0 {
		label, err = utils.LabelForEntityAction(m.Action)
		if err != nil {
			return
		}
		p.Action = label
	}

	if m.ObjectType > 0 {
		label, err = utils.LabelForEntityType(m.ObjectType)
		if err != nil {
			return
		}
		p.ObjectType = label
		p.ObjectID = m.ObjectID
		p.ObjectName = m.ObjectName
		if m.ObjectAttr > 0 {
			label, err = utils.LabelForEntityAttr(m.ObjectAttr)
			if err != nil {
				return
			}
			p.ObjectAttr = label
		}
	}

	p.OldValue = m.OldValue
	p.NewValue = m.NewValue

	if utils.NonEmptyString(m.Ext) {
		var ext interface{}
		if err = json.Unmarshal([]byte(m.Ext), &ext); err != nil {
			return
		}
		p.Ext = ext
	}

	payload = &p
	return
}

func LabelForMessageType(messageType int) (label string, err error) {
	switch messageType {
	case MessageTypeDiscussion:
		label = MessageTypeDiscussionLabel
	case MessageTypeSystem:
		label = MessageTypeSystemLabel
	default:
		err = errors.InvalidEnumError(messageType, "Message", errors.Type)
	}
	return
}

type ProjectIssueTypeExt struct {
	ProjectUUID   string `json:"project_uuid"`
	IssueTypeUUID string `json:"issue_type_uuid"`
}

type JobProgressExt struct {
	Success    int `json:"success"`
	Fail       int `json:"fail"`
	Incomplete int `json:"incomplete"`
}

type AddTaskFeildExt struct {
	Fields []TaskFieldExt `json:"field_values"`
}

type RelatedTestCaseExt struct {
	LibUUID      string `json:"library_uuid"`
	ModuleUUID   string `json:"module_uuid"`
	CaseUUID     string `json:"case_uuid"`
	TestCaseName string `json:"case_name"`
}

type RelatedPlanCaseExt struct {
	LibUUID      string `json:"library_uuid"`
	ModuleUUID   string `json:"module_uuid"`
	CaseUUID     string `json:"case_uuid"`
	TestCaseName string `json:"case_name"`
	PlanUUID     string `json:"plan_uuid"`
}

type TaskRelatedTestCasePlanExt struct {
	PlanUUID string `json:"plan_uuid"`
	PlanName string `json:"plan_name"`
}

type CopyTaskExt struct {
	UUID   string         `json:"uuid"`
	Name   string         `json:"name"`
	Fields []TaskFieldExt `json:"field_values"`
}
type SubtaskActionExt struct {
	Summary string `json:"summary,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type AttachmentsExt struct {
	Attachments []*UUIDNameExt `json:"attachments"`
}

type TaskTargetActionExt struct {
	Name  string `json:"name,omitempty"`
	Score *int64 `json:"score"`
	Unit  string `json:"unit"`
}

type PluginExt struct {
	UUID   string `json:"uuid"`
	IsOpen bool   `json:"is_open"`
}

type ManhourActionExt struct {
	Total int64 `json:"total"`
}

type TaskNoticeMessageExt struct {
	FieldUUID    string        `json:"field_uuid"`
	FieldName    string        `json:"field_name"`
	DistanceTime *DistanceTime `json:"distance_time"`
}

type DistanceTime struct {
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type TaskRelatedWikiPageExt struct {
	PageUUID  string `json:"page_uuid"`
	PageTitle string `json:"page_title"`
}

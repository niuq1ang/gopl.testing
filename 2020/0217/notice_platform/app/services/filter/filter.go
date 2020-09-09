package filter

import (
	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/message"
	taskModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/task"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/webhook"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"github.com/ngaut/log"
)

func DoFilter(m *message.Message, task *taskModel.Task, filterType int, filterUUID string) bool {
	log.Info("filer message: ", *m)
	switch filterType {
	case webhook.FilterTypeTaskFilter:
		if f, ok := TaskFilterMap[filterUUID]; ok {
			switch m.Type {
			case message.MessageTypeSystem:
				// TODO task_notice 判断方式 待观察
				// msg.ObjectAttr == utils.EntityAttrTaskNotice
				// m.ReferenceType == utils.EntityTypeProject && m.SubjectType == utils.EntityTypeTask
				if m.ObjectAttr == utils.EntityAttrTaskNotice {
					return true
				}
				issueTypeBit, err := getIssueTypeInt64(task.TeamUUID, task.StandardIssueTypeUUID)
				if err != nil {
					log.Error(errors.Trace(err))
					return false
				}
				// check issue_type
				if f.IssueTypeFilter&issueTypeBit == issueTypeBit {
					if issueFieldBit, ok := IssueFieldIntergeMap[GetMessageNoticeType(m)]; ok {
						// check issue_field
						if f.FieldFilter&issueFieldBit == issueFieldBit {
							return true
						}
					}
				}

			case message.MessageTypeDiscussion:
			}
		}
	}
	return false
}

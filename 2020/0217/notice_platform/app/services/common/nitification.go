package common

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/batch"
	fieldModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/field"
	issueTypeModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/issuetype"
	messageModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/message"
	resourceModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/resource"
	taskModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/task"
	userModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/user"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/common/i18n"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/common/mention"
	utilsModel "github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils/timestamp"
	"github.com/ngaut/log"
	"golang.org/x/text/language"
	"gopkg.in/gorp.v1"
)

var lang = []language.Tag{language.Chinese}

func BuildNoticeDescription(t *taskModel.Task, m *messageModel.Message, userMp map[string]*userModel.User) (actionTxt string, ok bool) {
	switch m.Type {
	case messageModel.MessageTypeSystem:
		actionTxt = buildNoticeDescriptionFromSystem(t, m, userMp)
	case messageModel.MessageTypeDiscussion:
		actionTxt = buildNoticeDescriptionFromDiscussion(t, m, userMp)
	default:
		//
	}
	ok = actionTxt != ""
	return
}

func buildNoticeDescriptionFromSystem(t *taskModel.Task, m *messageModel.Message, userMp map[string]*userModel.User) (actionTxt string) {
	switch m.ObjectAttr {
	case 0:
		actionTxt = buildNoticeDescriptionFromTask(t, m)
	case utilsModel.EntityAttrStatus:
		actionTxt = buildNoticeDescriptionFromTask(t, m)
	case utilsModel.EntityAttrField:
		actionTxt = buildNoticeDescriptionFromTaskField(t, m)
	case utilsModel.EntityAttrRelatedTask:
		actionTxt = buildNoticeDescriptionFromTaskRelatedTask(t, m)
	case utilsModel.EntityAttrSubtask:
		actionTxt = buildNoticeDescriptionFromTaskSubtask(t, m)
	case utilsModel.EntityAttrWatchers:
		actionTxt = buildNoticeDescriptionFromTaskWatchers(t, m, userMp)
	case utilsModel.EntityAttrAssessManhour:
		actionTxt = buildNoticeDescriptionFromTaskManhour(t, m)
	case utilsModel.EntityAttrAttachments:
		actionTxt = buildNoticeDescriptionFromTaskAttachments(t, m)
	case utilsModel.EntityAttrRelatedTestCase:
		actionTxt = buildNoticeDescriptionFromTaskRelatedTestCase(t, m)
	case utilsModel.EntityAttrRelatedPlanCase:
		actionTxt = buildNoticeDescriptionFromTaskRelatedPlanCase(t, m)
	case utilsModel.EntityAttrTaskRelatedTestCasePlan:
		actionTxt = buildNoticeDescriptionFromTaskRelatedPlan(t, m)
	case utilsModel.EntityAttrDeadline:
		actionTxt = buildNoticeDescriptionFromTaskDeadline(t, m)
	case utilsModel.EntityAttrIssueType:
		actionTxt = buildNoticeDescriptionFromIssueType(t, m)
	case utilsModel.EntityAttrParent:
		actionTxt = buildNoticeDescriptionFromUpdateParent(t, m)
	case utilsModel.EntityAttrManhours:
		actionTxt = buildNoticeDescriptionFromRecordManhour(t, m)
	case utilsModel.EntityAttrProduct:
		actionTxt = buildNoticeDescriptionFromTaskField(t, m)
	case utilsModel.EntityAttrTaskNotice:
		actionTxt = buildNoticeDescriptionFromTaskNotice(t, m)
	case utilsModel.EntityAttrTaskRelatedWikiPage:
		actionTxt = buildNoticeDescriptionFromTaskRelatedPage(t, m)
	default:
		log.Warn("buildNoticeDescriptionFromSystem not found type.")
		//
	}
	return
}

func buildNoticeDescriptionFromIssueType(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	var ext messageModel.TaskFieldExt
	ext.BatchAction = batch.BatchActionModifyIssueType

	if len(m.Ext) > 1 {
		err := json.Unmarshal([]byte(m.Ext), &ext)
		if err != nil {
			ext.BatchAction = batch.BatchActionModifyIssueType
		}
	}

	newIssueType, err := issueTypeModel.GetIssueType(db.BangDBMap, t.TeamUUID, m.NewValue)
	if err != nil {
		log.Warn("get new issue type was err: %v", err)
		return ""
	}
	oldIssueType, err := issueTypeModel.GetIssueType(db.BangDBMap, t.TeamUUID, m.OldValue)
	if err != nil {
		log.Warn("get old issue type was err: %v", err)
		return ""
	}

	switch ext.BatchAction {
	case batch.BatchActionModifyIssueType,
		batch.BatchActionSubToSubIssueType,
		batch.BatchActionStdToStdIssueType:
		actionTxt = i18n.Stringf(lang, "notice.updateIssueType", oldIssueType.Name, newIssueType.Name)
	case batch.BatchActionStdToSubIssueType:
		actionTxt = i18n.Stringf(lang, "notice.updateStdToSubIssueType", newIssueType.Name)
	case batch.BatchActionSubToStdIssueType:
		actionTxt = i18n.Stringf(lang, "notice.updateSubToStdIssueType", newIssueType.Name)
	}
	return
}

func buildNoticeDescriptionFromUpdateParent(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	var ext messageModel.TaskFieldExt
	ext.BatchAction = batch.BatchActionSubModifyParent

	if len(m.Ext) > 1 {
		err := json.Unmarshal([]byte(m.Ext), &ext)
		if err != nil {
			ext.BatchAction = batch.BatchActionSubModifyParent
		}
	}

	newTaskParent, err := taskModel.GetTask(db.BangDBMap, m.NewValue)
	if err != nil {
		log.Warn("get new issue type was err: %v", err)
		return ""
	}
	oldTaskParent, err := taskModel.GetTask(db.BangDBMap, m.OldValue)
	if err != nil {
		log.Warn("get old issue type was err: %v", err)
		return ""
	}

	actionTxt = i18n.Stringf(lang, "notice.updateSubIssueTypeParent", oldTaskParent.Summary, newTaskParent.Summary)
	return
}

////
func buildNoticeDescriptionFromDiscussion(t *taskModel.Task, m *messageModel.Message, userMp map[string]*userModel.User) (actionTxt string) {
	if m.ResourceUUID != "" {
		r, result := MustGetResource(db.BangDBMap, t.TeamUUID, m.ResourceUUID)
		if result != nil {
			return
		}
		actionTxt = i18n.Stringf(lang, "notice.fileUploaded", r.Name)
		return
	}
	text := m.Text
	allMentions := mention.Parse(text)
	for _, m := range allMentions {
		userUUID := m.Text
		user, found := userMp[userUUID]
		if found {
			old := fmt.Sprintf("@%s", userUUID)
			new := fmt.Sprintf("@%s", user.Name)
			text = strings.Replace(text, old, new, -1)
		}
	}
	actionTxt = i18n.Stringf(lang, "notice.sentAMessage", text)
	return
}

func buildNoticeDescriptionFromTask(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	taskName := t.Summary
	issueTypeName := ""
	it, _ := issueTypeModel.GetIssueType(db.BangDBMap, t.TeamUUID, t.StandardIssueTypeUUID)
	if it != nil {
		issueTypeName = it.Name
	}
	switch m.Action {
	case utilsModel.EntityActionAdd:
		actionTxt = i18n.Stringf(lang, "notice.taskCreated", issueTypeName, taskName)
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.taskDeleted", issueTypeName, taskName)
	case utilsModel.EntityActionCopy:
		actionTxt = i18n.String(lang, "notice.taskCopy")
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskField(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.TaskFieldExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	fieldUUID := ext.FieldUUID
	fieldName := ext.FieldName
	if fieldUUID == fieldModel.DescFieldUUID {
		actionTxt = i18n.String(lang, "notice.descChanged")
		return
	}

	addUnitFunc := func(value string) string {
		if fieldUUID == fieldModel.AssessManhourFieldUUID || fieldUUID == fieldModel.RemainingManhourFieldUUID {
			hourUnit := i18n.Stringf(lang, "unit.hour")
			return fmt.Sprintf("%s%s", value, hourUnit)
		}
		return value
	}

	switch m.Action {
	case utilsModel.EntityActionAdd:
		newValue := getFieldValueText(false, ext)
		actionTxt = i18n.Stringf(lang, "notice.setFieldTo", fieldName, addUnitFunc(newValue))
	case utilsModel.EntityActionUpdate:
		newValue := getFieldValueText(false, ext)
		actionTxt = i18n.Stringf(lang, "notice.updateFieldTo", fieldName, addUnitFunc(newValue))
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.deleteField", fieldName)
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromRecordManhour(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.ManhourActionExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	total := fmt.Sprintf("%v", float64(ext.Total)/100000)
	switch m.Action {
	case utilsModel.EntityActionAdd:
		i, _ := strconv.ParseInt(m.NewValue, 10, 64)
		value := fmt.Sprintf("%v", float64(i)/100000)
		actionTxt = i18n.Stringf(lang, "notice.addRecordManHour", value, total)
	case utilsModel.EntityActionUpdate:
		i, _ := strconv.ParseInt(m.NewValue, 10, 64)
		value := fmt.Sprintf("%v", float64(i)/100000)
		actionTxt = i18n.Stringf(lang, "notice.updatedRecordManHour", value, total)
	case utilsModel.EntityActionDelete:
		i, _ := strconv.ParseInt(m.OldValue, 10, 64)
		value := fmt.Sprintf("%v", float64(i)/100000)
		actionTxt = i18n.Stringf(lang, "notice.deleteRecordManHour", value, total)
	default:
	}
	return
}

func buildNoticeDescriptionFromTaskNotice(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.TaskNoticeMessageExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	distanceStr := ""
	if ext.DistanceTime.Day != 0 {
		distanceStr += fmt.Sprintf("%d天 ", ext.DistanceTime.Day)
	}
	if ext.DistanceTime.Hour != 0 {
		distanceStr += fmt.Sprintf("%d小时 ", ext.DistanceTime.Hour)
	}
	if ext.DistanceTime.Minute != 0 {
		distanceStr += fmt.Sprintf("%d分钟 ", ext.DistanceTime.Minute)
	}
	if m.Action == utilsModel.EntityActionBefore {
		actionTxt = fmt.Sprintf(i18n.String(lang, "notice_rule.before"), ext.FieldName, distanceStr)
	} else if m.Action == utilsModel.EntityActionAfter {
		actionTxt = fmt.Sprintf(i18n.String(lang, "notice_rule.after"), ext.FieldName, distanceStr)
	} else if m.Action == utilsModel.EntityActionNow || m.Action == utilsModel.EntityActionThatDay {
		actionTxt = fmt.Sprintf(i18n.String(lang, "notice_rule.nowOrThatday"), ext.FieldName)
	}
	return
}

func buildNoticeDescriptionFromTaskRelatedTask(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.UUIDNameExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	relatedTaskName := ext.Name
	switch m.Action {
	case utilsModel.EntityActionAdd:
		actionTxt = i18n.Stringf(lang, "notice.addedRelatedTask", relatedTaskName)
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.removedRelatedTask", relatedTaskName)
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskRelatedTestCase(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.RelatedTestCaseExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	relatedTestCaseName := ext.TestCaseName
	switch m.Action {
	case utilsModel.EntityActionAdd:
		actionTxt = i18n.Stringf(lang, "notice.addedRelatedTestCase", relatedTestCaseName)
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.removedRelatedTestCase", relatedTestCaseName)
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskRelatedPlanCase(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.RelatedPlanCaseExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	relatedTestCaseName := ext.TestCaseName
	switch m.Action {
	case utilsModel.EntityActionAdd:
		actionTxt = i18n.Stringf(lang, "notice.addedRelatedPlanCase", relatedTestCaseName)
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.removedRelatedPlanCase", relatedTestCaseName)
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskRelatedPlan(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.TaskRelatedTestCasePlanExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	relatedPlanName := ext.PlanName
	switch m.Action {
	case utilsModel.EntityActionAdd:
		actionTxt = i18n.Stringf(lang, "notice.addedRelatedPlan", relatedPlanName)
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.removedRelatedPlan", relatedPlanName)
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskRelatedPage(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.TaskRelatedWikiPageExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	relatedPageTitle := ext.PageTitle
	switch m.Action {
	case utilsModel.EntityActionAdd:
		actionTxt = i18n.Stringf(lang, "notice.addedTaskRelatedWikiPage", relatedPageTitle)
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.removedTaskRelatedWikiPage", relatedPageTitle)
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskDeadline(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	actionTxt = i18n.Stringf(lang, "notice.taskDeadline")
	return
}

func buildNoticeDescriptionFromTaskSubtask(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	ext := new(messageModel.SubtaskActionExt)
	if err := json.Unmarshal([]byte(m.Ext), &ext); err != nil {
		return
	}
	subTaskName := ext.Summary
	switch m.Action {
	case utilsModel.EntityActionAdd:
		actionTxt = i18n.Stringf(lang, "notice.addedSubTask", subTaskName)
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.removedSubTask", subTaskName)
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskWatchers(t *taskModel.Task, m *messageModel.Message, userMp map[string]*userModel.User) (actionTxt string) {
	switch m.Action {
	case utilsModel.EntityActionAdd:
		watcherName := GetUserNameFromUUID(m.NewValue, userMp)
		if watcherName != "" {
			actionTxt = i18n.Stringf(lang, "notice.taskWatcherAdded", watcherName)
		}
	case utilsModel.EntityActionDelete:
		watcherName := GetUserNameFromUUID(m.OldValue, userMp)
		if watcherName != "" {
			actionTxt = i18n.Stringf(lang, "notice.taskWatcherRemoved", watcherName)
		}
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskManhour(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	i, _ := strconv.ParseInt(m.NewValue, 10, 64)
	v := fmt.Sprintf("%d", i/100000)
	switch m.Action {
	case utilsModel.EntityActionUpdate:
		actionTxt = i18n.Stringf(lang, "notice.updatedAssessManHour", v)
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.String(lang, "notice.removedAssessManHour")
	default:
		//
	}
	return
}

func buildNoticeDescriptionFromTaskAttachments(t *taskModel.Task, m *messageModel.Message) (actionTxt string) {
	r, result := MustGetResource(db.BangDBMap, t.TeamUUID, m.ResourceUUID)
	if result != nil {
		return
	}
	switch m.Action {
	case utilsModel.EntityActionDelete:
		actionTxt = i18n.Stringf(lang, "notice.fileRemoved", r.Name)
	default:
		//
	}
	return
}

func createFieldValueDisplayText(fieldType int, value string) (newValue string) {
	switch fieldType {
	case fieldModel.FieldTypeInteger:
		i, _ := strconv.ParseInt(value, 10, 64)
		newValue = fmt.Sprintf("%d", i/100000)
	case fieldModel.FieldTypeFloat:
		i, _ := strconv.ParseInt(value, 10, 64)
		newValue = fmt.Sprintf("%v", float64(i)/100000)
	case fieldModel.FieldTypeDate:
		i, _ := strconv.ParseInt(value, 10, 64)
		newValue = timestamp.SecToDateString(i)
	case fieldModel.FieldTypeTime:
		i, _ := strconv.ParseInt(value, 10, 64)
		newValue = timestamp.SecToDateTimeString(i)
	case fieldModel.FieldTypeMultiOption,
		fieldModel.FieldTypeUserList,
		fieldModel.FieldTypeProduct:
		if len(value) == 0 {
			newValue = i18n.String(i18n.DefaultLang, "empty")
		} else {
			newValue = value
		}
	default:
		newValue = value
	}
	return
}

func getFieldValue(isOld bool, ext *messageModel.TaskFieldExt) string {
	if fieldModel.IsMultiOptionFieldType(ext.FieldType) {
		var options []messageModel.UUIDNameExt
		if isOld {
			options = ext.OldMultiOption
		} else {
			options = ext.NewMultiOption
		}
		ret := ""
		if options != nil {
			for _, op := range options {
				var prefix string
				if len(ret) > 0 {
					prefix = "、"
				}
				ret = ret + prefix + op.Name
			}
		}
		return ret
	}

	var v string
	if isOld {
		v = ext.OldValue
	} else {
		v = ext.NewValue
	}
	switch ext.FieldType {
	case fieldModel.FieldTypeProject, fieldModel.FieldTypeOption,
		fieldModel.FieldTypeSprint, fieldModel.FieldTypeIssueType,
		fieldModel.FieldTypeTaskStatus, fieldModel.FieldTypeUser:
		if isOld {
			if option := ext.OldOption; option != nil {
				return option.Name
			}
		} else {
			if option := ext.NewOption; option != nil {
				return option.Name
			}
		}
		return ""
	default:
		return v
	}
}

func getFieldValueText(isOld bool, ext *messageModel.TaskFieldExt) string {
	value := getFieldValue(isOld, ext)
	return createFieldValueDisplayText(ext.FieldType, value)
}

func GetUserNameFromUUID(uuidsString string, userMp map[string]*userModel.User) string {
	userUUIDs := strings.Split(uuidsString, ",")
	userNames := make([]string, 0)
	for _, userUUID := range userUUIDs {
		user, found := userMp[userUUID]
		if found {
			userNames = append(userNames, user.Name)
		} else {
			user, err := userModel.GetUserByUserUUID(db.BangDBMap, userUUID)
			if err == nil && user != nil {
				userNames = append(userNames, user.Name)
			}
		}
	}
	return strings.Join(userNames, ", ")
}

func GetResource(src gorp.SqlExecutor, teamUUID string, uuid string) (*resourceModel.Resource, error) {
	r, err := resourceModel.GetResource(src, uuid)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if r == nil || r.TeamUUID != teamUUID {
		return nil, nil
	}
	return r, nil
}

func MustGetResource(src gorp.SqlExecutor, teamUUID string, uuid string) (*resourceModel.Resource, error) {
	r, result := GetResource(src, teamUUID, uuid)
	if result != nil {
		result = errors.Trace(result)
		return nil, result
	}
	if r == nil {
		return nil, errors.AccessDeniedError(errors.Resource, "")
	}
	return r, nil
}

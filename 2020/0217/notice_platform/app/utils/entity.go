package utils

import (
	"fmt"
)

const (
	//以下为 entity 类型
	EntityTypeProject   = 1
	EntityTypeTask      = 2
	EntityTypeMessage   = 3
	EntityTypeTeam      = 4
	EntityTypeUser      = 5
	EntityTypeTemplate  = 6
	EntityTypeSprint    = 7
	EntityTypeSpace     = 8
	EntityTypeKanban    = 9
	EntityTypeVersion   = 10
	EntityTypeTestCase  = 11
	EntityTypeBatchTask = 12

	//以下为 entity 名称
	EntityTypeProjectLabel   = "project"
	EntityTypeTaskLabel      = "task"
	EntityTypeMessageLabel   = "message"
	EntityTypeTeamLabel      = "team"
	EntityTypeUserLabel      = "user"
	EntityTypeTemplateLabel  = "template"
	EntityTypeSprintLabel    = "sprint"
	EntityTypeSpaceLabel     = "space"
	EntityTypeKanbanLabel    = "kanban"
	EntityTypeVersionLabel   = "version"
	EntityTypeTestCaseLabel  = "testcase"
	EntityTypeBatchTaskLabel = "batch_task"

	EntityAttrStatus                  = 1
	EntityAttrName                    = 2
	EntityAttrSummary                 = 3
	EntityAttrMembers                 = 4
	EntityAttrDeadline                = 5
	EntityAttrAttachments             = 6
	EntityAttrAnnouncement            = 7
	EntityAttrDesc                    = 8
	EntityAttrTags                    = 9
	EntityAttrOwner                   = 10
	EntityAttrAssign                  = 11
	EntityAttrProject                 = 12
	EntityAttrPriority                = 13
	EntityAttrSubtask                 = 14
	EntityAttrTemplateUUID            = 15
	EntityAttrTarget                  = 16
	EntityAttrTemplateAttribute       = 17
	EntityAttrBoard                   = 18
	EntityAttrRelatedTask             = 19
	EntityAttrSprint                  = 20
	EntityAttrSink                    = 21
	EntityAttrPin                     = 22
	EntityAttrAssessManhour           = 23
	EntityAttrManhour                 = 24
	EntityAttrTask                    = 25
	EntityAttrGantt                   = 26
	EntityAttrPlugin                  = 27
	EntityAttrWatchers                = 28
	EntityAttrField                   = 29
	EntityAttrCodeCommit              = 30
	EntityAttrRelatedTestCase         = 31
	EntityAttrRelatedPlanCase         = 32
	EntityAttrTaskRelatedTestCasePlan = 33
	EntityAttrIssueType               = 34
	EntityAttrParent                  = 35
	EntityAttrManhours                = 36
	EntityAttrTaskNotice              = 37
	EntityAttrProduct                 = 38
	EntityAttrTaskRelatedWikiPage     = 39

	EntityAttrStatusLabel                  = "status"
	EntityAttrNameLabel                    = "name"
	EntityAttrSummaryLabel                 = "summary"
	EntityAttrMembersLabel                 = "members"
	EntityAttrDeadlineLabel                = "deadline"
	EntityAttrAttachmentsLabel             = "attachments"
	EntityAttrAnnouncementLabel            = "announcement"
	EntityAttrDescLabel                    = "desc"
	EntityAttrTagsLabel                    = "tags"
	EntityAttrOwnerLabel                   = "owner"
	EntityAttrAssignLabel                  = "assign"
	EntityAttrProjectLabel                 = "project"
	EntityAttrPriorityLabel                = "priority"
	EntityAttrSubtaskLabel                 = "subtask"
	EntityAttrTemplateUUIDLabel            = "template_uuid"
	EntityAttrTargetLabel                  = "target"
	EntityAttrTemplateAttributeLabel       = "template_attribute"
	EntityAttrBoardLabel                   = "board"
	EntityAttrRelatedTaskLabel             = "related_task"
	EntityAttrSprintLabel                  = "sprint"
	EntityAttrSinkLabel                    = "sink"
	EntityAttrPinLabel                     = "pin"
	EntityAttrAssessManhourLabel           = "assess_manhour"
	EntityAttrManhourLabel                 = "manhour"
	EntityAttrTaskLabel                    = "task"
	EntityAttrGanttLabel                   = "gantt"
	EntityAttrPluginLabel                  = "plugin"
	EntityAttrWatchersLabel                = "watchers"
	EntityAttrFieldLabel                   = "field"
	EntityAttrCodeCommitLabel              = "code_commit"
	EntityAttrRelatedTestCaseLabel         = "related_testcase"
	EntityAttrRelatedPlanCaseLabel         = "related_plan_case"
	EntityAttrTaskRelatedTestCasePlanLabel = "task_related_testcase_plan"
	EntityAttrIssueTypeLabel               = "issue_type"
	EntityAttrParentLabel                  = "parent"
	EntityAttrManhoursLabel                = "manhours"
	EntityAttrProductLabel                 = "product"
	EntityAttrTaskNoticeLabel              = "task_notice"
	EntityAttrTaskRelatedWikiPageLabel     = "task_related_wiki_page"

	EntityActionAdd     = 1
	EntityActionUpdate  = 2
	EntityActionDelete  = 3
	EntityActionRefresh = 4
	EntityActionMove    = 5
	EntityActionCopy    = 6
	EntityActionNotice  = 7
	EntityActionBefore  = 8
	EntityActionAfter   = 9
	EntityActionNow     = 10
	EntityActionThatDay = 11

	EntityActionAddLabel     = "add"
	EntityActionUpdateLabel  = "update"
	EntityActionDeleteLabel  = "delete"
	EntityActionRefreshLabel = "refresh"
	EntityActionMoveLabel    = "move"
	EntityActionCopyLabel    = "copy"
	EntityActionNoticeLabel  = "notice"
	EntityActionBeforeLabel  = "before"
	EntityActionAfterLabel   = "after"
	EntityActionNowLabel     = "now"
	EntityActionThatDayLabel = "that_day"
)

func LabelForEntityType(entityType int) (label string, err error) {
	switch entityType {
	case EntityTypeProject:
		label = EntityTypeProjectLabel
	case EntityTypeTask:
		label = EntityTypeTaskLabel
	case EntityTypeMessage:
		label = EntityTypeMessageLabel
	case EntityTypeTeam:
		label = EntityTypeTeamLabel
	case EntityTypeUser:
		label = EntityTypeUserLabel
	case EntityTypeTemplate:
		label = EntityTypeTemplateLabel
	case EntityTypeSprint:
		label = EntityTypeSprintLabel
	case EntityTypeSpace:
		label = EntityTypeSpaceLabel
	case EntityTypeKanban:
		label = EntityTypeKanbanLabel
	case EntityTypeVersion:
		label = EntityTypeVersionLabel
	case EntityTypeTestCase:
		label = EntityTypeTestCaseLabel
	case EntityTypeBatchTask:
		label = EntityTypeBatchTaskLabel
	default:
		err = fmt.Errorf("Invalid entity type: %d", entityType)
	}
	return
}

func TypeForEntityLabel(label string) (entityType int, err error) {
	switch label {
	case EntityTypeProjectLabel:
		entityType = EntityTypeProject
	case EntityTypeTaskLabel:
		entityType = EntityTypeTask
	case EntityTypeMessageLabel:
		entityType = EntityTypeMessage
	case EntityTypeTeamLabel:
		entityType = EntityTypeTeam
	case EntityTypeUserLabel:
		entityType = EntityTypeUser
	case EntityTypeTemplateLabel:
		entityType = EntityTypeTemplate
	case EntityTypeSprintLabel:
		entityType = EntityTypeSprint
	case EntityTypeSpaceLabel:
		entityType = EntityTypeSpace
	case EntityTypeKanbanLabel:
		entityType = EntityTypeKanban
	case EntityTypeVersionLabel:
		entityType = EntityTypeVersion
	case EntityTypeTestCaseLabel:
		entityType = EntityTypeTestCase
	case EntityTypeBatchTaskLabel:
		entityType = EntityTypeBatchTask
	case EntityAttrIssueTypeLabel:
		entityType = EntityAttrIssueType
	case EntityAttrParentLabel:
		entityType = EntityAttrParent
	case EntityAttrProductLabel:
		entityType = EntityAttrProduct
	default:
		err = fmt.Errorf("Invalid entity label: %s", label)
	}
	return
}

func LabelForEntityAttr(attr int) (label string, err error) {
	switch attr {
	case EntityAttrStatus:
		label = EntityAttrStatusLabel
	case EntityAttrName:
		label = EntityAttrNameLabel
	case EntityAttrSummary:
		label = EntityAttrSummaryLabel
	case EntityAttrMembers:
		label = EntityAttrMembersLabel
	case EntityAttrDeadline:
		label = EntityAttrDeadlineLabel
	case EntityAttrAttachments:
		label = EntityAttrAttachmentsLabel
	case EntityAttrAnnouncement:
		label = EntityAttrAnnouncementLabel
	case EntityAttrDesc:
		label = EntityAttrDescLabel
	case EntityAttrTags:
		label = EntityAttrTagsLabel
	case EntityAttrOwner:
		label = EntityAttrOwnerLabel
	case EntityAttrAssign:
		label = EntityAttrAssignLabel
	case EntityAttrProject:
		label = EntityAttrProjectLabel
	case EntityAttrPriority:
		label = EntityAttrPriorityLabel
	case EntityAttrSubtask:
		label = EntityAttrSubtaskLabel
	case EntityAttrTemplateUUID:
		label = EntityAttrTemplateUUIDLabel
	case EntityAttrTarget:
		label = EntityAttrTargetLabel
	case EntityAttrTemplateAttribute:
		label = EntityAttrTemplateAttributeLabel
	case EntityAttrBoard:
		label = EntityAttrBoardLabel
	case EntityAttrRelatedTask:
		label = EntityAttrRelatedTaskLabel
	case EntityAttrSprint:
		label = EntityAttrSprintLabel
	case EntityAttrSink:
		label = EntityAttrSinkLabel
	case EntityAttrPin:
		label = EntityAttrPinLabel
	case EntityAttrAssessManhour:
		label = EntityAttrAssessManhourLabel
	case EntityAttrManhour:
		label = EntityAttrManhourLabel
	case EntityAttrTask:
		label = EntityAttrTaskLabel
	case EntityAttrGantt:
		label = EntityAttrGanttLabel
	case EntityAttrPlugin:
		label = EntityAttrPluginLabel
	case EntityAttrWatchers:
		label = EntityAttrWatchersLabel
	case EntityAttrField:
		label = EntityAttrFieldLabel
	case EntityAttrCodeCommit:
		label = EntityAttrCodeCommitLabel
	case EntityAttrRelatedTestCase:
		label = EntityAttrRelatedTestCaseLabel
	case EntityAttrRelatedPlanCase:
		label = EntityAttrRelatedPlanCaseLabel
	case EntityAttrTaskRelatedTestCasePlan:
		label = EntityAttrTaskRelatedTestCasePlanLabel
	case EntityAttrIssueType:
		label = EntityAttrIssueTypeLabel
	case EntityAttrParent:
		label = EntityAttrParentLabel
	case EntityAttrManhours:
		label = EntityAttrManhoursLabel
	case EntityAttrProduct:
		label = EntityAttrProductLabel
	case EntityAttrTaskNotice:
		label = EntityAttrTaskNoticeLabel
	case EntityAttrTaskRelatedWikiPage:
		label = EntityAttrTaskRelatedWikiPageLabel
	default:
		err = fmt.Errorf("Invalid entity attribute: %d", attr)
	}
	return
}

func LabelForEntityAction(action int) (label string, err error) {
	switch action {
	case EntityActionAdd:
		label = EntityActionAddLabel
	case EntityActionUpdate:
		label = EntityActionUpdateLabel
	case EntityActionDelete:
		label = EntityActionDeleteLabel
	case EntityActionRefresh:
		label = EntityActionRefreshLabel
	case EntityActionMove:
		label = EntityActionMoveLabel
	case EntityActionCopy:
		label = EntityActionCopyLabel
	case EntityActionNotice:
		label = EntityActionNoticeLabel
	case EntityActionBefore:
		label = EntityActionBeforeLabel
	case EntityActionAfter:
		label = EntityActionAfterLabel
	case EntityActionNow:
		label = EntityActionNowLabel
	case EntityActionThatDay:
		label = EntityActionThatDayLabel
	default:
		err = fmt.Errorf("Invalid entity action: %d", action)
	}
	return
}

func ActionForEntityActionLabel(label string) (action int, err error) {
	switch label {
	case EntityActionAddLabel:
		action = EntityActionAdd
	case EntityActionUpdateLabel:
		action = EntityActionUpdate
	case EntityActionDeleteLabel:
		action = EntityActionDelete
	case EntityActionRefreshLabel:
		action = EntityActionRefresh
	case EntityActionMoveLabel:
		action = EntityActionMove
	case EntityActionCopyLabel:
		action = EntityActionCopy
	case EntityActionNoticeLabel:
		action = EntityActionNotice
	case EntityActionBeforeLabel:
		action = EntityActionBefore
	case EntityActionAfterLabel:
		action = EntityActionAfter
	case EntityActionNowLabel:
		action = EntityActionNow
	case EntityActionThatDayLabel:
		action = EntityActionThatDay
	default:
		err = fmt.Errorf("Invalid entity action: %s", label)
	}
	return
}

type Tuple2_Int_Int struct {
	Ele_1 int64 `db:"_1"`
	Ele_2 int64 `db:"_2"`
}

type Tuple2_String_Int struct {
	Ele_1 string `db:"_1"`
	Ele_2 int64  `db:"_2"`
}

type Tuple2_Int_String struct {
	Ele_1 int64  `db:"_1"`
	Ele_2 string `db:"_2"`
}

type Tuple2_String_String struct {
	Ele_1 string `db:"_1"`
	Ele_2 string `db:"_2"`
}

type Tuple3_String_Int_Int struct {
	Ele_1 string `db:"_1"`
	Ele_2 int64  `db:"_2"`
	Ele_3 int64  `db:"_3"`
}
type Tuple3_String_String_Int struct {
	Ele_1 string `db:"_1"`
	Ele_2 string `db:"_2"`
	Ele_3 int64  `db:"_3"`
}

type Tuple3_Int_Int_Int struct {
	Ele_1 int64 `db:"_1"`
	Ele_2 int64 `db:"_2"`
	Ele_3 int64 `db:"_3"`
}

type Tuple3_String_String_String struct {
	Ele_1 string `db:"_1"`
	Ele_2 string `db:"_2"`
	Ele_3 string `db:"_3"`
}

type Tuple2_String_NString struct {
	Ele_1 string  `db:"_1"`
	Ele_2 *string `db:"_2"`
}

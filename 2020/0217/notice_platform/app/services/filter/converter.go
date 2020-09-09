package filter

import (
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/issuetype"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
)

// var IssueTypeMap map[string]*
type FieldChooseEvent struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	IsSelected bool   `json:"is_selected"`
}

func ListAllFieldEvent() []*FieldChooseEvent {
	var result []*FieldChooseEvent
	for k := range IssueFieldIntergeMap {
		item := &FieldChooseEvent{
			UUID: k,
			Name: NoticeTypeNames[k],
		}
		result = append(result, item)
	}
	return result
}

func SelectedFieldToInt64(fieldList []*FieldChooseEvent) int64 {
	var sum int64
	for _, field := range fieldList {
		if v, ok := IssueFieldIntergeMap[field.UUID]; ok {
			if field.IsSelected {
				sum += v
			}
		}
	}
	return sum
}

func Int64ToSelectedField(i int64) []*FieldChooseEvent {
	var result []*FieldChooseEvent
	for k, v := range IssueFieldIntergeMap {
		item := &FieldChooseEvent{
			UUID: k,
			Name: NoticeTypeNames[k],
		}
		if i&v == v {
			item.IsSelected = true
		} else {
			item.IsSelected = false
		}
		result = append(result, item)
	}
	return result
}

func ListAllIssueTypeEvent(teamUUID string) ([]*FieldChooseEvent, error) {
	var result []*FieldChooseEvent
	issueTypeList, err := issuetype.ListTeamIssueType(db.BangDBMap, teamUUID)
	if err != nil {
		return nil, err
	}
	for _, issue := range issueTypeList {
		if issue.Status == issuetype.IssueTypeStatusNormal {
			item := &FieldChooseEvent{
				UUID: issue.UUID,
				Name: issue.Name,
			}
			result = append(result, item)
		}
	}
	return result, nil
}

func SelectedIssueTypeToInt64(teamUUID string, events []*FieldChooseEvent) (int64, error) {
	issueTypeMap, err := getIssueTypeMap(teamUUID)
	if err != nil {
		return 0, err
	}

	var sum int64
	for _, event := range events {
		if v, ok := issueTypeMap[event.UUID]; ok {
			if event.IsSelected && v.Status == issuetype.IssueTypeStatusNormal {
				sum += getIndexBit(v.Index)
			}
		}
	}
	return sum, nil
}

func Int64ToSelectedIssueType(teamUUID string, i int64) ([]*FieldChooseEvent, error) {
	issueTypeMap, err := getIssueTypeMap(teamUUID)
	if err != nil {
		return nil, err
	}

	var result []*FieldChooseEvent
	for k, v := range issueTypeMap {
		if v.Status == issuetype.IssueTypeStatusNormal {
			item := &FieldChooseEvent{
				UUID: k,
				Name: v.Name,
			}
			if i&getIndexBit(v.Index) == getIndexBit(v.Index) {
				item.IsSelected = true
			} else {
				item.IsSelected = false

			}
			result = append(result, item)
		}
	}
	return result, nil
}

func getIssueTypeMap(teamUUID string) (map[string]*issuetype.IssueType, error) {
	issueTypeList, err := issuetype.ListTeamIssueType(db.BangDBMap, teamUUID)
	if err != nil {
		return nil, err
	}
	// 刷新缓存
	SetIssueType(db.RedisClient, teamUUID, issueTypeList)
	issueTypeMap := make(map[string]*issuetype.IssueType, len(issueTypeList))
	for k, v := range issueTypeList {
		v.Index = k
		issueTypeMap[v.UUID] = v
	}
	return issueTypeMap, nil
}

func getIssueTypeInt64(teamUUID, issueTypeUUID string) (int64, error) {
	issueTypeList, err := GetIssueType(db.RedisClient, teamUUID)
	if err != nil || len(issueTypeList) == 0 {
		issueTypeList, err = issuetype.ListTeamIssueType(db.BangDBMap, teamUUID)
		if err != nil {
			return 0, err
		}
	}
	issueTypeMap := make(map[string]*issuetype.IssueType, len(issueTypeList))
	for k, v := range issueTypeList {
		v.Index = k
		issueTypeMap[v.UUID] = v
	}
	if err != nil {
		return 0, err
	}
	if v, ok := issueTypeMap[issueTypeUUID]; ok {
		return getIndexBit(v.Index), nil
	}
	return 0, nil
}

func getIndexBit(index int) int64 {
	return 1 << uint(index)
}

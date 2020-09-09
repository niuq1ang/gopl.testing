package filter

import (
	"encoding/json"
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/issuetype"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/team"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/filter"
	"gopkg.in/redis.v5"
)

const (
	RedisKeyIssueType = "issue_type"
)

var TaskFilterMap map[string]*filter.TaskFilter

func InitTaskFilterMap() error {
	taskFilterList, err := filter.ListAllTaskFilter(db.NoticeDBMap)
	if err != nil {
		return err
	}
	TaskFilterMap = make(map[string]*filter.TaskFilter, len(taskFilterList))
	for _, taskFilter := range taskFilterList {
		TaskFilterMap[taskFilter.UUID] = taskFilter
	}
	return nil
}

func InitTeamIssueType() error {
	teamList, err := team.ListAllTeam(db.BangDBMap)
	if err != nil {
		fmt.Println("teamList", err)
		return err
	}
	for _, team := range teamList {
		issueTypeList, err := issuetype.ListTeamIssueType(db.BangDBMap, team.UUID)
		if err != nil {
			fmt.Println("issueTypeList", team.UUID, err)
			return err
		}
		err = SetIssueType(db.RedisClient, team.UUID, issueTypeList)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetIssueType(client *redis.Client, teamUUID string, notice []*issuetype.IssueType) error {
	data, _ := json.Marshal(notice)
	err := client.HSet(RedisKeyIssueType, teamUUID, data).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetIssueType(client *redis.Client, teamUUID string) ([]*issuetype.IssueType, error) {
	data, err := client.HGet(RedisKeyIssueType, teamUUID).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, errors.Redis(err)
	}
	var results []*issuetype.IssueType
	_ = json.Unmarshal(data, results)
	return results, nil
}

package payload

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	messageModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/message"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/notice"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/organization"
	taskModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/task"
	userModel "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/user"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	noticeModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/notice"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/common"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils/config"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils/timestamp"
	"github.com/ngaut/log"
)

const (
	ListTypeUnread     = 1
	ListTypeRead       = 2
	ListNoticeInterval = 3 * time.Second
	MaxLimitCount      = 200
)

var nextFoundTime = int64(1578374840942336)

func TimeTicker() {
	// webhook 默认处理启动前三分钟的 message
	nextFoundTime = timestamp.Next() - 3*int64(time.Minute)
	for {
		for PayLoadMessages() {
		}
		time.Sleep(ListNoticeInterval)
	}
}

func ListMessages() ([]*messageModel.Message, error) {
	messages, err := messageModel.ListMessageWithLimit(db.BangDBMap, nextFoundTime, MaxLimitCount)
	if err != nil {
		return nil, err
	}
	if len(messages) > 0 {
		nextFoundTime = messages[len(messages)-1].SendTime
	}
	return messages, nil
}

func PayLoadMessages() bool {
	messagesList, err := ListMessages()
	if err != nil {
		log.Error(err)
		return false
	}
	for _, msg := range messagesList {
		payload, task, err := BuildMessagePayLoad(msg)
		if err != nil {
			log.Error(err)
			continue
		}
		msgBytes, _ := json.Marshal(payload)
		seqNo := timestamp.Next()
		noticePayload := &noticeModel.Notice{
			Message:               msg,
			SeqNo:                 seqNo,
			TeamUUID:              task.TeamUUID,
			TaskUUID:              task.UUID,
			FromUserUUID:          msg.FromUUID,
			MessageUUID:           msg.UUID,
			Task:                  task,
			MessagePayLoadContent: string(msgBytes),
			MessagePayLoad:        payload,
		}
		if err := noticeModel.SaveNotice(db.NoticeDBMap, noticePayload); err != nil {
			log.Error(err)
			continue
		}
		// 不太需要异步
		// noticeModel.SaveNoticeChan(noticePayload)
		SetNotice(db.RedisClient, noticePayload)
		GenerateAndSaveRocords(task.TeamUUID, seqNo, noticePayload)
	}
	return false
}

func BuildMessagePayLoad(msg *messageModel.Message) (*noticeModel.MessagePayLoad, *taskModel.Task, error) {
	noticeList, err := notice.ListNoticesByMessageUUID(db.BangDBMap, msg.UUID)
	if err != nil {
		return nil, nil, err
	}

	taskUUID, toUsers := "", []string{}
	if len(noticeList) > 0 {
		taskUUID = noticeList[0].TaskUUID
		for _, notice := range noticeList {
			toUsers = append(toUsers, notice.UserUUID)
		}
		if msg.SubjectType == utils.EntityTypeUser {
			toUsers = append(toUsers, msg.SubjectID)
		}
	} else {
		return nil, nil, errors.Errorf(errors.InvalidOption, fmt.Sprintf("message toUsers is nil, message_uuid [%s]", msg.UUID))
	}

	task, err := taskModel.GetTaskIfExists(db.BangDBMap, taskUUID)
	if err != nil {
		return nil, nil, err
	}
	if task == nil {
		return nil, nil, errors.Errorf(errors.InvalidOption, "task is nil")
	}
	var userMp map[string]*userModel.User
	userMp, err = userModel.MapUserByUUIDs(db.BangDBMap, toUsers)
	if err != nil {
		return nil, nil, err
	}
	if userMp == nil {
		return nil, nil, errors.Errorf(errors.InvalidOption, "user map is nil")
	}
	var users []*noticeModel.User
	for _, v := range userMp {
		user := &noticeModel.User{
			UUID: v.UUID,
			Name: v.Name,
		}
		users = append(users, user)
	}
	desc, ok := common.BuildNoticeDescription(task, msg, userMp)
	if !ok {
		return nil, nil, errors.Errorf(errors.InvalidOption, "build message error")

	}
	msgCreaterName := ""
	if msg.SubjectType == utils.EntityTypeUser {
		user, ok := userMp[msg.SubjectID]
		if ok {
			msgCreaterName = user.Name
		}
	}
	switch msg.Type {
	case messageModel.MessageTypeSystem:
		if msg.ObjectAttr == utils.EntityAttrTaskNotice {
			desc = fmt.Sprintf("%s", desc)
		} else {
			desc = fmt.Sprintf("%s %s", msgCreaterName, desc)
		}
	case messageModel.MessageTypeDiscussion:
		desc = fmt.Sprintf("%s:%s", msgCreaterName, desc)
	default:
	}

	teamNameMap, err := organization.MapTeamNameInOrganizationByTeamUUID(db.BangDBMap, task.TeamUUID)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}
	msgTitle := task.Summary
	if len(teamNameMap) > 1 {
		msgTitle = fmt.Sprintf("[%s]%s", teamNameMap[task.TeamUUID], msgTitle)
	}
	path := fmt.Sprintf("/team/%s/project/%s/issue_type/%s/task/%s", task.TeamUUID, task.ProjectUUID,
		task.GetIssueTypeUUID(), task.UUID)

	payload := &noticeModel.MessagePayLoad{
		ToUsers:     users,
		MessageDesc: desc,
		Title:       msgTitle,
		URL:         MakeApiURL(path),
	}
	return payload, task, nil
}

func AssembleNotice(seqNo int64) (*noticeModel.Notice, error) {
	n, err := noticeModel.ListNoticeBySeqNo(db.NoticeDBMap, seqNo)
	if err != nil {
		return nil, err
	}
	msg, err := messageModel.GetMessageByUUID(db.NoticeDBMap, n.MessageUUID)
	if err != nil {
		return nil, err
	}
	payload, task, err := BuildMessagePayLoad(msg)
	if err != nil {
		return nil, errors.Trace(err)
	}
	msgBytes, _ := json.Marshal(payload)
	noticePayload := &noticeModel.Notice{
		Message:               msg,
		SeqNo:                 seqNo,
		TeamUUID:              task.TeamUUID,
		TaskUUID:              task.UUID,
		FromUserUUID:          msg.FromUUID,
		MessageUUID:           msg.UUID,
		Task:                  task,
		MessagePayLoadContent: string(msgBytes),
		MessagePayLoad:        payload,
	}
	return noticePayload, nil
}

// 生成 Web APP 用的 url，appPath 即 # 后面的一段
func MakeApiURL(appPath string) string {
	u, err := url.Parse(config.Config.ApiBaseURL)
	if err != nil {
		log.Errorf("invalid api base url: %s", config.Config.ApiBaseURL)
	}
	u.Path = path.Clean(u.Path) + "/"
	u.Fragment = appPath
	return u.String()
}

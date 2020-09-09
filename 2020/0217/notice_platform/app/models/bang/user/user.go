package user

import (
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"gopkg.in/gorp.v1"
)

const (
	UserStatusNormal  = 1
	UserStatusDeleted = 2
	UserStatusPending = 3

	UserVerifyStatusConfirmed = 1
	UserVerifyStatusPending   = 2

	UserStatusNormalLabel  = "normal"
	UserStatusDeletedLabel = "deleted"
	UserStatusPendingLabel = "pending"
)

type User struct {
	UserProfile
	Channel          string  `db:"channel" json:"-"`
	Password         string  `db:"password" json:"-"`
	OrgUUID          string  `db:"org_uuid" json:"-"`
	SyncId           *string `db:"sync_id"`
	ThirdPartySyncID *string `db:"third_party_sync_id"`
}

type UserProfile struct {
	UUID             string  `db:"uuid" json:"uuid"`
	Email            *string `db:"email" json:"email"`
	Name             string  `db:"name" json:"name"`
	NamePinyin       string  `db:"name_pinyin" json:"name_pinyin"`
	Title            string  `db:"title" json:"title"`
	Avatar           string  `db:"avatar" json:"avatar"`
	Phone            *string `db:"phone" json:"phone"`
	TeamUUID         string  `db:"team_uuid" json:"-"`
	CreateTime       int64   `db:"create_time" json:"create_time"`
	AccessTime       int64   `db:"access_time" json:"access_time"`
	Status           int     `db:"status" json:"status"`
	VerifyStatus     int     `db:"verify_status" json:"verify_status"`
	TeamMemberStatus int     `db:"team_member_status" json:"team_member_status"`
	InviterUUID      string  `db:"inviter_uuid" json:"inviter_uuid"`
}

func GetUserByUserUUID(src gorp.SqlExecutor, userUUID string) (*User, error) {
	sql := "SELECT uuid, email, phone, status, verify_status, name, password FROM user WHERE uuid=? AND status!=?;"
	user := new(User)
	rows, err := src.Select(user, sql, userUUID, UserStatusDeleted)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		return rows[0].(*User), nil
	}
	return nil, nil
}

func MapNamesByUserUUIDs(src gorp.SqlExecutor, uuids []string) (map[string]string, error) {
	result := make(map[string]string)
	length := len(uuids)
	if length == 0 {
		return result, nil
	}
	sql := "SELECT uuid AS _1, name AS _2 FROM user WHERE uuid IN(%s) AND status!=?;"
	sql = fmt.Sprintf(sql, utils.SqlPlaceholds(length))
	args, _ := utils.BuildSqlArgs(uuids, UserStatusDeleted)
	data := make([]*utils.Tuple2_String_String, 0)
	if _, err := src.Select(&data, sql, args...); err != nil {
		return nil, errors.Sql(err)
	}
	for _, t := range data {
		result[t.Ele_1] = t.Ele_2
	}
	return result, nil
}

func MapUserByUUIDs(src gorp.SqlExecutor, uuids []string) (map[string]*User, error) {
	length := len(uuids)
	if length == 0 {
		return nil, nil
	}
	users := make([]*User, 0)
	sql := "SELECT uuid, email, name, status, verify_status FROM user WHERE uuid IN(%s);"
	sql = fmt.Sprintf(sql, utils.SqlPlaceholds(length))
	args, _ := utils.BuildSqlArgs(uuids)
	if _, err := src.Select(&users, sql, args...); err != nil {
		return nil, errors.Sql(err)
	}
	result := make(map[string]*User)
	for _, u := range users {
		result[u.UUID] = u
	}
	return result, nil
}

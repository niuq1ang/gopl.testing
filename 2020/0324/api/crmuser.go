package api

import (
	"fmt"
	"log"

	"github.com/bangwork/mygo/2020/0324/config"
)

var (
	userMap         map[string]User
	RegionalHeadMap = map[string]string{
		"FSUID_2732B63C295920A346DD9B8C48633E6C": "黄冰懿",
		"FSUID_3D7388136039B510FC633531AD5D483A": "沈娟",
		"FSUID_8D3B18EA5D209F3825CA5E7B6CD2CE6F": "罗伟",
	}
)

func GetCRMLeaderID(id string) string {
	if _, ok := RegionalHeadMap[id]; ok {
		return id
	}
	user, ok := userMap[id]
	if ok {
		_, ok := RegionalHeadMap[user.LeaderID]
		if ok {
			return user.LeaderID
		} else {
			return GetCRMLeaderID(user.LeaderID)
		}
	}
	return ""
}

const (
	RegionalHead = "区域负责人"
)

func InitUserMap() {
	userList := RequestAllDepartUser()
	userMap = make(map[string]User, len(userList))
	for _, u := range userList {
		userMap[u.OpenUserID] = u
		if u.Name == "" {
			u.Name = u.NickName
		}
		fmt.Println(fmt.Sprintf("%s-%s-%s", u.OpenUserID, u.Name, u.LeaderID))
	}
}

type UserList struct {
	BaseRes
	UserList []User `json:"userlist"`
}

type User struct {
	OpenUserID string `json:"openUserId"`
	Name       string `json:"name"`
	NickName   string `json:"nickName"`
	LeaderID   string `json:"leaderId"`
	Position   string `json:"position"`
}

func RequestAllDepartUser() []User {
	client := NewClient(config.Config.URL)
	token := GetToken()
	paramMap := map[string]interface{}{
		"corpAccessToken":   token.CorpAccessToken,
		"corpId":            token.CorpID,
		"currentOpenUserId": config.Config.AppUserId,
		"departmentId":      999999,
		"fetchChild":        true,
	}
	var result UserList
	_, err := client.Post("/cgi/user/list", paramMap, &result)
	if err != nil {
		log.Printf("request crm_simple_list error: %+v", err)
	}
	return result.UserList
}

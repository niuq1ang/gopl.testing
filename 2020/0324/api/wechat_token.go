package api

import (
	"fmt"
	"net/http"

	"github.com/bangwork/mygo/2020/0324/config"
)

var wechatToken *WechatAccessTokenPayload

const (
	invalidAccessTokenCode  = 40014
	NumberOfRequestAttempts = 3
)

var (
	contentTemplete = `立项通知
%s - 已立项
客户跟进信息已建立: %s
需求分析（BA）: @%s
销售经理: @%s
项目经理: @吴丹阳`
)

type ErrPayload struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WechatAccessTokenPayload struct {
	ErrPayload
	CurrentTime int64
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` //秒
}

type WechatUserRes struct {
	ErrPayload
	UserList []WechatUser `json:"userlist"`
}

type WechatUser struct {
	UserID     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
}

type WechatGroupChatRes struct {
	ErrPayload
	ChatID string `json:"chatid"`
}

func InitWecahtToken() error {
	var err error
	wechatToken, err = RequestWechatAccessToken()
	if err != nil {
		return err
	}
	return nil
}

func GetAccessToken() (string, error) {
	var err error
	if wechatToken == nil || wechatToken.ErrCode != 0 || GetSec()-wechatToken.CurrentTime > wechatToken.ExpiresIn-30 {
		wechatToken, err = RequestWechatAccessToken()
		if err != nil {
			return "", err
		}
		if wechatToken.ErrCode != 0 {
			return "", fmt.Errorf(wechatToken.ErrMsg)
		}
	}
	return wechatToken.AccessToken, nil
}

func RequestWechatAccessToken() (*WechatAccessTokenPayload, error) {
	_path := fmt.Sprintf("/cgi-bin/gettoken?corpid=%s&corpsecret=%s", config.Config.CorpID, config.Config.CorpSecret)
	client := NewClient(config.Config.WechatBaseURL)
	resp := new(WechatAccessTokenPayload)
	_, err := client.Request(http.MethodGet, _path, nil, resp)
	if err != nil {
		return nil, err
	}
	resp.CurrentTime = GetSec()
	return resp, nil
}

func RequestWechatUserList() (map[string]string, error) {
	token, err := GetAccessToken()
	if err != nil {
		InitWecahtToken()
		token = wechatToken.AccessToken
	}
	userMap := make(map[string]string)
	resp := new(WechatUserRes)
	for i := 0; i < NumberOfRequestAttempts; i++ {
		// department_id=3 服务部
		_path := fmt.Sprintf("/cgi-bin/user/simplelist?access_token=%s&department_id=%s&fetch_child=1", token, "3")
		client := NewClient(config.Config.WechatBaseURL)
		_, err = client.Request(http.MethodGet, _path, nil, resp)
		if err != nil || resp != nil && resp.ErrCode != 0 {
			InitWecahtToken()
			token = wechatToken.AccessToken
			continue
		} else {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	if resp.ErrCode != 0 {
		return nil, fmt.Errorf(wechatToken.ErrMsg)
	}
	for _, user := range resp.UserList {
		userMap[user.Name] = user.UserID
	}
	return userMap, nil
}

func CreateGroupChat(stage, customerName string, users []string) (string, error) {
	token, err := GetAccessToken()
	if err != nil {
		InitWecahtToken()
		token = wechatToken.AccessToken
	}
	name := fmt.Sprintf("KA服务-%s-%s", stage, customerName)
	// users = append(users, "dare", "YeShenYue")
	users = append(users, "ADu", "jim", "ZouYuHang", "jasonfong", "dare")

	var groupUser []string
	for _, user := range users {
		if user != "" {
			groupUser = append(groupUser, user)
		}
	}
	resp := new(WechatGroupChatRes)
	for i := 0; i < NumberOfRequestAttempts; i++ {
		_path := fmt.Sprintf("/cgi-bin/appchat/create?access_token=%s", token)
		client := NewClient(config.Config.WechatBaseURL)
		req := map[string]interface{}{
			"name":     name,
			"owner":    "ADu", // 群主
			"userlist": groupUser,
		}
		_, err = client.Request(http.MethodPost, _path, req, resp)
		if err != nil || resp != nil && resp.ErrCode != 0 {
			InitWecahtToken()
			token = wechatToken.AccessToken
			continue
		} else {
			break
		}
	}
	if err != nil {
		return "", err
	}
	if resp.ErrCode != 0 {
		return "", fmt.Errorf(wechatToken.ErrMsg)
	}
	return resp.ChatID, nil
}

func WechatSendGroupChatText(chatID, customerName, wikiURL, ba string, ownersName []string) {
	token, err := GetAccessToken()
	if err != nil {
		InitWecahtToken()
		token = wechatToken.AccessToken
	}
	// 添加默认成员
	resp := new(WechatGroupChatRes)
	content := fmt.Sprintf(contentTemplete, customerName, wikiURL, ba, ArrayToStringLine(ownersName))
	for i := 0; i < NumberOfRequestAttempts; i++ {
		_path := fmt.Sprintf("/cgi-bin/appchat/send?access_token=%s", token)
		client := NewClient(config.Config.WechatBaseURL)
		req := map[string]interface{}{
			"chatid":  chatID,
			"msgtype": "text",
			"text": struct {
				Content string `json:"content"`
			}{Content: content},
		}
		_, err = client.Request(http.MethodPost, _path, req, resp)
		if err != nil || resp != nil && resp.ErrCode != 0 {
			InitWecahtToken()
			token = wechatToken.AccessToken
			continue
		} else {
			break
		}
	}
	if err != nil {
		fmt.Printf("\nsend group chat text err: %v", err)
	}
}

func ArrayToStringLine(array []string) string {
	if len(array) == 0 {
		return ""
	}
	var name string
	for i, n := range array {
		if i == 0 {
			name = n
		} else {
			name += ", " + n
		}
	}
	return name
}

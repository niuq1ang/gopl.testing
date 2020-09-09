package api

import (
	"fmt"
	"net/http"

	"github.com/bangwork/mygo/2020/0324/config"
)

const (
// teamUUID = "RDjYMhKq"
// userID   = ""
)

var (
	token    = "vTyeWrCaOKRD96FGhExrChfekRChscyK2bzyEqLBapWkOYQfnmbSSwKyvV5g5m8m"
	userUUID = "5ZPerY1K"
	teamUUID = "RDjYMhKq"
)

func GetONESToken() error {
	user, err := LoginONES()
	if err != nil {
		return err
	}
	userUUID = user.UUID
	token = user.Token
	return nil
}

type ONESUser struct {
	UUID         string `json:"uuid"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	NamePinyin   string `json:"name_pinyin"`
	Title        string `json:"title"`
	Avatar       string `json:"avatar"`
	Phone        string `json:"phone"`
	CreateTime   int64  `json:"create_time"`
	Status       int    `json:"status"`
	Channel      string `json:"channel"`
	Token        string `json:"token"`
	LicenseTypes []int  `json:"license_types"`
}

type ONESLogin struct {
	User ONESUser `json:"user"`
}

func LoginONES() (*ONESUser, error) {
	_path := fmt.Sprintf("/project/v1/auth/login")
	client := NewClient(config.Config.ONESBaseURL)

	req := map[string]interface{}{
		"email":    "robot@ones.ai",
		"password": "Njw7DzNi5Ca",
	}
	resp := new(ONESLogin)
	_, err := client.Request(http.MethodPost, _path, req, resp)
	if err != nil {
		return nil, err
	}
	return &resp.User, nil
}

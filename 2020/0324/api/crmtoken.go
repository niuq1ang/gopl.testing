package api

import (
	"log"
	"time"

	"github.com/bangwork/mygo/2020/0324/config"
)

var crmToken CropAccessTokenRes

type BaseRes struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type CropAccessTokenRes struct {
	BaseRes
	CorpAccessToken string `json:"corpAccessToken"`
	CorpID          string `json:"corpId"`
	ExpiresIn       int64  `json:"expiresIn"`
	LastRequestTime int64
}

type ObjectListResult struct {
	ErrorCode    int       `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
	Objects      []Objects `json:"objects"`
}

type Objects struct {
	APIName     string `json:"api_name"`
	DisplayName string `json:"display_name"`
}

// GetToken 过期时间剩余5分钟尝试自动刷新token
func GetToken() CropAccessTokenRes {
	if crmToken.CorpAccessToken == "" || crmToken.ErrorCode != 0 || GetSec()-crmToken.LastRequestTime >= crmToken.ExpiresIn-10*60 {
		crmToken = RequestCorpAccessToken()
	}
	return crmToken
}

// RequestCorpAccessToken 获取AccessToken
func RequestCorpAccessToken() CropAccessTokenRes {
	client := NewClient(config.Config.URL)
	paramMap := map[string]interface{}{
		"appId":         config.Config.AppID,
		"appSecret":     config.Config.AppSecret,
		"permanentCode": config.Config.PermanentCode,
	}

	var result CropAccessTokenRes
	_, err := client.Post("/cgi/corpAccessToken/get/V2", paramMap, &result)
	if err != nil {
		log.Fatalf("request crm token error: %+v", err)
	}
	result.LastRequestTime = GetSec()
	return result
}

func GetSec() int64 {
	return time.Now().Unix()
}

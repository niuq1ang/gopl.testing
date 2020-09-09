package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bangwork/mygo/2020/0324/config"
)

var (
	line           = "<tr>\n\t\t\t<td>{{number}}</td>\n\t\t\t<td>{{demand_detail}}</td>\n\t\t\t<td>{{create_time}}</td>\n\t\t\t<td>-</td>\n\t\t</tr>"
	demandTemplete = "<table style=\"width:801px\">\n\t<colgroup>\n\t\t<col style=\"width:10%\" />\n\t\t<col style=\"width:66.5%\" />\n\t\t<col style=\"width:13%\" />\n\t\t<col style=\"width:10.5%\" />\n\t</colgroup>\n\t<thead>\n\t\t<tr>\n\t\t\t<th>序号</th>\n\t\t\t<th>描述</th>\n\t\t\t<th>创建时间</th>\n\t\t\t<th>相关图片</th>\n\t\t</tr>\n\t</thead>\n\t<tbody>\n\t\t{{trs}}\n\t</tbody>\n</table>\n\n<p>&nbsp;</p>\n"
	wikiTemplete   = "<h3 data-oid=\"zzvi51bp\">负责人信息</h3>\n\n<table style=\"width:632px\">\n\t<colgroup>\n\t\t<col style=\"width:18.7005%\" />\n\t\t<col style=\"width:81.2995%\" />\n\t</colgroup>\n\t<thead>\n\t\t<tr>\n\t\t\t<th>角色</th>\n\t\t\t<th>负责人</th>\n\t\t</tr>\n\t</thead>\n\t<tbody>\n\t\t<tr>\n\t\t\t<td>项目经理</td>\n\t\t\t<td>吴丹阳</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>BA</td>\n\t\t\t<td>{{ba}}</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>销售经理</td>\n\t\t\t<td>{{sales_manager}}</td>\n\t\t</tr>\n\t</tbody>\n</table>\n\n<h3 data-oid=\"pb7zdpdu\">客户跟进记录</h3>\n\n<table style=\"width:638px\">\n\t<colgroup>\n\t\t<col style=\"width:18.2104%\" />\n\t\t<col style=\"width:12.0879%\" />\n\t\t<col style=\"width:15.0706%\" />\n\t\t<col style=\"width:54.6311%\" />\n\t</colgroup>\n\t<thead>\n\t\t<tr>\n\t\t\t<th>时间</th>\n\t\t\t<th>事项</th>\n\t\t\t<th>参与人</th>\n\t\t\t<th>概要描述（信息过多时请单开页面，此处附链接）</th>\n\t\t</tr>\n\t</thead>\n\t<tbody>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>线上沟通</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>线下拜访</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>...</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t</tbody>\n</table>\n\n<p>&nbsp;</p>\n\n<p>&nbsp;</p>\n"
	spaceUUID      = "TF34tMF1"
	ParentPageUUID = "QrDb4iUw"
	pageUUIDMap    = map[string]string{
		"":    "QrDb4iUw",
		"黄冰懿": "TW22NB9P",
		"沈娟":  "91M4aXEx",
		"罗伟":  "WRmpamAx",
	}
)

// {"uuid":"65kVDu4r","space_uuid":"TF34tMF1","page_uuid":"QrDb4iUw","owner_uuid":"5ZPerY1K","status":1,"create_time":1585126749}
type DraftRes struct {
	UUID       string `json:"uuid"`
	SpaceUUID  string `json:"space_uuid"`
	PageUUID   string `json:"page_uuid"`
	OwnerUUID  string `json:"owner_uuid"`
	Status     int    `json:"status"`
	CreateTime int    `json:"create_time"`
}

func CreateWiki(stage, customerName, ba, districtManagerName string, ownersName []string, demands []DemandDetailByProjectRes) (string, error) {
	var wikiAccessURL string
	content := strings.Replace(wikiTemplete, "{{ba}}", ba, -1)
	content = strings.Replace(content, "{{sales_manager}}", ArrayToStringLine(ownersName), -1)
	title := fmt.Sprintf("[%s]%s", stage, customerName)
	draft, err := CreateDraft(title, pageUUIDMap[districtManagerName])
	if err != nil {
		return "", err
	}
	draft, err = PublishedPage(title, content, draft)
	if err != nil {
		return "", err
	}
	wikiAccessURL = fmt.Sprintf("https://ones.ai/wiki/#/team/%s/space/%s/page/%s", teamUUID, spaceUUID, draft.PageUUID)
	if len(demands) == 0 {
		return wikiAccessURL, nil
	}
	var tr string
	var trs string
	for i, demand := range demands {
		if i == 0 {
			tr = strings.Replace(line, "{{number}}", strconv.Itoa(i), -1)
			tr = strings.Replace(tr, "{{demand_detail}}", demand.DemandDetail, -1)
			tr = strings.Replace(tr, "{{create_time}}", msToTimeString(demand.CreateTime), -1)
		} else {
			tr = strings.Replace("\n\t\t"+line, "{{number}}", strconv.Itoa(i), -1)
			tr = strings.Replace(tr, "{{demand_detail}}", demand.DemandDetail, -1)
			tr = strings.Replace(tr, "{{create_time}}", msToTimeString(demand.CreateTime), -1)
		}
		trs += tr
	}
	content = strings.Replace(demandTemplete, "{{trs}}", trs, -1)
	title = fmt.Sprintf("[%s]需求列表", customerName)
	draft, err = CreateDemandDraft(title, draft.PageUUID)
	if err != nil {
		return wikiAccessURL, err
	}
	draft, err = PublishedPage(title, content, draft)
	if err != nil {
		return wikiAccessURL, err
	}
	return wikiAccessURL, nil
}

// {"copy_src_type":"template","copy_src_uuid":"6taiwtXY","page_uuid":"QrDb4iUw","status":1}
func CreateDraft(title string, pageUUID string) (*DraftRes, error) {
	_path := fmt.Sprintf("/wiki/v1/team/%s/space/%s/drafts/add", teamUUID, spaceUUID)
	client := NewClient(config.Config.ONESBaseURL)
	req := map[string]interface{}{
		"copy_src_type": "template",
		"copy_src_uuid": "6taiwtXY",
		"page_uuid":     pageUUID,
		"title":         title,
		"status":        1,
	}
	resp := new(DraftRes)
	_, err := client.Request(http.MethodPost, _path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateDemandDraft(title, pageUUID string) (*DraftRes, error) {
	_path := fmt.Sprintf("/wiki/v1/team/%s/space/%s/drafts/add", teamUUID, spaceUUID)
	client := NewClient(config.Config.ONESBaseURL)
	req := map[string]interface{}{
		"copy_src_type": "global_template",
		"copy_src_uuid": "PcHwAMii",
		"page_uuid":     pageUUID,
		"title":         title,
		"status":        1,
	}
	resp := new(DraftRes)
	_, err := client.Request(http.MethodPost, _path, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// {
// 	"uuid": "65kVDu4r",
// 	"space_uuid": "TF34tMF1",
// 	"page_uuid": "QrDb4iUw",
// 	"from_version": -1,
// 	"title": "[k7] test",
// 	"content": "<h3 data-oid=\"zzvi51bp\">负责人信息</h3>\n\n<table style=\"width:632px\">\n\t<colgroup>\n\t\t<col style=\"width:18.7005%\" />\n\t\t<col style=\"width:81.2995%\" />\n\t</colgroup>\n\t<thead>\n\t\t<tr>\n\t\t\t<th>角色</th>\n\t\t\t<th>负责人</th>\n\t\t</tr>\n\t</thead>\n\t<tbody>\n\t\t<tr>\n\t\t\t<td>项目经理</td>\n\t\t\t<td>杜生军</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>BA</td>\n\t\t\t<td>方志恒/龙宇双/刘晶晶</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>销售经理</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t</tbody>\n</table>\n\n<h3 data-oid=\"pb7zdpdu\">客户跟进记录</h3>\n\n<table style=\"width:638px\">\n\t<colgroup>\n\t\t<col style=\"width:18.2104%\" />\n\t\t<col style=\"width:12.0879%\" />\n\t\t<col style=\"width:15.0706%\" />\n\t\t<col style=\"width:54.6311%\" />\n\t</colgroup>\n\t<thead>\n\t\t<tr>\n\t\t\t<th>时间</th>\n\t\t\t<th>事项</th>\n\t\t\t<th>参与人</th>\n\t\t\t<th>概要描述（信息过多时请单开页面，此处附链接）</th>\n\t\t</tr>\n\t</thead>\n\t<tbody>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>线上沟通</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>线下拜访</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>...</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t\t<tr>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t\t<td>&nbsp;</td>\n\t\t</tr>\n\t</tbody>\n</table>\n\n<p>&nbsp;</p>\n\n<p>&nbsp;</p>\n",
// 	"status": 1,
// 	"create_time": 1000000000,
// 	"updated_time": 1000000000,
// 	"is_published": true
// }
func PublishedPage(title, contents string, draft *DraftRes) (*DraftRes, error) {
	_path := fmt.Sprintf("/wiki/v1/team/%s/space/%s/draft/%s/update", teamUUID, spaceUUID, draft.UUID)
	client := NewClient(config.Config.ONESBaseURL)
	req := map[string]interface{}{
		"uuid":         draft.UUID,
		"space_uuid":   draft.SpaceUUID,
		"content":      contents,
		"title":        title,
		"page_uuid":    draft.PageUUID,
		"status":       1,
		"from_version": -1,
		"is_published": true,
	}
	resp := new(DraftRes)
	_, err := client.Request(http.MethodPost, _path, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GotaRandomBA() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(4096)
	return config.Config.BAUsers[n%len(config.Config.BAUsers)]
}

func msToTimeString(ms int64) string {
	tm := time.Unix(ms/1000, 0)
	return tm.Format("2006-01-02 15:04:05")
}

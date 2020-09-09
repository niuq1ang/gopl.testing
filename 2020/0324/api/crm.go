package api

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/bangwork/mygo/2020/0324/config"
)

// create ONES wiki title：[阶段] <客户名称>

const (
	NewOpportunityObj = "NewOpportunityObj"
	// 项目需求明细 对象名称
	ProjectDemandDetail = "object_T91oR__c"
	// 项目立项 对象名称
	ProjectApprovalName = "object_KRt7b__c"
	AccountObjName      = "AccountObj"
)

var RecordMap map[string]bool

func init() {
	RecordMap = make(map[string]bool)
}

var (
	regStage = regexp.MustCompile(`[A-Z][0-9]{1,2}`)
)

// SalesSatgeText 商机阶段
var SalesSatgeText = map[string]string{
	"OF1126j8j": "M5初步沟通",
	"351hOhy8F": "M4产品演示",
	"aHkr8t47M": "M3跟进试用",
	"PrFu9FrIp": "M2合同",
	"pQvB82jp9": "M1签约",
	"mdpyytUzv": "M0回款确认",
	"qjJ91UcUl": "K10初步接触",
	"5132E0SmH": "K9提案",
	"bq1JOTHIc": "K8试用|POC",
	"zhpNupmDm": "K7招投标",
	"8V98rd6z1": "K6合同",
	"bKi9o5mfT": "K5签约",
	"xkuXA89F6": "K4首款确认",
	"L74B9oJpF": "K3首次交付",
	"iy1208oth": "K2持续开发",
	"Cpf0t1woe": "K1二次交付",
	"x5c1CP627": "K0尾款赢单",
	"1":         "验证客户",
	"2":         "需求确定",
	"3":         "方案/报价",
	"4":         "谈判审核",
	"5":         "赢单",
	"6":         "输单",
	"7":         "无效",
}

// 项目立项
type ProjectApproval struct {
	// 项目立项ID
	ID string `json:"_id"`
	// 项目标号
	Name string `json:"name"`
	// 客户名称
	CustomerID string `json:"field_r9A1x__c"`
	// 商机名称
	OpportunityName string `json:"field_av94E__c"`
	// 负责人
	Owner []string `json:"owner"`
	// 负责人主属部门
	OwnerDepartment string `json:"owner_department"`
}

// CRM 传过来的值
type CRMData struct {
	UserID   string          `json:"userId"`
	TenantID string          `json:"tenantId"`
	Data     ProjectApproval `json:"data"`
	// Details  interface{}     `json:"details"`
}

func HandlerCRMRequest(data []byte) error {
	var crmData CRMData
	_ = json.Unmarshal(data, &crmData)

	if RecordMap[crmData.Data.ID] {
		log.Println("crmData.Data.ID duplicated")
		return nil
	} else {
		RecordMap[crmData.Data.ID] = true
	}

	opportiunity := QueryNewOpportunityStageByID(crmData.Data.OpportunityName)
	var stage string
	if opportiunity.Data.SaleStage == "" {
		stage = "未找到"
	} else {
		stage = SalesSatgeText[opportiunity.Data.SaleStage]
		if len(opportiunity.Data.SaleStage) > 1 {
			stage = regStage.FindString(stage)
		}
	}

	projectApproval := QueryProjectApprovalByID(crmData.Data.ID)
	accountObj := QueryAccountObjByID(crmData.Data.CustomerID)
	customerName := accountObj.Data.Name
	// 刷新 crm user map
	InitUserMap()
	// owners 销售经理ID, ownersName 销售经理 name, districtManager 大区经理ID, districtManagerName 大区经理 name
	owners, ownersName, districtManager, districtManagerName := projectApproval.Data.Owner, []string{}, "", ""
	for _, owner := range owners {
		if u, ok := userMap[owner]; ok {
			ownersName = append(ownersName, u.Name)
			districtManager = GetCRMLeaderID(owner)
			districtManagerName = RegionalHeadMap[districtManager]
		}
	}
	// 获取客户需求
	demands := QueryDemandDetailByProjectID(crmData.Data.ID)
	// 刷新 ONES token
	GetONESToken()
	// 创建 wiki 页面
	ba := GotaRandomBA()
	wikiURL, err := CreateWiki(stage, customerName, ba, districtManagerName, ownersName, demands)
	if err != nil {
		fmt.Printf("\ncreate wiki err: %v", err)
	}
	// 添加 project 自定义工作项属性值
	if err := AddCustomerField(customerName); err != nil {
		fmt.Printf("\nadd project field options err: %v", err)
	}

	// 获取企业微信用户
	wechatUsers, err := RequestWechatUserList()
	if err != nil {
		fmt.Printf("\nget wechat users err: %v", err)
		return err
	}
	// wechatOwners wechat 销售经理ID, wechatOwners wechat 大区经理ID
	wechatOwners := []string{}
	for _, n := range ownersName {
		if u, ok := wechatUsers[n]; ok {
			wechatOwners = append(wechatOwners, u)
		}
	}
	// add 大区经理 ba
	wechatOwners = append(wechatOwners, wechatUsers[districtManagerName], wechatUsers[ba])
	// 创建微信群
	chatID, err := CreateGroupChat(stage, customerName, wechatOwners)
	if err != nil {
		fmt.Printf("\ncreate wechat group chat err: %v", err)
		return err
	}
	WechatSendGroupChatText(chatID, customerName, wikiURL, ba, ownersName)
	return nil
}

type DemandDetailByProjectRes struct {
	// 需求编号
	DemandID string `json:"name"`
	// 需求明细
	DemandDetail string `json:"field_2vowR__c"`
	// 需求阶段
	DemandStage string `json:"field_3ku25__c"`
	// 负责人
	Owner []string `json:"owner"`
	// 项目编号
	ProjectID string `json:"field_idnYm__c"`
	// 相关图片
	// [{"ext":"jpeg","path":"6f8ce53b-2a19-4fd4-ae59-b3c9f2025f38"}]
	// Photo []byte `json:"field_2FP8a__c"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
}

// 通过项目编号查找项目需求明细
func QueryDemandDetailByProjectID(projectID string) []DemandDetailByProjectRes {
	var demandList []DemandDetailByProjectRes
	searchInfo := SearchQueryInfo{
		Limit:  50,
		Offset: 0,
		Filters: []map[string]interface{}{
			map[string]interface{}{
				"field_name":   "field_idnYm__c", // 生命状态 不等于作废
				"field_values": []string{projectID},
				"operator":     "EQ",
			},
		},
		FieldProjection: []string{
			"name",
			"field_2vowR__c",
			"field_3ku25__c",
			"owner",
			"field_idnYm__c",
			"field_2FP8a__c",
			"create_time",
		},
	}
	result := queryDefined(searchInfo, ProjectDemandDetail)
	if result.ErrorCode != 0 {
		log.Printf("Call OpenAPI Error, errCode=[%d], errMsg=%s", result.ErrorCode, result.ErrorMessage)
		return demandList
	}
	for _, raw := range result.Data.DataList {
		var result DemandDetailByProjectRes
		err := json.Unmarshal(raw, &result)
		if err != nil {
			fmt.Println(err)
			continue
		}
		demandList = append(demandList, result)
	}
	return demandList
}

// 项目立项返回值
type ProjectApprovalRes struct {
	BaseRes
	Data ProjectApproval `json:"data"`
}

// 查找项目立项信息
func QueryProjectApprovalByID(id string) ProjectApprovalRes {
	var results ProjectApprovalRes
	client := NewClient(config.Config.URL)
	token := GetToken()
	paramMap := map[string]interface{}{
		"corpAccessToken":   token.CorpAccessToken,
		"corpId":            token.CorpID,
		"currentOpenUserId": config.Config.AppUserId,
		"data": map[string]interface{}{
			"objectDataId":      id,
			"dataObjectApiName": ProjectApprovalName,
		},
	}
	_, err := client.Post("/cgi/crm/custom/data/get", paramMap, &results)
	if err != nil {
		log.Printf("query data err: %+v", err)
	}
	return results
}

// 商机返回字段
type NewOpportunityRes struct {
	BaseRes
	Data NewOpportunityData `json:"data"`
}

type NewOpportunityData struct {
	TeamUUID string `json:"field_vnWq1__c"`
	// 商机名称
	Name string `json:"name"`
	// 阶段状态
	SaleStatus string `json:"sales_status"`
	// 商机阶段
	SaleStage string `json:"sales_stage"`
}

// 查找商机
func QueryNewOpportunityStageByID(id string) NewOpportunityRes {
	var results NewOpportunityRes
	client := NewClient(config.Config.URL)
	token := GetToken()
	paramMap := map[string]interface{}{
		"corpAccessToken":   token.CorpAccessToken,
		"corpId":            token.CorpID,
		"currentOpenUserId": config.Config.AppUserId,
		"data": map[string]interface{}{
			"objectDataId":      id,
			"dataObjectApiName": NewOpportunityObj,
		},
	}
	_, err := client.Post("/cgi/crm/v2/data/get", paramMap, &results)
	if err != nil {
		log.Printf("query data err: %+v", err)
	}
	return results
}

// 客户返回字段
type AccountObjRes struct {
	BaseRes
	Data NewOpportunityData `json:"data"`
}

type AccountObj struct {
	TeamUUID string `json:"UDSText2__c"`
	// 客户名称
	Name string `json:"name"`
}

// 查找客户
func QueryAccountObjByID(id string) AccountObjRes {
	var results AccountObjRes
	client := NewClient(config.Config.URL)
	token := GetToken()
	paramMap := map[string]interface{}{
		"corpAccessToken":   token.CorpAccessToken,
		"corpId":            token.CorpID,
		"currentOpenUserId": config.Config.AppUserId,
		"data": map[string]interface{}{
			"objectDataId":      id,
			"dataObjectApiName": AccountObjName,
		},
	}
	_, err := client.Post("/cgi/crm/v2/data/get", paramMap, &results)
	if err != nil {
		log.Printf("query data err: %+v", err)
	}
	return results
}

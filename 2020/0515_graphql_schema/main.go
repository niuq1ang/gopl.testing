package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"unicode"
)

type DataBody struct {
	Data FieldBody `json:"data"`
}

type FieldBody struct {
	Field []FieldItem `json:"fields"`
}
type FieldItem struct {
	Aliases      []string    `json:"aliases"`
	UUID         string      `json:"uuid"`
	FieldType    string      `json:"fieldType"`
	Name         string      `json:"name"`
	Required     bool        `json:"required"`
	DefaultValue interface{} `json:"defaultValue"`
	AllowEmpty   bool        `json:"allowEmpty"`
	CanUpdate    bool        `json:"canUpdate"`
}

var (
	Aliases   = "aliases"
	FieldType = "type"
	Required  = "required"
	CanUpdate = "canUpdate"
	// AllowEmpty  = "allowEmpty"
	Description = "description"
)

func main() {
	var result DataBody
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		panic(err)
	}

	lenMap := initLenMap(result)
	for i, fieldItem := range result.Data.Field {
		if i == 0 {
			// fmt.Printf("|aliases|fieldType|required|canUpdate|allowEmpty|description|\n")
			// fmt.Printf("|%s|%s|%s|%s|%s|%s|\n", printString(Aliases, " ", lenMap[Aliases][0]), printString(FieldType, " ", lenMap[FieldType][0]), printString(Required, " ", lenMap[Required][0]), printString(CanUpdate, " ", lenMap[CanUpdate][0]), printString(AllowEmpty, " ", lenMap[AllowEmpty][0]), printString(Description, " ", lenMap[Description][0]))
			// fmt.Printf("|%s|%s|%s|%s|%s|%s|\n", printAlign(lenMap[Aliases][0]), printAlign(lenMap[FieldType][0]), printAlign(lenMap[Required][0]), printAlign(lenMap[CanUpdate][0]), printAlign(lenMap[AllowEmpty][0]), printAlign(lenMap[Description][0]))
			fmt.Printf("|%s|%s|%s|%s|%s|\n", printString(Aliases, " ", lenMap[Aliases][0]), printString(FieldType, " ", lenMap[FieldType][0]), printString(Required, " ", lenMap[Required][0]), printString(CanUpdate, " ", lenMap[CanUpdate][0]), printString(Description, " ", lenMap[Description][0]))
			fmt.Printf("|%s|%s|%s|%s|%s|\n", printAlign(lenMap[Aliases][0]), printAlign(lenMap[FieldType][0]), printAlign(lenMap[Required][0]), printAlign(lenMap[CanUpdate][0]), printAlign(lenMap[Description][0]))
		}
		// fmt.Printf("|%s|%s|%s|%s|%s|%s|\n", printString(fieldItem.Aliases[0], " ", lenMap[Aliases][0]), printString(fieldItem.FieldType, " ", lenMap[FieldType][0]), printString(boolToText(fieldItem.Required), " ", lenMap[Required][0]), printString(boolToText(fieldItem.CanUpdate), " ", lenMap[CanUpdate][0]), printString(boolToText(fieldItem.AllowEmpty), " ", lenMap[AllowEmpty][0]), printString(fieldItem.Name, " ", lenMap[Description][0]))
		fmt.Printf("|%s|%s|%s|%s|%s|\n", printString(fieldItem.Aliases[0], " ", lenMap[Aliases][0]), printString(fieldItem.FieldType, " ", lenMap[FieldType][0]), printString(boolToText(fieldItem.Required), " ", lenMap[Required][0]), printString(boolToText(fieldItem.CanUpdate), " ", lenMap[CanUpdate][0]), printString(fieldItem.Name, " ", lenMap[Description][0]))

	}
}

func initLenMap(data DataBody) map[string][]int {
	lenMap := make(map[string][]int, 0)
	lenMap[Aliases] = []int{len(Aliases)}
	lenMap[FieldType] = []int{len(FieldType)}
	lenMap[Required] = []int{len(Required)}
	lenMap[CanUpdate] = []int{len(CanUpdate)}
	// lenMap[AllowEmpty] = []int{len(AllowEmpty)}
	lenMap[Description] = []int{len(Description)}
	for _, fieldItem := range data.Data.Field {
		lenMap[Aliases] = append(lenMap[Aliases], lenString(fieldItem.Aliases[0]))
		lenMap[FieldType] = append(lenMap[FieldType], lenString(fieldItem.FieldType))
		lenMap[Required] = append(lenMap[Required], 1)
		lenMap[CanUpdate] = append(lenMap[CanUpdate], 1)
		// lenMap[AllowEmpty] = append(lenMap[AllowEmpty], 1)
		lenMap[Description] = append(lenMap[Description], lenString(fieldItem.Name))

	}

	for k, v := range lenMap {
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		lenMap[k] = v
	}
	return lenMap
}

func lenString(source string) int {
	length := 0
	for _, s := range source {
		if unicode.Is(unicode.Han, s) {
			length += 2
		} else {
			length += 1
		}
	}
	return length
}

func printString(source, placeholder string, length int) string {
	if lenString(source) == length {
		return source
	}
	replaceHolder := ""
	for i := lenString(source); i < length; i++ {
		replaceHolder += placeholder

	}
	return source + replaceHolder
}

func printAlign(length int) string {
	placeholder := "-"
	source := " :---"
	replaceHolder := ""
	for i := lenString(source); i < length-1; i++ {
		replaceHolder += placeholder

	}
	return source + replaceHolder + " "
}

func boolToText(b bool) string {
	if b {
		return "T"
	}
	return "F"
}

func printDefaultValue(x interface{}) string {
	if x == nil {
		return "null"
	}
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case int64:
		return strconv.FormatInt(x, 10)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array、chan、func、map、pointer、slice、struct
		return "???" + reflect.ValueOf(x).Type().String()
	}
}

var data = `
{
    "data": {
        "fields": [
            {
                "aliases": [
                    "uuid"
                ],
                "allowEmpty": false,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "text",
                "name": "[UUID]",
                "options": [],
                "required": false,
                "uuid": "uuid"
            },
            {
                "aliases": [
                    "item_type"
                ],
                "allowEmpty": false,
                "builtIn": true,
                "defaultValue": "task",
                "fieldType": "text",
                "name": "[Item类型]",
                "options": [],
                "required": true,
                "uuid": ""
            },
            {
                "aliases": [
                    "key"
                ],
                "allowEmpty": false,
                "builtIn": true,
                "defaultValue": "task",
                "fieldType": "text",
                "name": "[Item Key]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "name",
                    "summary",
                    "title"
                ],
                "allowEmpty": false,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "text",
                "name": "任务名称",
                "options": [],
                "required": true,
                "uuid": ""
            },
            {
                "aliases": [
                    "name_pinyin"
                ],
                "allowEmpty": false,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "text",
                "name": "[名称拼音]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "owner"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "user",
                "name": "任务创建人",
                "options": [],
                "required": false,
                "uuid": "field003"
            },
            {
                "aliases": [
                    "assign"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "user",
                "name": "任务负责人",
                "options": [],
                "required": true,
                "uuid": "field004"
            },
            {
                "aliases": [
                    "project"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "project",
                "name": "所属项目",
                "options": [],
                "required": true,
                "uuid": "field006"
            },
            {
                "aliases": [
                    "sprint"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "sprint",
                "name": "所属迭代",
                "options": [],
                "required": true,
                "uuid": "field011"
            },
            {
                "aliases": [
                    "issue_type"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "issue_type",
                "name": "所属任务类型",
                "options": [],
                "required": true,
                "uuid": ""
            },
            {
                "aliases": [
                    "sub_issue_type"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "issue_type",
                "name": "子任务类型",
                "options": [],
                "required": true,
                "uuid": ""
            },
            {
                "aliases": [
                    "alternative_issue_type"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "issue_type",
                "name": "特殊任务类型（含父子任务类型）",
                "options": [],
                "required": true,
                "uuid": ""
            },
            {
                "aliases": [
                    "status"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "status",
                "name": "任务状态",
                "options": [],
                "required": false,
                "uuid": null
            },
            {
                "aliases": [
                    "status_category"
                ],
                "allowEmpty": false,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "text",
                "name": "[状态类型]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "description"
                ],
                "allowEmpty": true,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "rich_text",
                "name": "任务描述",
                "options": [],
                "required": true,
                "uuid": "field016"
            },
            {
                "aliases": [
                    "description_text",
                    "announcement"
                ],
                "allowEmpty": true,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "text",
                "name": "任务描述文本",
                "options": [],
                "required": true,
                "uuid": ""
            },
            {
                "aliases": [
                    "create_time"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "time",
                "name": "创建时间",
                "options": [],
                "required": false,
                "uuid": "field009"
            },
            {
                "aliases": [
                    "priority"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "option",
                "name": "优先级",
                "options": [
                    {
                        "uuid": "SyvYXmkE",
                        "value": "最高"
                    },
                    {
                        "uuid": "EWrf118M",
                        "value": "较高"
                    },
                    {
                        "uuid": "EDjvSU7Q",
                        "value": "普通"
                    },
                    {
                        "uuid": "BYeB8nCn",
                        "value": "较低"
                    },
                    {
                        "uuid": "MgCLeFSs",
                        "value": "最低"
                    },
                    {
                        "uuid": "Q97Rzb7N",
                        "value": "xx"
                    }
                ],
                "required": true,
                "uuid": "field012"
            },
            {
                "aliases": [
                    "deadline"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "date",
                "name": "截止日期",
                "options": [],
                "required": true,
                "uuid": "field013"
            },
            {
                "aliases": [
                    "number"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "integer",
                "name": "任务号",
                "options": [],
                "required": true,
                "uuid": "field015"
            },
            {
                "aliases": [
                    "parent"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": "",
                "fieldType": "task",
                "name": "父任务",
                "options": [],
                "required": true,
                "uuid": ""
            },
            {
                "aliases": [
                    "path"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": "",
                "fieldType": "text",
                "name": "层级路径",
                "options": [],
                "required": true,
                "uuid": ""
            },
            {
                "aliases": [
                    "watchers"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "user",
                "name": "[任务关注者]",
                "options": [],
                "required": true,
                "uuid": "field008"
            },
            {
                "aliases": [
                    "estimated_hours"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "integer",
                "name": "预估工时",
                "options": [],
                "required": true,
                "uuid": "field018"
            },
            {
                "aliases": [
                    "remaining_manhour"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "integer",
                "name": "剩余工时",
                "options": [],
                "required": true,
                "uuid": "field020"
            },
            {
                "aliases": [
                    "total_manhour"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": 0,
                "fieldType": "integer",
                "name": "[已登记工时]",
                "options": [],
                "required": false,
                "uuid": "field019"
            },
            {
                "aliases": [
                    "estimate_variance"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": 0,
                "fieldType": "integer",
                "name": "[预估偏差]",
                "options": [],
                "required": false,
                "uuid": "field025"
            },
            {
                "aliases": [
                    "time_progress"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": 0,
                "fieldType": "integer",
                "name": "[工时进度]",
                "options": [],
                "required": false,
                "uuid": "field026"
            },
            {
                "aliases": [
                    "server_update_stamp"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "time",
                "name": "更新时间",
                "options": [],
                "required": false,
                "uuid": "field010"
            },
            {
                "aliases": [
                    "sub_tasks"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "task",
                "name": "[子任务]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "related_count"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "integer",
                "name": "关联任务数量",
                "options": [],
                "required": false,
                "uuid": "field023"
            },
            {
                "aliases": [
                    "related_tasks"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "task",
                "name": "[关联任务]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "gantt_chart_uuid"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "text",
                "name": "[甘特图UUID]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "plan_start_date",
                    "plan_start_time"
                ],
                "allowEmpty": true,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "date",
                "name": "[任务计划开始日期]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "plan_end_date",
                    "plan_end_time"
                ],
                "allowEmpty": true,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "date",
                "name": "[任务计划结束日期]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "products"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "product",
                "name": "[所属产品]",
                "options": [],
                "required": false,
                "uuid": "field029"
            },
            {
                "aliases": [
                    "product_modules"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "product_module",
                "name": "[所属功能模块]",
                "options": [],
                "required": false,
                "uuid": "field030"
            },
            {
                "aliases": [
                    "progress"
                ],
                "allowEmpty": true,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "integer",
                "name": "[工作项进度]",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "schedule_group_type"
                ],
                "allowEmpty": true,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "enum",
                "name": "项目计划分组类型",
                "options": [],
                "required": false,
                "uuid": ""
            },
            {
                "aliases": [
                    "field012"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": "EDjvSU7Q",
                "fieldType": "option",
                "name": "优先级",
                "options": [
                    {
                        "uuid": "SyvYXmkE",
                        "value": "最高"
                    },
                    {
                        "uuid": "EWrf118M",
                        "value": "较高"
                    },
                    {
                        "uuid": "EDjvSU7Q",
                        "value": "普通"
                    },
                    {
                        "uuid": "BYeB8nCn",
                        "value": "较低"
                    },
                    {
                        "uuid": "MgCLeFSs",
                        "value": "最低"
                    },
                    {
                        "uuid": "Q97Rzb7N",
                        "value": "xx"
                    }
                ],
                "required": false,
                "uuid": "field012"
            },
            {
                "aliases": [
                    "field027"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "date",
                "name": "计划开始日期",
                "options": [],
                "required": false,
                "uuid": "field027"
            },
            {
                "aliases": [
                    "field028"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "date",
                "name": "计划完成日期",
                "options": [],
                "required": false,
                "uuid": "field028"
            },
            {
                "aliases": [
                    "field033"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "float",
                "name": "进度",
                "options": [],
                "required": false,
                "uuid": "field033"
            },
            {
                "aliases": [
                    "status"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "status",
                "name": "任务状态",
                "options": [],
                "required": false,
                "uuid": null
            },
            {
                "aliases": [
                    "manhours"
                ],
                "allowEmpty": null,
                "builtIn": true,
                "defaultValue": null,
                "fieldType": "manhour",
                "name": "[任务登记工时]",
                "options": [],
                "required": true,
                "uuid": null
            },
            {
                "aliases": [
                    "links"
                ],
                "allowEmpty": false,
                "builtIn": true,
                "defaultValue": false,
                "fieldType": "task_link",
                "name": "[任务关联关系]",
                "options": [],
                "required": false,
                "uuid": null
            }
        ]
    }
}
`

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"unicode"

	"github.com/tidwall/gjson"
)

const (
	defaultConfigPath = "/Users/niuqiang/go_home/src/github.com/bangwork/mygo/2020/0721_field_table/data.json"
	uuid              = "uuid"
	name              = "name"
	typ               = "type"
	desc              = "desc"
)

var (
	header = []string{"UUID", "名称", "类型", "说明"}
	values = []string{"uuid", "name", "type", "desc"}
)

func main() {
	fmt.Println("--------------------")
	data, err := ioutil.ReadFile(defaultConfigPath)
	if err != nil {
		panic(err)
	}
	dataMap := render(data)
	dataMap = calculateLineMaxLine(dataMap)
	print(dataMap)
}

func render(data []byte) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	result = append(result, map[string]interface{}{uuid: 0, name: 0, typ: 0, desc: 0})

	appendFunc := func(key, value gjson.Result) bool {
		fmt.Printf("key=%v, value=%v", key.String(), value.String())
		m := make(map[string]interface{}, 4)
		m[uuid] = gjson.Get(value.Raw, uuid).String()
		m[name] = gjson.Get(value.Raw, name).String()
		m[typ] = typeMap[int(gjson.Get(value.Raw, typ).Int())]
		m[desc] = gjson.Get(value.Raw, desc).String()
		fmt.Printf("1: %#v\n", m)
		result = append(result, m)
		return true
	}

	gjson.Get(string(data), "fields.#(fixed=true)#").ForEach(appendFunc)

	return result
}

func print(items []map[string]interface{}) {
	lenmap := items[0]
	formart := calculateFormat()
	// header
	fmt.Println(fmt.Sprintf(formart, printHeader(lenmap)[0], printHeader(lenmap)[1], printHeader(lenmap)[2], printHeader(lenmap)[3]))
	// align
	fmt.Println(fmt.Sprintf(formart, printAlign(lenmap)[0], printAlign(lenmap)[1], printAlign(lenmap)[2], printAlign(lenmap)[3]))
	for _, valueMap := range items[1:] {
		// content
		fmt.Println(fmt.Sprintf(formart, printString(valueMap[uuid], lenmap[uuid].(int)), printString(valueMap[name], lenmap[name].(int)), printString(valueMap[typ], lenmap[typ].(int)), printString(valueMap[desc], lenmap[desc].(int))))
	}

}

func printHeader(m map[string]interface{}) []string {
	args := make([]string, len(header))
	for i, v := range header {
		args[i] = printString(v, m[values[i]].(int))
	}
	return args
}

func printAlign(m map[string]interface{}) []string {
	result := func(length int) string {
		placeholders := ""
		for i := 1; i < length-1; i++ {
			placeholders += "-"
		}
		return ":" + placeholders
	}
	results := []string{}
	for _, v := range m {
		results = append(results, result(v.(int)))
	}
	return results
}

func printString(v interface{}, length int) string {
	result := ""
	switch v.(type) {
	case string:
		result = v.(string)
	case *string:
		result = *v.(*string)
	case int:
		result = strconv.Itoa(v.(int))
	case int64:
		result = strconv.FormatInt(v.(int64), 10)
	case bool:
		if v.(bool) {
			result = "T"
		} else {
			result = "F"
		}
	default:
		result = fmt.Sprintf("%v", v)
	}

	placeholders := ""
	for i := interfaceLength(result); i < length; i++ {
		placeholders += " "

	}
	return result + placeholders
}

func calculateFormat() string {
	result, placeholder, length := "|", "", len(header)
	for ; length != 0; length-- {
		placeholder += "%s|"
	}
	return result + placeholder
}

func calculateLineMaxLine(items []map[string]interface{}) []map[string]interface{} {
	if len(items) == 0 {
		return items
	}
	lenMap := items[0]
	for _, valueMap := range items[1:] {
		for k, v := range valueMap {
			lenv := interfaceLength(v)
			if lenv > lenMap[k].(int) {
				lenMap[k] = lenv
			}
		}
	}
	items[0] = lenMap
	return items
}

func interfaceLength(v interface{}) int {
	switch v.(type) {
	case string:
		return stringLength(v.(string))
	case *string:
		return stringLength(*v.(*string))
	case int:
		return intLength(v.(int))
	case int64:
		return intLength(int(v.(int64)))
	case bool:
		return 4
	default:
		return 8
	}
}

func intLength(num int) int {
	length := 0
	for num == 0 {
		num = num / 10
		length++
	}
	return length
}

func stringLength(source string) int {
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

const (
	FieldTypeAny            = "any"
	FieldTypeOption         = "option"
	FieldTypeText           = "text"
	FieldTypeInteger        = "integer"
	FieldTypeFloat          = "float"
	FieldTypeDate           = "date"
	FieldTypeTime           = "time"
	FieldTypeUser           = "user"
	FieldTypeStatus         = "status"
	FieldTypeUserList       = "user_list"
	FieldTypeNumber         = "number"
	FieldTypeMultiLineText  = "multi_line_text"
	FieldTypeMultiOption    = "multi_option"
	FieldTypeStatusCategory = "status_category"
	FieldTypeBoolean        = "boolean"
	FieldTypeContext        = "context"
	FieldTypeRichText       = "rich_text"
	FieldTypeOptions        = "options"
	FieldTypeStatuses       = "statuses"
	FieldTypeEnum           = "enum"
	FieldTypeAliases        = "aliases"
	FieldTypeBucket         = "bucket"
	FieldTypeTimeSeries     = "time_series"
	FieldTypeUserDomains    = "user_domains"
	FieldTypeDefaultValue   = "default_value"
	FieldTypePercentage     = "percentage"
)

var typeMap = map[int]string{
	1:  FieldTypeOption,
	2:  FieldTypeText,
	3:  FieldTypeInteger,
	4:  FieldTypeFloat,
	5:  FieldTypeDate,
	6:  FieldTypeTime,
	8:  FieldTypeUser,
	13: FieldTypeUserList,
	14: FieldTypeNumber,
	15: FieldTypeMultiLineText,
	16: FieldTypeMultiOption,
	17: FieldTypeStatusCategory,
	18: FieldTypeContext,
	19: FieldTypeBoolean,
	20: FieldTypeRichText,
	21: FieldTypeAny,
	22: FieldTypeStatus,
	23: FieldTypeOptions,
	24: FieldTypeStatuses,
	25: FieldTypeEnum,
	26: FieldTypeAliases,
	27: FieldTypeBucket,
	29: FieldTypeUserDomains,
	34: FieldTypeDefaultValue,
	48: FieldTypePercentage,
}

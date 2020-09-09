package utils

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
)

const (
	sqlDefaultSep       = " ,"
	sqlPlaceholderLimit = 1024
)

func GetPlacholderCount(columns string) int {
	return strings.Count(columns, ",") + 1
}

// 传入列，根据key生产 key.列
// columns e.g. "uuid, project_uuid, status_uuid"
// key e.g. "f1"
// return "f1.uuid, f1.project_uuid, f1.status_uuid"
//
func SqlMultiColumnsWithTable(table string, columns string) string {
	columnsRaw := strings.Split(columns, ",")
	length := len(columnsRaw)
	if length <= 0 || len(table) == 0 {
		return columns
	}

	for i := 0; i < length; i++ {
		column := strings.TrimSpace(columnsRaw[i])
		columnsRaw[i] = fmt.Sprintf("%s.%s", table, column)
	}

	return strings.Join(columnsRaw, ",")
}

// SqlPlacholdsFromString 根据传入的字符串，生成 placholder。
// columns: e.g. "uuid, project_uuid, status_uuid"
func SqlPlacholdsFromString(columns string) string {
	return SqlPlaceholds(GetPlacholderCount(columns))
}

// 创建 sql in 查询条件
func SqlPlaceholds(count int) string {
	return appendDuplicateString("?", sqlDefaultSep, count)
}

func SqlLikeOrCondition(column string, length int) string {
	where := fmt.Sprintf("%s LIKE ?", column)
	return appendDuplicateString(where, " OR ", length)
}

// sql 多字段查询／插入条件
func SqlMultiColumnPlaceholders(columnCount int, count int) string {
	if columnCount <= 0 || count <= 0 {
		return ""
	}
	one := fmt.Sprintf("(%s)", appendDuplicateString("?", sqlDefaultSep, columnCount))
	return appendDuplicateString(one, sqlDefaultSep, count)
}

// 创建Sql 多个or 查询条件
func SqlMultiCondition(singleCondition string, count int) string {
	return appendDuplicateString(singleCondition, " or ", count)
}

// 创建 sql 批量插入条件
func SqlInMultiInsertValues(columnCount int, count int) string {
	return SqlMultiColumnPlaceholders(columnCount, count)
}

func appendDuplicateString(character, separator string, count int) string {
	if count <= 0 {
		return ""
	}
	var b bytes.Buffer
	for i := 0; i < count; i++ {
		if i > 0 {
			b.WriteString(separator)
		}
		b.WriteString(character)
	}
	return b.String()
}
func BuildSqlArgs(args ...interface{}) ([]interface{}, error) {
	newArgs := make([]interface{}, 0)
	addEleFun := func(ele interface{}) {
		newArgs = append(newArgs, ele)
		return
	}
	for _, arg := range args {
		switch v := arg.(type) {
		case string, int, int32, int64, bool, *string, *int, *int32, *int64:
			addEleFun(v)
		case []string:
			for _, e := range v {
				addEleFun(e)
			}
		case []int:
			for _, e := range v {
				addEleFun(e)
			}
		case []int32:
			for _, e := range v {
				addEleFun(e)
			}
		case []int64:
			for _, e := range v {
				addEleFun(e)
			}
		case []*string:
			for _, e := range v {
				addEleFun(e)
			}
		default:
			return nil, errors.TypeMismatchError(arg,
				"string",
				"int",
				"int32",
				"int64",
				"bool",
				"*string",
				"*int",
				"*int32",
				"*int64",
				"[]string",
				"[]int",
				"[]int32",
				"[]int64",
				"[]string")
		}
	}
	return newArgs, nil
}

package main

import (
	"bufio"
	"fmt"
	"github.com/bangwork/mygo/20191115/config"
	"github.com/bangwork/mygo/20191115/models"
	"unicode"

	"gopkg.in/gorp.v1"
	"log"
	"os"
	"strings"
)

type ColumnInfo struct {
	ColumnName string ` db:"column_name"`
	ColumeType string ` db:"column_type"`
	Default    string `db:"column_default"`
	IsNullAble string `db:"is_nullable"`
	Comment    string `db:"column_comment"`
	Key        string `db:"column_key"`
}

func main() {
	if err := config.LoadConfigs(""); err != nil {
		log.Fatalf("init config err:% v", err)
	}

	if err := models.InitDB(); err != nil {
		log.Fatalf("init db err:% v", err)
	}

	printDataBase()
}

func printDataBase() {
	var lineList []string
	lineList = append(lineList, "# ONES 数据字典\n")
	lineList = append(lineList, "\n")
	tableList, err := GetTables(models.ProjectDBMap)
	if err != nil {
		panic(err)
	}
	for _, tableName := range tableList {
		lineList = append(lineList, printTable(models.ProjectDBMap, "project", tableName)...)
	}

	tableList, err = GetTables(models.WikiDBMap)
	if err != nil {
		panic(err)
	}
	for _, tableName := range tableList {
		lineList = append(lineList, printTable(models.WikiDBMap, "wiki", tableName)...)
	}

	writeFile("ONES_data_dict.md", lineList)
}

func GetTables(src gorp.SqlExecutor) ([]string, error) {
	tables := make([]string, 0)
	_, err := src.Select(&tables, "show tables")
	return tables, err
}

func GetTableColumns(src gorp.SqlExecutor, tableName string) ([]ColumnInfo, error) {
	var columns []ColumnInfo
	sql := fmt.Sprintf("select column_name, column_type, IFNULL(column_default,'无') as column_default, IFNULL(is_nullable,'无') as is_nullable, column_comment, column_key from information_schema.columns where table_name='%s'", tableName)
	_, err := src.Select(&columns, sql)
	return columns, err
}

func printTable(src gorp.SqlExecutor, tableSchema, tableName string) []string {
	var table []string
	columns, err := GetTableColumns(src, tableName)
	if err != nil {
		panic(err)
	}

	tableHead := fmt.Sprintf("## `%s`", tableName)
	tableComment := GetTableComment(src, tableSchema, tableName)
	if tableComment != "" {
		tableComment = " - " + tableComment
	}

	table = append(table, tableHead, tableComment, "\n")
	table = append(table, "\n")
	table = append(table, "| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |\n")
	table = append(table, "| :--- | :--- | :--- | :--- | :--- |\n")

	for _, column := range columns {
		// 授权相关不提供注释
		if tableName == "team" || tableName == "organization" || tableName == "product_authorization" {
			column.Comment = ""
		}
		table = append(table, formatMarkDownOneLine(column))
	}
	table = append(table, "\n")
	return table
}

func GetTableComment(src gorp.SqlExecutor, tableSchema, tableName string) string {
	sql := fmt.Sprintf("select table_comment from information_schema.tables where table_schema='%s' and table_name='%s';", tableSchema, tableName)
	comment, err := src.SelectStr(sql)
	if err != nil {
		panic(err)
	}
	return comment
}
func formatMarkDownOneLine(info ColumnInfo) string {
	info.Comment = appendBR(strings.TrimSpace(info.Comment))
	info.ColumeType = appendBR(info.ColumeType)
	info.Default = strings.TrimSpace(info.Default)
	if info.Default == "" {
		info.Default = "空串"
	}
	if strings.TrimSpace(info.IsNullAble) == "NO" {
		info.IsNullAble = "no"
	}
	if strings.TrimSpace(info.IsNullAble) == "YES" {
		info.IsNullAble = "yes"
	}
	return fmt.Sprintf("| %s | %s | %s | %s | %s |\n", fmt.Sprintf("`%s`", info.ColumnName), strings.TrimSpace(info.ColumeType), info.Default, info.IsNullAble, info.Comment)
}

func appendBR(str string) string {
	if len(str) > 25 {
		// return `<span style="word-break: break-all;white-space: normal;">` + str + `</span>`
		newstr := ""
		i := 0
		for _, c := range str {
			if unicode.Is(unicode.Han, c) {
				i = i + 2
			} else {
				i = i + 1
			}
			if i == 40 || i == 41 {
				i = 0
				newstr += string(c) + "<br>"
				continue
			}
			newstr += string(c)
		}
		return newstr
	}
	return str
}
func writeFile(fileName string, lines []string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	for _, item := range lines {
		_, err := w.WriteString(item)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
}

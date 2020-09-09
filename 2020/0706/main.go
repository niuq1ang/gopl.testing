package main

import (
	"bytes"
	"fmt"

	"github.com/tidwall/gjson"
)

var str = `
{
    "data": {
        "tasks": [
            {
                "_NnAcdq9R": "US1xmvXf",
                "_PUuiRRim": {
                    "uuid": "N3xRCJzH",
                    "value": "提供安装包"
                },
                "_TioFkeZn": {
                    "name": "林灏华",
                    "uuid": "5DFNahq5"
                },
                "_U1Zf7epq": null,
                "name": "小米私有部署初装申请",
                "subTasks": [
                    {
                        "name": "【检查】确保开启CAS单点登录",
                        "uuid": "8uJU1VNkSDnZxKFj"
					},
					{
                        "name": "【检查cc】确保开启CAS单点登录",
                        "uuid": "8uJU1VNkSDnZxKFj"
                    }
                ],
                "uuid": "8uJU1VNkVWY7eERs"
            }
        ]
    }
}
`

func main() {
	TestGJosn()
}

func TestGJosn() {
	results := []string{}
	appendFunc := func(key, value gjson.Result) bool {
		fmt.Println("key: ", key.String())
		fmt.Println("value: ", value.String())
		results = append(results, value.String())
		return true
	}
	fmt.Println(gjson.Get(str, "data.tasks.0.subTasks.#.name").String())
	gjson.Get(str, "data.tasks.0.subTasks.#.name").ForEach(appendFunc)
	fmt.Println(results)
}
func SliceToGraphqlArray(array []string) string {
	buf := &bytes.Buffer{}
	for _, arr := range array {
		buf.WriteByte('"')
		buf.WriteString(arr)
		buf.WriteByte('"')
		buf.WriteByte(',')
		buf.WriteByte(' ')
	}
	return buf.String()
}

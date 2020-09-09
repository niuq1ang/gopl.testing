package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	var releatedTaskUUID string
	pickRelatedTaskUUIDFunc := func(key, value gjson.Result) bool {
		fmt.Println(gjson.Get(value.String(), "taskLinkTypeUUID").String())
		fmt.Println(gjson.Get(value.String(), "taskUUID").String())

		if gjson.Get(value.String(), "taskLinkTypeUUID").String() == "UoSpjAPC" {
			releatedTaskUUID = gjson.Get(value.String(), "taskUUID").String()
			return false
		}
		return true
	}
	gjson.Get(josnData, "data.tasks.0.links").ForEach(pickRelatedTaskUUIDFunc)
	fmt.Println("---", releatedTaskUUID)
}

var josnData = `{
	"data": {
		"tasks": [
			{
				"links": [
					{
						"taskLinkTypeUUID": "Nom5Tqcq",
						"taskUUID": "5ZPerY1K4iGscPrC"
					},
					{
						"taskLinkTypeUUID": "Nom5Tqcq",
						"taskUUID": "5ZPerY1KC50oX8Cu"
					},
					{
						"taskLinkTypeUUID": "UoSpjAPC",
						"taskUUID": "5ZPerY1KYmXpHGvc"
					},
					{
						"taskLinkTypeUUID": "Nom5Tqcq",
						"taskUUID": "5ZPerY1Kk3IBNAz7"
					},
					{
						"taskLinkTypeUUID": "Nom5Tqcq",
						"taskUUID": "5ZPerY1KqZ7SepXH"
					},
					{
						"taskLinkTypeUUID": "Nom5Tqcq",
						"taskUUID": "5ZPerY1Ks2uAeh0z"
					},
					{
						"taskLinkTypeUUID": "Nom5Tqcq",
						"taskUUID": "5ZPerY1KsOg4aGov"
					}
				]
			}
		]
	}
}`

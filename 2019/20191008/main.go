package main

import (
	"encoding/json"
	"fmt"
)

var str = `{"user_name":"txx"}`
var str2 = `{"user_name":"xtt"}`

type User struct {
	UserName string `json:"user_name"`
}

func main() {
	req := new(User)
	// req := User{}
	results := funcTypeTest(req)

	for _, result := range results {
		user, ok := result.(*User)
		fmt.Println("----", ok)
		if ok {
			fmt.Println("----", user.UserName)
		}
		fmt.Println(result)
	}

}

func funcTypeTest(req interface{}) []interface{} {
	var results []interface{}
	err := json.Unmarshal([]byte(str), &req)
	if err != nil {
		fmt.Println(err)
	}
	results = append(results, req)

	err = json.Unmarshal([]byte(str2), &req)
	if err != nil {
		fmt.Println(err)
	}
	results = append(results, req)
	return results
}

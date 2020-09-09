package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type BaseRes struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	RequestBody  []byte `json:"-"`
	Identifer    string `json:"-"`
}

func main() {
	res := BaseRes{
		ErrorCode:    100,
		ErrorMessage: "funny test",
	}

	resBytes, err := json.Marshal(res)
	var baseRes BaseRes
	baseRes.Identifer = "Identifer,123456"
	baseRes.RequestBody = resBytes
	err = json.Unmarshal(resBytes, &baseRes)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(baseRes)
}

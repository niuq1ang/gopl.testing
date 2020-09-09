package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int64
}

func main() {

	m := map[string]string{
		"1": "1",
		"a": "a",
	}
	m["a"] = "b"
	fmt.Println(m)

	stuMap := map[string]Student{
		"XiaoMing": Student{
			"Xiaoming",
			18,
		},
	}

	stuMap["XiaoMing"] = Student{
		"wanglufei",
		17,
	}

	fmt.Println()
	fmt.Println()

	mapMap := map[string]map[string]string{
		"A": map[string]string{
			"a": "aa",
		},
	}
	mapMap["A"]["a"] = "bb"
	fmt.Println(mapMap)

	fmt.Println()
	mapInterface := make(map[string]interface{})
	s := "ss"
	mapInterface["A"] = s
	mapInterface["A"] = "b"
	mapInterface["B"] = 3
	mapInterface["C"] = stuMap
	fmt.Println(mapInterface, "\n")

	params := make([]map[string]interface{}, 0)
	tempMap := make(map[string]interface{})
	for k, v := range mapInterface {
		tempMap["value"] = v
		tempMap["key"] = k
		tempMap["common"] = "common"
		fmt.Println(tempMap)

		params = append(params, tempMap)
		fmt.Println(params)

	}
	fmt.Println()
	fmt.Println(params)
}

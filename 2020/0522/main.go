package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	u := GetUser()
	fmt.Println("111111111")
	PrintUser(u.Name)
	fmt.Println("22222222")

}

func IntArrayDifference(old []int, new []int) (additions []int, deletions []int) {
	additionsMap := make(map[int]struct{})
	deletionsMap := make(map[int]struct{})
	for _, s := range old {
		deletionsMap[s] = struct{}{}
	}
	for _, s := range new {
		additionsMap[s] = struct{}{}
	}
	for s, _ := range additionsMap {
		if _, ok := deletionsMap[s]; !ok {
			additions = append(additions, s)
		}
	}
	for s, _ := range deletionsMap {
		if _, ok := additionsMap[s]; !ok {
			deletions = append(deletions, s)
		}
	}
	return
}

func GetUser() *User {
	return nil
}

func PrintUser(i string) {
	fmt.Println("3333333333")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(i)
}

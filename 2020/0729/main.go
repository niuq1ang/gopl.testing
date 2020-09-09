package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []*User{
		{"xiaoming", 18},
		{"xiaohua", 20},
	}
	updateUser(users...)
	fmt.Println(users[0].Age, users[0].Age)
}

func updateUser(users ...*User) {
	for _, u := range users {
		u.Age = 100
	}
}

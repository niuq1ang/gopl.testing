package main

import (
	"fmt"
)

type User struct {
	age, name string
}

func main() {

	var user1 []*User
	user2 := make([]*User, 0)

	fmt.Println(user1, len(user1), cap(user1), user2, len(user2), cap(user2))

	fmt.Println(int(^uint(0) >> 1))
	s := []int{0, 1, 2}
	s2 := []int{4, 5, 6, 7}

	s1 := make([]int, 0, len(s)+len(s2))
	s1 = s1[:len(s)]
	fmt.Println("1:", s1, len(s1))
	copy(s1, s)
	fmt.Println(s1, len(s1))

	n := copy(s1[len(s):], s2) // n == 2, s == []int{1, 2, 2}
	fmt.Println(n, s1)

	fmt.Println([]byte("\n"))
}

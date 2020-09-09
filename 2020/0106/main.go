package main

import (
	"fmt"
	"strings"
)

const (
	i = 1 << iota
	j = 3 << iota
	k = 3 << iota
	l = 3 << iota
)

func main() {

	fmt.Println(i, j, k, l)

}

var method uint8

func calculateMetheds(methods ...string) uint8 {
	var sum uint8
	for _, m := range methods {
		sum += calculateMethed(m)
	}
	return sum
}

func calculateMethed(m string) uint8 {
	switch strings.ToLower(m) {
	case "get":
		return 1 << 0
	case "post":
		return 1 << 1
	case "put":
		return 1 << 2
	case "delete":
		return 1 << 3
	}
	return 0
}

func validMethod(n, m uint8) bool {
	return n&m > 0
}

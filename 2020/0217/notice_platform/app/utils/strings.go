package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	emptyString = ""
)

func IsSpaceString(s string) bool { return strings.TrimSpace(s) == emptyString }

func IsEmptyString(s string) bool { return s == emptyString }

func IsBlankString(s *string) bool { return s == nil || len(*s) == 0 }

func NonEmptyString(s string) bool { return !IsEmptyString(s) }

const code = "0123456789ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvxwyz-*"
const upperCaseOffset = 62
const lowerCaseOffset = 36
const numberOffset = 10

func randomString(size int, max int) string {
	rand.Seed(time.Now().UnixNano())
	buffer := make([]byte, size, size)
	for i := 0; i < size; i++ {
		buffer[i] = code[rand.Intn(max)]
	}
	return string(buffer[:size])
}

//含小写英数
func Random36String(size int) string {
	return randomString(size, lowerCaseOffset)
}

//含大小写英数
func Random62String(size int) string {
	return randomString(size, upperCaseOffset)
}

func RandomNumberString(size int) string {
	return randomString(size, numberOffset)
}

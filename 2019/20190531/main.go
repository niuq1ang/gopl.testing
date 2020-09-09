package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ipToStringName(""))
}
func ipToStringName(ip string) string {
	var b bytes.Buffer
	parts := strings.Split(ip, ".")
	for _, val := range parts {
		for len(val) < 3 {
			val = fmt.Sprintf("%s%s", "0", val)
		}
		b.WriteString(val)
	}
	return b.String()
}

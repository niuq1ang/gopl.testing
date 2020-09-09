package main

import (
	"fmt"
	"strings"
)

func main() {
	uri := "/organization/{organizationUUID}/create_team"
	parts := strings.Split(uri, "/")
	fmt.Println(len(parts))
	fmt.Printf("%v", parts)
}

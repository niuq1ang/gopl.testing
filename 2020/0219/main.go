package main

import (
	"fmt"
	"time"
)

func main() {

LOOP:
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("aa")
			break LOOP
			fmt.Println("bb")
		}

	}

	//goto直接调到LAbEL2
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LAbEL2
			}
		}
	}
	fmt.Println("PreLAbEL2")
LAbEL2:
	fmt.Println("LastLAbEL2")

	//break跳出和LAbEL1同一级别的循环,继续执行其他的
LAbEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LAbEL1
			}
		}
	}
	fmt.Println("OK")

	//continue
LABEL3:

	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			continue LABEL3
		}
	}
	fmt.Println("ok")
}

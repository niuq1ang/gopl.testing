package main

import (
	"fmt"
)

type TestStruct struct {
	Number   int
	Boolen   bool
	String   string
	Byte     byte
	IntArray []int
	Pointer  uintptr
}

func (self TestStruct) valueTest() {
	fmt.Printf("Value: %p\n", &self)
}

func (self *TestStruct) pointerTest() {
	fmt.Printf("Pointer: %p\n", self)
}

func main() {
	d := TestStruct{}
	p := &d

	fmt.Printf("Data: %p\n", p)

	d.valueTest()
	d.pointerTest()

	p.valueTest()
	p.pointerTest()
}


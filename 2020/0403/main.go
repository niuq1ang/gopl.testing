package main

import (
	"plugin"
)

func main() {
	p, err := plugin.Open("plugin/plugin.so")
	if err != nil {
		panic(err)
	}
	// A Symbol is a pointer to a variable or function.
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	// *v.(*int) = 7
	pointerV := v.(*int)
	*pointerV = 7
	*v.(*int) = 7
	*(v.(*int)) = 7

	f.(func())() // prints "Hello, number 7"
}

type p int

func (p) f() {

}

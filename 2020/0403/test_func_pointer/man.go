package main

import "fmt"

type p int

type FT func(int)

func Fa(int) {
	fmt.Println("------")
}
func Test(FT) {
	fmt.Println("Test")

}

func main() {
	Test(Fa) //pass function as parameter
}

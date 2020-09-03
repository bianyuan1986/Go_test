package main

import (
	"fmt"
)

type A struct {
	name string
}

type B struct {
	member A
}

func (A) Amethod() {
	fmt.Println("A method!")
}

func main() {
	var a = A{name: "A name"}
	B{a}.member.Amethod()
}

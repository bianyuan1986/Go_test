package main

import (
	"fmt"
)

type TestData struct {
	name string
}

func (d *TestData) PrintName() {
	fmt.Println(d.name)
}

func SliceAnalysis() {

}

func InterfaceAnalysis() {

}

func StructAnalysis() {

}

func StringAnalysis() {

}

func ChannelAnalysis() {

}

func ArrayAnalysis() {

}

func MapAnalysis() {

}

func ExternalInterface(v interface{}) {
	fmt.Println("Get struct!")
	d := v.(TestData)
	a := &d
	a.PrintName()
	fmt.Println("Get struct pointer!")
	v.(*TestData).PrintName()
}

func ExternalInterfacePointer(v *interface{}) {

}

func main() {
	fmt.Println("Go struct analysis!")
	a := TestData{name: "2020 Go!"}
	ExternalInterface(a)
	ExternalInterface(&a)
}

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

func ByteToStringAnalysis() {
	var data []byte
	data = make([]byte, 5, 7)
	for i := 0; i < 5; i++ {
		data[i] = byte(101 + i)
	}
	str := string(data[:])
	fmt.Println("Str:", str)
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
	/*fmt.Println("Go struct analysis!")
	a := TestData{name: "2020 Go!"}
	ExternalInterface(a)
	ExternalInterface(&a)*/
	ByteToStringAnalysis()
}

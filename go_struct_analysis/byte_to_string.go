package main

import (
	"fmt"
)

type tiger1 struct {
	name string
}

type tiger2 struct {
	count int
}

func changeString(v interface{}) {
	d := v.(tiger1)
	d.name = "20201!"
}

func changeStringPointer(v interface{}) {
	d := v.(*tiger1)
	d.name = "20201!"
}

func changeInt(v interface{}) {
	d := v.(tiger2)
	d.count = 2021
}

func changeIntPointer(v interface{}) {
	d := v.(*tiger2)
	d.count = 2021
}

func main() {
	var b []byte = make([]byte, 0, 0)
	var b1 []byte = make([]byte, 4, 4)
	s := string(b[:])
	fmt.Println(s)
	s = string(b1[:])
	fmt.Println(s)
	b1[0] = 'a'
	b1[1] = 'b'
	s = string(b1[:])
	fmt.Println(s)

	t1 := tiger1{name: "2020 Go!"}
	changeString(t1)
	fmt.Printf("%s\n", t1.name)
	changeStringPointer(&t1)
	fmt.Printf("%s\n", t1.name)
	t2 := tiger2{count: 2020}
	changeInt(t2)
	fmt.Printf("%d\n", t2.count)
	changeIntPointer(&t2)
	fmt.Printf("%d\n", t2.count)
}

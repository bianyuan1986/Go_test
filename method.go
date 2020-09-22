package main

import (
	"fmt"
)

type Tiger struct {
	name string
}

func (t *Tiger) GetName() {
	fmt.Println(t.name)
}

/*保存某个对象实例的方法，记录方法地址的同时关联了对象*/
func main() {
	data := make(map[string]func())
	a := &Tiger{name: "Tiger1"}
	b := &Tiger{name: "Tiger2"}
	data["t1"] = a.GetName
	data["t2"] = b.GetName

	for _, v := range data {
		v()
	}
}

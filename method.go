package main

import (
	"fmt"
)

type Tiger struct {
	name string
}

type Tiger1 struct {
	name string
}

type I interface {
	Dump()
}

func (t *Tiger) GetName() {
	fmt.Println(t.name)
}

/*
func (t *Tiger) Dump() {
	fmt.Println(t.name)
}
*/

func (t Tiger) Dump() {
	fmt.Println(t.name)
}

func checkInterfaceImplementI(object interface{}) {
	if _, ok := object.(I); !ok {
		fmt.Println("Type doesn't implement interface I")
		return
	}
	fmt.Println("Type implement interface I")
}

func _checkInterfaceImplementI(object *interface{}) {
	if _, ok := (*object).(I); !ok {
		fmt.Println("Type doesn't implement interface I")
		return
	}
	fmt.Println("Type implement interface I")
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
	t := Tiger{"2020"}
	t1 := Tiger1{"2021"}
	fmt.Println("Check struct [Tiger]")
	checkInterfaceImplementI(t)
	fmt.Println("Check struct [Tiger1]")
	checkInterfaceImplementI(t1)
	fmt.Println("Check struct [Tiger]")
	checkInterfaceImplementI(&t)
	fmt.Println("Check struct [Tiger1]")
	checkInterfaceImplementI(&t1)
}

/*
接口方法实现的接收者为指针类型，当参数为interface{}类型时如果通过值传递则
判断当前对象未实现该接口，如果传递指针类型则判断实现该接口
func (t *Tiger) Dump() {
	fmt.Println(t.name)
}
Tiger1
Tiger2
Check struct [Tiger]
Type doesn't implement interface I
Check struct [Tiger1]
Type doesn't implement interface I
Check struct [Tiger]
Type implement interface I
Check struct [Tiger1]
Type doesn't implement interface I

接口方法实现的接收者为对象类型，当参数为interface{}类型时如果通过值传递或者
通过指针传递则都判断实现该接口
func (t Tiger) Dump() {
	fmt.Println(t.name)
}
Tiger1
Tiger2
Check struct [Tiger]
Type implement interface I
Check struct [Tiger1]
Type doesn't implement interface I
Check struct [Tiger]
Type implement interface I
Check struct [Tiger1]
Type doesn't implement interface I
*/

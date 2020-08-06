package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Array1 []int

func (arr Array1) Len() int {
	fmt.Printf("Array1 Len():")
	return len(arr)
}

func (arr Array1) Less(i, j int) bool {
	fmt.Printf("Array1 Less():")
	return arr[i] < arr[j]
}

func (arr Array1) Swap(i, j int) {
	fmt.Println("Array1 Swap():")
	arr[i], arr[j] = arr[j], arr[i]
}

type Array2 []int

func (arr Array2) Len() int {
	fmt.Printf("Array2 Len():")
	return len(arr)
}

func (arr Array2) Less(i, j int) bool {
	fmt.Printf("Array2 Less():")
	return arr[i] < arr[j]
}

func (arr Array2) Swap(i, j int) {
	fmt.Println("Array2 Swap():")
	arr[i], arr[j] = arr[j], arr[i]
}

/*内嵌接口类型，不依赖与具体的实现，只要实现了sort.Interface接口的对象都可以*/
type Sortable struct {
	sort.Interface
	Type string
}

/*重写Len方法，任何实现了sort.Interface接口的对象都可以被重写*/
func (s Sortable) Len() int {
	fmt.Println("Override Len()")
	return s.Interface.Len()
}

func NewSortable(i sort.Interface) Sortable {
	t := reflect.TypeOf(i).String()
	return Sortable{Interface: i, Type: t}
}

func DoSomething(s Sortable) {
	fmt.Println(s.Type)
	fmt.Println(s.Len())
}

/*内嵌结构体类型，通过组合来继承方法*/
type Sortable1 struct {
	Array2
	Type string
}

/*重写方法，通过内嵌结构体的方式来重写方法，此时Sortable1类型依赖与具体的实现，
即依赖与Array2的实现，此时只能重写Array2类型的Len方法*/
func (s Sortable1) Len() int {
	fmt.Println("Override Len()")
	return s.Array2.Len()
}

func NewSortable1(arr Array2) Sortable1 {
	t := reflect.TypeOf(arr).String()
	return Sortable1{Array2: arr, Type: t}
}

func NewSortable2(arr *Array2) Sortable1 {
	t := reflect.TypeOf(arr).String()
	return Sortable1{Array2: *arr, Type: t}
}
func DoSomethingNew(s Sortable1) {
	fmt.Println(s.Type)
	fmt.Println(s.Len())
}

func main() {
	arr1 := Array1{1, 2, 3}
	arr2 := Array2{3, 2, 1, 0}
	DoSomething(NewSortable(arr1))
	fmt.Println("-----------------")
	DoSomething(NewSortable(arr2))
	fmt.Println("-----------------")
	DoSomethingNew(NewSortable1(arr2))
	fmt.Println("-----------------")
	s1 := NewSortable2(&arr2)
	fmt.Println(s1.Len())
	fmt.Println("-----------------")
	fmt.Println(s1.Array2.Len())
	fmt.Println("-----------------")
	fmt.Println(arr2.Len())
}

/*
Output:

main.Array1
Override Len()
Array1 Len():3
-----------------
main.Array2
Override Len()
Array2 Len():4
-----------------
main.Array2
Override Len()
Array2 Len():4
-----------------
Override Len()
Array2 Len():4
-----------------
Array2 Len():4
-----------------
Array2 Len():4
*/

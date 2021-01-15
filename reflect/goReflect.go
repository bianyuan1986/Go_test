package main

import (
	"fmt"
	"reflect"
)

type subRule struct {
	Sid   int
	Sname string
}

type rule struct {
	Id     int
	Name   string
	SRules subRule
}

func testReflect() {
	r := rule{Id: 2021, Name: "xss"}
	//r.sRules = append(r.sRules, subRule{sid: 3, sname: "ddd"})
	r.SRules.Sid = 3
	r.SRules.Sname = "ddd"
	parseStructType(r, 0)
}

func printPrefix(level int) {
	for i := 0; i < level*4; i++ {
		fmt.Printf(" ")
	}
}

func parseStructType(data interface{}, level int) interface{} {
	t := reflect.TypeOf(data)
	var v reflect.Value
	if t.Kind() == reflect.Ptr {
		realData := reflect.ValueOf(data).Elem().Interface()
		t = reflect.TypeOf(realData)
		v = reflect.ValueOf(realData)
	} else if t.Kind() == reflect.Struct {
		v = reflect.ValueOf(data)
	} else {
		printPrefix(level)
		fmt.Println("Not Struct Type!")
		return nil
	}
	printPrefix(level)
	fmt.Printf("Type of data:%10s\n", t)
	fieldNum := t.NumField()
	printPrefix(level)
	fmt.Printf("Struct [%s] contain [%d] fields\n", t.Name(), fieldNum)

	newValue := reflect.New(t).Elem()

	printPrefix(level + 1)
	fmt.Printf("------------------------------\n")
	for i := 0; i < fieldNum; i++ {
		printPrefix(level + 1)
		fmt.Printf("[%02d]: [%s]\n", i, t.Field(i).Name)

		subValue := v.Field(i)
		subType := t.Field(i).Type
		printPrefix(level + 1)
		fmt.Printf("SubType:%15s |", subType)
		kind := subValue.Kind()
		/*printPrefix(level)
		fmt.Printf("Kind: [%s]\n", kind.String())*/
		switch kind {
		case reflect.Int:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.String:
			fmt.Println("Value:", subValue)
		case reflect.Struct:
			fmt.Println("Value:", subValue)
			parseStructType(reflect.New(subType).Interface(), level+1)
		default:
			fmt.Println("Unsupported value!")
		}

		switch kind {
		case reflect.Int:
			newValue.Field(i).SetInt(2022)
			//*(subValue.Addr().Interface().(**int)) = new(int)
		case reflect.Int32:
		case reflect.String:
			newValue.Field(i).SetString("sql")
		}
	}
	printPrefix(level + 1)
	fmt.Printf("------------------------------\n")
	printPrefix(level + 1)
	fmt.Println("New Value:", newValue)

	return newValue.Interface()
}

func main() {
	testReflect()
}

/*
Type of data: main.rule
Struct [rule] contain [3] fields
    ------------------------------
    [00]: [Id]
    SubType:            int |Value: 2021
    [01]: [Name]
    SubType:         string |Value: xss
    [02]: [SRules]
    SubType:   main.subRule |Value: {3 ddd}
    Type of data:main.subRule
    Struct [subRule] contain [2] fields
        ------------------------------
        [00]: [Sid]
        SubType:            int |Value: 0
        [01]: [Sname]
        SubType:         string |Value:
        ------------------------------
        New Value: {2022 sql}
    ------------------------------
    New Value: {2022 sql {0 }}

*/

package main

import (
	"./lua"
	"./waf"
	"fmt"
)

func main() {
	w := waf.NewWaf()
	var i interface{} = w
	if handler, ok := i.(lua.LuaHandler); ok {
		fmt.Println("Implement!")
		handler.ExecuteLua("test.lua")
	} else {
		fmt.Println("Not Implement!")
	}
}

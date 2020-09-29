package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"strings"
	"time"
	"unsafe"
)

/*
#cgo LDFLAGS: -L./ -llua
#cgo CFLAGS: -I./ -g

#include "lauxlib.h"
#include "lualib.h"
#include "lua.h"
#include <stdlib.h>

int sum(int range)
{
	int total = 0;
	int i = 0;

	for( ; i < range; i++)
	{
		total += i;
	}

	return total;
}
*/
import "C"

const (
	REPEAT_CNT = 4
)

func sum(boundary int) int {
	var total int = 0
	var i int = 0
	for i = 0; i < boundary; i++ {
		total += i
	}

	return total
}

func golang_test() {
	loopCnt := 1
	for i := 0; i < REPEAT_CNT; i++ {
		loopCnt *= 10
		start := time.Now()
		for j := 0; j < loopCnt; j++ {
			sum(i)
		}
		elapse := time.Now().Sub(start)
		fmt.Println("Golang Time consumed:\n", elapse)
	}
	fmt.Println("\n\n")
}

func cgo_test() {
	loopCnt := 1
	for i := 0; i < REPEAT_CNT; i++ {
		loopCnt *= 10
		start := time.Now()
		for j := 0; j < loopCnt; j++ {
			C.sum(C.int(i))
		}
		elapse := time.Now().Sub(start)
		fmt.Println("Cgo Time consumed:\n", elapse)
	}
	fmt.Println("\n\n")
}

func golang_lua_test() {
	start := time.Now()
	L := lua.NewState()
	if err := L.DoString(`print("Hello World!")`); err != nil {
		panic(err)
	}
	elapse := time.Now().Sub(start)
	fmt.Println("Golang lua Time consumed:\n", elapse)
	fmt.Println("\n\n")

	L.Close()
}

func golang_lua_precompile_test() {
	start := time.Now()
	r := strings.NewReader(`print("Hello World!")`)
	chunk, err := parse.Parse(r, "tiger")
	if err != nil {
		fmt.Println("Parse failed!")
		return
	}
	proto, err := lua.Compile(chunk, "tiger")
	if err != nil {
		fmt.Println("Compile failed!")
		return
	}
	L1 := lua.NewState()
	lfunc := L1.NewFunctionFromProto(proto)
	L1.Push(lfunc)
	L1.PCall(0, lua.MultRet, nil)
	elapse := time.Now().Sub(start)
	fmt.Println("Golang lua with precompile Time consumed:\n", elapse)
	fmt.Println("\n\n")

	L1.Close()
}

func golang_lua_bytecode_run_test() {
	r := strings.NewReader(`print("Hello World!")`)
	chunk, err := parse.Parse(r, "tiger")
	if err != nil {
		fmt.Println("Parse failed!")
		return
	}
	proto, err := lua.Compile(chunk, "tiger")
	if err != nil {
		fmt.Println("Compile failed!")
		return
	}
	start := time.Now()
	L1 := lua.NewState()
	elapse := time.Now().Sub(start)
	fmt.Println("Golang lua create vm Time consumed:\n", elapse)

	start = time.Now()
	lfunc := L1.NewFunctionFromProto(proto)
	L1.Push(lfunc)
	L1.PCall(0, lua.MultRet, nil)
	elapse = time.Now().Sub(start)
	fmt.Println("Golang lua run bytecode Time consumed:\n", elapse)
	fmt.Println("\n\n")

	L1.Close()
}

func cgo_lua_test() {
	start := time.Now()
	l := C.luaL_newstate()
	C.luaL_openlibs(l)
	s := C.CString(`print("Hello World!")`)
	ret := C.luaL_loadstring(l, s)
	elapse := time.Now().Sub(start)
	if ret == 0 {
		fmt.Println("Cgo lua prepare Time consumed:\n", elapse)
		start = time.Now()
		C.lua_pcall(l, 0, -1, 0)
	}
	elapse = time.Now().Sub(start)
	fmt.Println("Cgo lua run bytecode Time consumed:\n", elapse)
	fmt.Println("\n\n")

	C.lua_close(l)
	C.free(unsafe.Pointer(s))
}

func main() {
	fmt.Println("Begin Test!")
	cgo_test()
	golang_test()
	golang_lua_test()
	cgo_lua_test()
	golang_lua_precompile_test()
	golang_lua_bytecode_run_test()
	fmt.Println("End Test!")
}

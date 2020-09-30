package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"os"
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
#include <string.h>

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

int find(char *payload, char flag)
{
	int i = 0;
	int pLen = strlen(payload);

	for( ; i < pLen; i++)
	{
		if(payload[i] == flag)
		{
			break;
		}
	}
}
*/
import "C"

const (
	REPEAT_CNT = 4
)

var luaStatement string = `print("Hello World!")`

func sum(boundary int) int {
	var total int = 0
	var i int = 0
	for i = 0; i < boundary; i++ {
		total += i
	}

	return total
}

func golang_string_test() {
	loopCnt := 1
	for i := 0; i < REPEAT_CNT; i++ {
		loopCnt *= 10
		start := time.Now()
		for j := 0; j < loopCnt; j++ {
			strings.IndexByte("testdwefweefsewfewfewwwfwewfefwfweweoojooojoljojoefwefez", 'z')
		}
		elapse := time.Now().Sub(start)
		fmt.Printf("Golang string Time consumed:[%s]\n", elapse.String())
	}
	fmt.Println("\n")
}

func cgo_string_test() {
	loopCnt := 1
	for i := 0; i < REPEAT_CNT; i++ {
		loopCnt *= 10
		start := time.Now()
		for j := 0; j < loopCnt; j++ {
			C.find(C.CString("testdwefweefsewfewfewwwfwewfefwfweweoojooojoljojoefwefez"), C.char('z'))
		}
		elapse := time.Now().Sub(start)
		fmt.Printf("Cgo string Time consumed:[%s]\n", elapse.String())
	}
	fmt.Println("\n")
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
		fmt.Printf("Golang Time consumed:[%s]\n", elapse.String())
	}
	fmt.Println("\n")
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
		fmt.Printf("Cgo Time consumed:[%s]\n", elapse.String())
	}
	fmt.Println("\n")
}

func golang_lua_test() {
	start := time.Now()
	L := lua.NewState()
	if err := L.DoString(luaStatement); err != nil {
		panic(err)
	}
	elapse := time.Now().Sub(start)
	fmt.Printf("Golang lua Time consumed:[%s]\n", elapse.String())
	fmt.Println("\n")

	L.Close()
}

func golang_lua_precompile_test() {
	start := time.Now()
	r := strings.NewReader(luaStatement)
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
	fmt.Printf("Golang lua with precompile Time consumed:[%s]\n", elapse.String())
	fmt.Println("\n")

	L1.Close()
}

func golang_lua_bytecode_run_test() {
	r := strings.NewReader(luaStatement)
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
	fmt.Printf("Golang lua create vm Time consumed:[%s]\n", elapse.String())

	start = time.Now()
	lfunc := L1.NewFunctionFromProto(proto)
	L1.Push(lfunc)
	L1.PCall(0, lua.MultRet, nil)
	elapse = time.Now().Sub(start)
	fmt.Printf("Golang lua run bytecode Time consumed:[%s]\n", elapse.String())
	fmt.Println("\n")

	L1.Close()
}

func cgo_lua_test() {
	start := time.Now()
	l := C.luaL_newstate()
	C.luaL_openlibs(l)
	s := C.CString(luaStatement)
	ret := C.luaL_loadstring(l, s)
	elapse := time.Now().Sub(start)
	if ret == 0 {
		fmt.Printf("Cgo lua prepare Time consumed:[%s]\n", elapse.String())
		start = time.Now()
		C.lua_pcall(l, 0, -1, 0)
	}
	elapse = time.Now().Sub(start)
	fmt.Printf("Cgo lua run bytecode Time consumed:[%s]\n", elapse.String())
	fmt.Println("\n")

	C.lua_close(l)
	C.free(unsafe.Pointer(s))
}

func main() {
	if len(os.Args) > 1 {
		luaStatement = string(os.Args[1])
		fmt.Println("Lua statement:[%s]", luaStatement)
	}

	fmt.Println("Begin Test!")
	cgo_test()
	golang_test()
	cgo_string_test()
	golang_string_test()
	golang_lua_test()
	cgo_lua_test()
	golang_lua_precompile_test()
	golang_lua_bytecode_run_test()
	fmt.Println("End Test!")
}

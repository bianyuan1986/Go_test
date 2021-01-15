package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"os"
	"runtime"
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
	SIZE       = 1024 * 1024
)

var luaStatement string = "s = 'select trunc( categoryid/pkg_roles.Fgetmancat(2)) as CATEGORYID, sum(nvl(a.salevalue,0)) as salevalue, sum(nvl(a.netsalevalue,0)) as netsalevalue, sum(nvl(a.salevalueratio,0)) as salevalueratio, sum(nvl(a.costvalue,0)) as costvalue, sum(nvl(a.netcostvalue,0)) as netcostvalue, sum(nvl(a.grossprofit,0)) as grossprofit, sum(nvl(a.grossprofit,0))/f0tonull( sum(nvl(a.salevalue,0)))*100 as grossmargin, sum(nvl(a.grossprofitratio,0)) as grossprofitratio,salechannel from TMP_RPT_M85151006 a where 1=1 group by salechannel, trunc( categoryid/pkg_roles.Fgetmancat(2)) order by trunc( categoryid/pkg_roles.Fgetmancat(2))'; if string.sub(s,1,6)=='select' and string.find(s,'where') then ok = 1 end"
var L1 *lua.LState

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

func golang_lua_engine_memory_test(count int) {
	var st1 runtime.MemStats
	var st2 runtime.MemStats
	runtime.ReadMemStats(&st1)
	fmt.Printf("Count:%d\n", count)
	fmt.Printf("Sys:[%7.2f]M HeapSys:[%7.2f]M HeapIdle:[%7.2f]M HeapInuse:[%7.2f]M StackSys:[%7.2f]M StackInuse:[%7.2f]M\n", float64(st1.Sys)/SIZE, float64(st1.HeapSys)/SIZE, float64(st1.HeapIdle)/SIZE, float64(st1.HeapInuse)/SIZE, float64(st1.StackSys)/SIZE, float64(st1.StackInuse)/SIZE)
	for i := 0; i < count; i++ {
		L1 = lua.NewState()
		L1.DoString("s = 'select where'; if string.sub(s,1,6)=='select' and string.find(s,'where') then ok = 1 end")
	}
	runtime.ReadMemStats(&st2)
	fmt.Printf("Sys:[%7.2f]M HeapSys:[%7.2f]M HeapIdle:[%7.2f]M HeapInuse:[%7.2f]M StackSys:[%7.2f]M StackInuse:[%7.2f]M\n", float64(st2.Sys)/SIZE, float64(st2.HeapSys)/SIZE, float64(st2.HeapIdle)/SIZE, float64(st2.HeapInuse)/SIZE, float64(st2.StackSys)/SIZE, float64(st2.StackInuse)/SIZE)
	fmt.Printf("\n\n")
}

func memory_test() {
	count := 10
	for i := 1; i < 4; i++ {
		count *= 10
		golang_lua_engine_memory_test(count)
	}
}

func GetMatchRuleId(l *lua.LState) int {
	fmt.Println("Match rule id is 2020")
	l.Push(lua.LNumber(2020))

	return 1
}

func register_func() {
	L := lua.NewState()
	tbl := L.NewTable()
	L.SetGlobal("waf", tbl)
	L.SetField(tbl, "GetMatchRuleId", L.NewFunction(GetMatchRuleId))
	if err := L.DoString("id = waf.GetMatchRuleId(); print(id); if id == \"2020\" then print(\"sql\") else print(\"not sql\") end"); err != nil {
		fmt.Println("Failed")
	}
}

func main() {
	if len(os.Args) > 1 {
		luaStatement = string(os.Args[1])
		fmt.Println("Lua statement:[%s]", luaStatement)
	}

	fmt.Println("Begin Test!")

	/*cgo_test()
	golang_test()
	cgo_string_test()
	golang_string_test()
	golang_lua_test()*/
	cgo_lua_test()
	/*golang_lua_precompile_test()*/
	golang_lua_bytecode_run_test()
	memory_test()

	fmt.Println("End Test!")
}

package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"os"
)

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
	if err := L.DoString("id = waf.GetMatchRuleId(); print(id); if id == 2020 then print(\"sql\") else print(\"not sql\") end"); err != nil {
		fmt.Println("Failed")
	}
}

func main() {
	if len(os.Args) > 1 {
		luaStatement := string(os.Args[1])
		fmt.Println("Lua statement:[%s]", luaStatement)
	}

	fmt.Println("Begin Test!")

	register_func()

	fmt.Println("End Test!")
}

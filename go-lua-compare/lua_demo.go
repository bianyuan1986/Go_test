package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"net"
	"net/http"
	"os"
	"strings"
)

var globalName string = "One"
var globalLuaName lua.LString = "Four"
var L *lua.LState = nil
var reqUd *lua.LUserData = nil
var luaByteCode map[string]*lua.FunctionProto

func GetMatchRuleId(l *lua.LState) int {
	fmt.Println("Match rule id is 2020")
	l.Push(lua.LNumber(2020))

	return 1
}

func GetRuleId(l *lua.LState) int {
	id := l.ToInt(1)
	l.Push(lua.LNumber(id))

	return 1
}

func GetName(l *lua.LState) int {
	name := l.GetGlobal("name").(lua.LString)
	fmt.Println(name)
	l.Push(name)

	return 1
}

func GetLuaName(l *lua.LState) int {
	ud := l.GetGlobal("luaName").(*lua.LUserData)
	name := ud.Value.(*lua.LString)
	fmt.Println(name.String())
	l.Push(*name)

	return 1
}

const (
	FIELD_URI        = 0
	FIELD_HOST       = 1
	FIELD_COOKIE     = 2
	FIELD_USER_AGENT = 3
	FIELD_CONNECTION = 4
	FIELD_UID        = 5
)

func GetHttpField(l *lua.LState) int {
	field := l.ToInt(1)
	pri := l.GetGlobal("req").(*lua.LUserData)
	r := pri.Value.(*http.Request)
	result := ""

	/*fmt.Println(r.Header)
	r.ParseForm()
	fmt.Println(r.Form)*/
	switch field {
	case FIELD_URI:
		result = r.RequestURI
	case FIELD_HOST:
		result = r.Host
	case FIELD_COOKIE:
		result = r.Header["Cookie"][0]
	case FIELD_USER_AGENT:
		result = r.Header["User-Agent"][0]
	case FIELD_CONNECTION:
		result = r.Header["Connection"][0]
	case FIELD_UID:
		result = r.Header["Uid"][0]
	default:
		fmt.Println("Unknown!")
	}

	l.Push(lua.LString(result))

	return 1
}

func InitLua() {
	L = lua.NewState()

	/*Just for test!*/
	ud := L.NewUserData()
	ud.Value = &globalLuaName
	L.SetGlobal("name", lua.LString(globalName))
	L.SetGlobal("luaName", ud)

	reqUd = L.NewUserData()
	L.SetGlobal("req", reqUd)
	luaByteCode = make(map[string]*lua.FunctionProto)
	registerModuleV1("waf", L)
}

func TestLua() {
	if err := L.DoString("id = waf.GetMatchRuleId(); print(id); if id == 2020 then print(\"2020 sql\") else print(\"not sql\") end"); err != nil {
		fmt.Println("Failed")
	}
	if err := L.DoString("tmp = 2021; id = waf.GetRuleId(tmp); print(id); if id == 2021 then print(\"2021 sql\") else print(\"not sql\") end"); err != nil {
		fmt.Println("Failed")
	}
	if err := L.DoString("name = waf.GetLuaName(); print(\"luaname:\"..name)"); err != nil {
		fmt.Println("Failed")
	}
	globalLuaName = "Five"
	if err := L.DoString("name = waf.GetLuaName(); print(\"luaname:\"..name)"); err != nil {
		fmt.Println("Failed")
	}
}

var exports = map[string]lua.LGFunction{
	"GetMatchRuleId": GetMatchRuleId,
	"GetRuleId":      GetRuleId,
	"GetName":        GetName,
	"GetLuaName":     GetLuaName,
	"GetHttpField":   GetHttpField,
}

func registerModuleV1(name string, L *lua.LState) int {
	tbl := L.NewTable()
	L.SetGlobal(name, tbl)
	for name, handler := range exports {
		L.SetField(tbl, name, L.NewFunction(handler))
	}

	return 0
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)

	return 1
}

func registerModuleV2(name string) int {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule(name, Loader)

	return 1
}

func bianyuan_test(w http.ResponseWriter, req *http.Request) {
	reqUd.Value = req
	if err := L.DoString("uri = waf.GetHttpField(0); host = waf.GetHttpField(1); print(\"uri:\"..uri); print(\"host:\"..host)"); err != nil {
		fmt.Println("Failed")
	}
}

func tiger_test(w http.ResponseWriter, req *http.Request) {
	reqUd.Value = req
	uid := req.Header["Uid"]
	if uid == nil {
		fmt.Println("Lack uid info!")
		return
	}
	proto := luaByteCode[uid[0]]
	if proto == nil {
		fmt.Println("No script!")
		return
	}
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	L.PCall(0, lua.MultRet, nil)
}

func NewServer(addr string) bool {
	mux := http.NewServeMux()
	mux.HandleFunc("/bianyuan_test", bianyuan_test)
	mux.HandleFunc("/tiger_test", tiger_test)
	s := http.Server{Addr: addr, Handler: mux, ConnState: nil}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Listen failed!")
		return false
	}
	s.Serve(l)

	return true
}

func LoadLuaScript(scripts map[string]string) bool {
	for name, content := range scripts {
		r := strings.NewReader(content)
		chunk, err := parse.Parse(r, "tiger")
		if err != nil {
			fmt.Println("Parse failed!")
			return false
		}
		proto, err := lua.Compile(chunk, name)
		if err != nil {
			fmt.Println("Compile failed!")
			return false
		}
		luaByteCode[name] = proto
	}

	/*fmt.Println(luaByteCode)*/
	return true
}

func main() {
	if len(os.Args) > 1 {
		luaStatement := string(os.Args[1])
		fmt.Println("Lua statement:[%s]", luaStatement)
	}

	luaScripts := map[string]string{
		"202020": "uri = waf.GetHttpField(0); print(\"uri:\"..uri)",
		"202021": "host = waf.GetHttpField(1); print(\"host:\"..host)",
		"202022": "cookie = waf.GetHttpField(2); print(\"cookie:\"..cookie)",
		"202023": "agent = waf.GetHttpField(3); print(\"user-agent:\"..agent)",
		"202024": "connection = waf.GetHttpField(4); print(\"connection:\"..connection)",
		"202025": "uid = waf.GetHttpField(5); print(\"uid:\"..uid)",
	}

	InitLua()
	LoadLuaScript(luaScripts)
	/*TestLua()*/
	NewServer("30.27.153.142:2020")
}

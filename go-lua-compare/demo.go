package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"net"
	"net/http"
	"strings"
)

var L *lua.LState = nil
var reqUd *lua.LUserData = nil
var luaByteCode map[string]*lua.FunctionProto
var globalResult string = ""

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

func SetResult(l *lua.LState) int {
	globalResult = l.ToString(1)

	return 0
}

func InitLua() {
	L = lua.NewState()

	reqUd = L.NewUserData()
	L.SetGlobal("req", reqUd)
	luaByteCode = make(map[string]*lua.FunctionProto)
	registerModuleV1("waf", L)
}

var exports = map[string]lua.LGFunction{
	"GetHttpField": GetHttpField,
	"SetResult":    SetResult,
}

func registerModuleV1(name string, L *lua.LState) int {
	tbl := L.NewTable()
	L.SetGlobal(name, tbl)
	for name, handler := range exports {
		L.SetField(tbl, name, L.NewFunction(handler))
	}

	return 0
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
	fmt.Fprintf(w, globalResult)
}

func NewServer(addr string) bool {
	mux := http.NewServeMux()
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
	luaScripts := map[string]string{
		"202020": "uri = waf.GetHttpField(0); result = \"uri:\"..uri; waf.SetResult(result); print(result)",
		"202021": "host = waf.GetHttpField(1); result = \"host:\"..host; waf.SetResult(result); print(result)",
		"202022": "cookie = waf.GetHttpField(2); result = \"cookie:\"..cookie; waf.SetResult(result); print(result)",
		"202023": "agent = waf.GetHttpField(3); result = \"user-agent:\"..agent; waf.SetResult(result); print(result)",
		"202024": "connection = waf.GetHttpField(4); result = \"connection:\"..connection; waf.SetResult(result); print(result)",
		"202025": "uid = waf.GetHttpField(5); result = \"uid:\"..uid; waf.SetResult(result)print(result)",
	}

	InitLua()
	LoadLuaScript(luaScripts)
	NewServer("30.27.153.132:2020")
}

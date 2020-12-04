package lua

type LuaHandler interface {
	ExecuteLua(script string)
}

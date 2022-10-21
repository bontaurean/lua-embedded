package handlers

import (
	lua "github.com/yuin/gopher-lua"
)

var values map[string]int

var exports = map[string]lua.LGFunction{
	"load":    loadValue,
	"store":   storeValue,
	"process": processValue,
}

func init() {
	values = make(map[string]int)
}

func Configure(l *lua.LState) int {
	l.Push(l.SetFuncs(l.NewTable(), exports))
	return 1
}

func loadValue(l *lua.LState) int {
	k := l.ToString(1)
	v, ok := values[k]
	if !ok {
		v = 0
		values[k] = v
	}
	l.Push(lua.LNumber(v))
	return 1
}

func storeValue(l *lua.LState) int {
	k, v := l.ToString(1), l.ToInt(2)
	values[k] = v
	return 0
}

func processValue(l *lua.LState) int {
	v := l.ToInt(1)
	v += 1
	l.Push(lua.LNumber(v))
	return 1
}

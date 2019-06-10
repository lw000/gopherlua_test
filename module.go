package main

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"myfunc": myFunc,
}

// Loader 加载库函数
func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func myFunc(L *lua.LState) int {
	v := L.ToString(-1)
	L.Push(lua.LString(v))

	log.Println("test myfun")

	return 1
}

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

var exports = map[string]lua.LGFunction{
	"myfunc": myFunc,
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func myFunc(L *lua.LState) int {
	v := L.ToInt(-1)
	L.Push(lua.LString(v))

	log.Println("test myfun")

	return 1
}

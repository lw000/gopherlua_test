package main

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

// LuaDouble ...
func LuaDouble(L *lua.LState, a lua.LValue) int {
	er := L.CallByParam(lua.P{Fn: L.GetGlobal("double"), NRet: 1, Protect: true}, a)
	if er != nil {
		log.Println(er)
		return 1
	}
	ret := L.Get(-1)
	L.Pop(-1)

	num, ok := ret.(lua.LNumber)
	if ok {
		return int(num)
	}
	return -1
}

// LuaMax ...
func LuaMax(L *lua.LState, a, b lua.LValue) int {
	er := L.CallByParam(lua.P{Fn: L.GetGlobal("max"), NRet: 1, Protect: true}, a, b)
	if er != nil {
		log.Println(er)
		return 1
	}
	v := L.Get(-1)
	L.Pop(-1)

	max, ok := v.(lua.LNumber)
	if ok {
		return int(max)
	}
	return -1
}

// LuaMin ...
func LuaMin(L *lua.LState, a, b lua.LValue) int {
	er := L.CallByParam(lua.P{Fn: L.GetGlobal("min"), NRet: 1, Protect: true}, a, b)
	if er != nil {
		log.Println(er)
		return 1
	}
	v := L.Get(-1)
	L.Pop(-1)

	min, ok := v.(lua.LNumber)
	if ok {
		return int(min)
	}
	return -1
}

// LuaMaxmin ...
func LuaMaxmin(L *lua.LState, a, b lua.LValue) (max int, min int) {
	er := L.CallByParam(lua.P{Fn: L.GetGlobal("maxmin"), NRet: 2, Protect: true}, a, b)
	if er != nil {
		log.Println(er)
		return 0, 0
	}

	n1 := L.Get(-2)
	L.Pop(-2)

	v1, ok := n1.(lua.LNumber)
	if ok {
		max = int(v1)
	}

	n2 := L.Get(-1)
	L.Pop(-1)

	v2, ok := n2.(lua.LNumber)
	if ok {
		min = int(v2)
	}
	return max, min
}

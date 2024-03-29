package main

import (
	"log"
	"sync"
	"time"

	lua "github.com/yuin/gopher-lua"
)

// GoDouble ...
func GoDouble(L *lua.LState) int {
	v := L.ToInt(1)
	L.Push(lua.LNumber(v * 2))
	return 1
}

// Update ...
func Update(wg *sync.WaitGroup, L *lua.LState) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	t := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-t.C:
			er := L.CallByParam(lua.P{
				Fn:      L.GetGlobal("update"),
				NRet:    0,
				Protect: true,
			}, lua.LNumber(time.Now().UnixNano()/1e6))
			if er != nil {
				log.Println(er)
				return
			}
		}
	}
}

func main() {
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("levi", Loader)
	L.SetGlobal("GoDouble", L.NewFunction(GoDouble))

	if er := L.DoFile("./script/main.lua"); er != nil {
		log.Panic(er)
	}

	log.Println(LuaDouble(L, lua.LNumber(10)))
	log.Println(LuaMax(L, lua.LNumber(100), lua.LNumber(200)))
	log.Println(LuaMin(L, lua.LNumber(100), lua.LNumber(200)))
	log.Println(LuaMaxmin(L, lua.LNumber(100), lua.LNumber(200)))

	// wg := &sync.WaitGroup{}
	// wg.Add(1)
	// go Update(wg, L)
	// wg.Wait()
}

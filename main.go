package main

import (
	"github.com/yuin/gopher-lua"
	"log"
	"sync"
	"time"
)

func GoDouble(L *lua.LState) int {
	v := L.ToInt(1)
	L.Push(lua.LNumber(v * 2))
	return 1
}

func RunLuaApp() {
	L := lua.NewState()
	defer L.Close()

	L.OpenLibs()

	L.PreloadModule("levi", Loader)

	L.SetGlobal("GoDouble", L.NewFunction(GoDouble))

	if er := L.DoFile("./main.lua"); er != nil {
		log.Panic(er)
	}

	LuaDouble(L, lua.LNumber(10))
	LuaMax(L, lua.LNumber(100), lua.LNumber(200))
	LuaMin(L, lua.LNumber(100), lua.LNumber(200))
	log.Println(LuaMaxmin(L, lua.LNumber(100), lua.LNumber(200)))

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go Update(wg, L)
	wg.Wait()
}

func Update(wg *sync.WaitGroup, L *lua.LState) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	t := time.NewTicker(time.Millisecond * 1000)
	for {
		select {
		case <-t.C:
			er := L.CallByParam(lua.P{
				Fn:      L.GetGlobal("update"),
				NRet:    1,
				Protect: true,
			}, lua.LNumber(time.Now().Unix()))
			if er != nil {
				log.Println(er)
				return
			}
			ret := L.Get(-1)
			L.Pop(-1)

			res, ok := ret.(lua.LNumber)
			if ok {
				log.Println(int(res))
			}
		}
	}
}

func main() {
	RunLuaApp()
}

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
	"sync"
	"time"
)

var (
	wg *sync.WaitGroup
)

func GoDouble(L *lua.LState) int {
	lv := L.ToInt(1)
	L.Push(lua.LNumber(lv * lv))
	return 1
}

func LuaCallGo() {
	defer func() {
		wg.Done()
	}()

	L := lua.NewState()
	defer L.Close()

	L.OpenLibs()

	L.SetGlobal("GoDouble", L.NewFunction(GoDouble))

	if er := L.DoFile("./main.lua"); er != nil {
		log.Panic(er)
	}

	er := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("Double"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(10))
	if er != nil {
		log.Panic(er)
	}
	ret := L.Get(-1)
	L.Pop(-1)

	res, ok := ret.(lua.LNumber)
	if ok {
		log.Println(int(res))
	}

	wg = &sync.WaitGroup{}

	wg.Add(1)
	go Update(L)
	wg.Wait()
}

func Update(L *lua.LState) {
	defer func() {
		wg.Done()
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
				log.Panic(er)
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
	LuaCallGo()
}

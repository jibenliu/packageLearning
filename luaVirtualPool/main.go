package main

import (
	"bufio"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"os"
	"sync"
)

type lStatePool struct {
	m     sync.Mutex
	saved []*lua.LState
}

func (pl *lStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

func (pl *lStatePool) New() *lua.LState {
	L := lua.NewState()
	// setting the L up here.
	// load scripts, set global variables, share channels, etc...
	return L
}

func (pl *lStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}

func (pl *lStatePool) Shutdown() {
	for _, L := range pl.saved {
		L.Close()
	}
}

// Global LState pool
var luaPool = &lStatePool{
	saved: make([]*lua.LState, 0, 4),
}

func MyWorker() {
	L := luaPool.Get()
	defer luaPool.Put(L)
	/* your code here */
}

// CompileLua reads the passed lua file from disk and compiles it.
func CompileLua(filePath string) (*lua.FunctionProto, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	chunk, err := parse.Parse(reader, filePath)
	if err != nil {
		return nil, err
	}
	proto, err := lua.Compile(chunk, filePath)
	if err != nil {
		return nil, err
	}
	return proto, nil
}

// DoCompiledFile takes a FunctionProto, as returned by CompileLua, and runs it in the LState. It is equivalent
// to calling DoFile on the LState with the original source file.
func DoCompiledFile(L *lua.LState, proto *lua.FunctionProto) error {
	lfunc := L.NewFunctionFromProto(proto)
	L.Push(lfunc)
	return L.PCall(0, lua.MultRet, nil)
}

func GetLuaPath(L *lua.LState) int {
	// 绝对路径
	L.Push(lua.LString("./scripts/lua"))
	return 1
}

// Example shows how to share the compiled byte code from a lua script between multiple VMs.
func Example() {
	codeToShare, _ := CompileLua("mylua.lua")
	a := lua.NewState()
	b := lua.NewState()
	c := lua.NewState()
	_ = DoCompiledFile(a, codeToShare)
	_ = DoCompiledFile(b, codeToShare)
	_ = DoCompiledFile(c, codeToShare)
}

func main() {
	defer luaPool.Shutdown()
	go MyWorker()
	go MyWorker()
	/* etc... */
}

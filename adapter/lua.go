package adapter

import (
	"encoding/base64"
	"fmt"

	"github.com/aarzilli/golua/lua"
)

// TODO: aSec sync is enough?
type Lua struct {
	state *lua.State
}

type LuaParams struct {
	LuaCode64 string
}

func NewLua(params LuaParams) (*Lua, error) {
	luaCode, err := base64.StdEncoding.DecodeString(params.LuaCode64)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(luaCode))

	state := lua.NewState()
	state.OpenLibs()

	// Load and execute the Lua code
	if err := state.DoString(string(luaCode)); err != nil {
		return nil, fmt.Errorf("Error loading Lua code: %v\n", err)
	}
	return &Lua{state: state}, nil
}

func (l *Lua) Close() {
	l.state.Close()
}

func (l *Lua) Generate(data []byte, counter int64) ([]byte, error) {
	// Push the function onto the stack
	l.state.GetGlobal("D_gen")

	// Push the argument
	l.state.PushBytes(data)
	l.state.PushInteger(counter)

	if err := l.state.Call(2, 1); err != nil {
		return nil, fmt.Errorf("Error calling Lua function: %v\n", err)
	}

	result := l.state.ToBytes(-1)
	l.state.Pop(1)

	fmt.Printf("Result: %s\n", string(result))
	return result, nil
}

func (l *Lua) Parse(data []byte) ([]byte, error) {
	// Push the function onto the stack
	l.state.GetGlobal("D_parse")

	// Push the argument
	l.state.PushBytes(data)
	if err := l.state.Call(1, 1); err != nil {
		return nil, fmt.Errorf("Error calling Lua function: %v\n", err)
	}

	result := l.state.ToBytes(-1)
	l.state.Pop(1)

	fmt.Printf("Result: %s\n", string(result))
	return result, nil
}

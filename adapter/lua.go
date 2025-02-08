package main

import (
	"encoding/base64"
	"fmt"

	"github.com/aarzilli/golua/lua"
)

func main() {
	// luaB64 := `bG9jYWwgZnVuY3Rpb24gZF9nZW4oZGF0YSwgY291bnRlcikKCUhlYWRlciA9IHN0cmluZy5jaGFyKDB4MTIsIDB4MzQsIDB4NTYsIDB4NzgpCgktLSBsb2NhbCB0cyA9IG9zLnRpbWUoKQoJcmV0dXJuIEhlYWRlciAuLiBkYXRhCmVuZAoKbG9jYWwgZnVuY3Rpb24gZF9wYXJzZShkYXRhKQoJcmV0dXJuIHN0cmluZy5zdWIoZGF0YSwgI0hlYWRlcikKZW5kCg==`
	// only d_gen
	// luaB64 := `ZnVuY3Rpb24gRF9nZW4oZGF0YSkKCS0tIEhlYWRlciA9IHN0cmluZy5jaGFyKDB4MTIsIDB4MzQsIDB4NTYsIDB4NzgpCglsb2NhbCBIZWFkZXIgPSAiXHgxMlx4MzRceDU2XHg3OCIKCS0tIGxvY2FsIHRzID0gb3MudGltZSgpCglyZXR1cm4gSGVhZGVyIC4uIGRhdGEKZW5kCg==`
	luaB64 := `ZnVuY3Rpb24gRF9nZW4oZGF0YSkKCS0tIEhlYWRlciA9IHN0cmluZy5jaGFyKDB4MTIsIDB4MzQsIDB4NTYsIDB4NzgpCglsb2NhbCBIZWFkZXIgPSAiXHgxMlx4MzRceDU2XHg3OCIKCWxvY2FsIHRzID0gb3MudGltZSgpCglyZXR1cm4gSGVhZGVyIC4uIGRhdGEKZW5kCg==`
	sDec, _ := base64.StdEncoding.DecodeString(luaB64)
	fmt.Println(string(sDec))
	luaCode := sDec
	L := lua.NewState()
	L.OpenLibs()
	defer L.Close()

	// Load and execute the Lua code
	if err := L.DoString(string(luaCode)); err != nil {
		fmt.Printf("Error loading Lua code: %v\n", err)
		return
	}

	// Push the function onto the stack
	L.GetGlobal("D_gen")

	// Push the argument
	L.PushString("data")

	if err := L.Call(1, 1); err != nil {
		fmt.Printf("Error calling Lua function: %v\n", err)
		return
	}

	result := L.ToString(-1)
	L.Pop(1)

	// Print the result
	// fmt.Printf("Result: %x\n", []byte(result))
	fmt.Printf("Result: %s\n", result)
}

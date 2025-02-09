package adapter

import (
	"testing"
)

func newLua() *Lua {
	lua, _ := NewLua(LuaParams{
		/*
			function d_gen(msg_type, data, counter)
				local header = "header"
				return counter .. header .. data
			end

			function d_parse(data)
				local header = "10header"
				return string.sub(data, #header+1)
			end
		*/
		Base64LuaCode: "CmZ1bmN0aW9uIGRfZ2VuKG1zZ190eXBlLCBkYXRhLCBjb3VudGVyKQoJbG9jYWwgaGVhZGVyID0gImhlYWRlciIKCXJldHVybiBjb3VudGVyIC4uIGhlYWRlciAuLiBkYXRhCmVuZAoKZnVuY3Rpb24gZF9wYXJzZShkYXRhKQoJbG9jYWwgaGVhZGVyID0gIjEwaGVhZGVyIgoJcmV0dXJuIHN0cmluZy5zdWIoZGF0YSwgI2hlYWRlcisxKQplbmQK",
	})
	return lua
}

func TestLua_Generate(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := newLua()
		defer l.Close()
		got, err := l.Generate(1, []byte("test"))
		if err != nil {
			t.Errorf(
				"Lua.Generate() error = %v, wantErr %v",
				err,
				nil,
			)
			return
		}

		want := "1headertest"
		if string(got) != want {
			t.Errorf("Lua.Generate() = %v, want %v", string(got), want)
		}
	})
}

func TestLua_Parse(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := newLua()
		defer l.Close()
		got, err := l.Parse([]byte("10headertest"))
		if err != nil {
			t.Errorf("Lua.Parse() error = %v, wantErr %v", err, nil)
			return
		}
		want := "test"
		if string(got) != want {
			t.Errorf("Lua.Parse() = %v, want %v", got, want)
		}
	})
}

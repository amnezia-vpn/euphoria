package adapter

import (
	"testing"
)

func newLua() *Lua {
	lua, _ := NewLua(LuaParams{
		/*
		   function d_gen(data, counter)
		   	local header = "header"
		   	return counter .. header .. data
		   end

		   function d_parse(data)
		   	local header = "10header"
		   	return string.sub(data, #header+1)
		   end
		*/
		Base64LuaCode: "ZnVuY3Rpb24gZF9nZW4oZGF0YSwgY291bnRlcikKCWxvY2FsIGhlYWRlciA9ICJoZWFkZXIiCglyZXR1cm4gY291bnRlciAuLiBoZWFkZXIgLi4gZGF0YQplbmQKCmZ1bmN0aW9uIGRfcGFyc2UoZGF0YSkKCWxvY2FsIGhlYWRlciA9ICIxMGhlYWRlciIKCXJldHVybiBzdHJpbmcuc3ViKGRhdGEsICNoZWFkZXIrMSkKZW5kCg==",
	})
	return lua
}

func TestLua_Generate(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l := newLua()
		defer l.Close()
		got, err := l.Generate([]byte("test"), 10)
		if err != nil {
			t.Errorf(
				"Lua.Generate() error = %v, wantErr %v",
				err,
				nil,
			)
			return
		}

		want := "10headertest"
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

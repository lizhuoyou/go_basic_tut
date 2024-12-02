package c9string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // initial value should be empty
	s = "hello"
	t.Log(len(s))
	// s[1] = "3" invalid since since string is immutable byte slice
	s = "\xE4\xB8\xA5"	// string can store any binary data
	t.Log(s)
	s = "中"
	t.Log(len(s))

	c:=[]rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		// t.Log(string(c), c)
		t.Logf("%[1]c %[1]x", c)
	}
}

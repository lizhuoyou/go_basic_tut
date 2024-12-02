package c9string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for i, part := range parts {
		t.Log(i, part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringToNum(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str"+ s)
	if num, err := strconv.Atoi(s); err == nil {
		t.Log(num * 10)
	}
}

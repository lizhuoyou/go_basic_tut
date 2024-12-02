package c8map

import (
	"math"
	"testing"
)

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int)int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return int(math.Pow(float64(op), 2))}
	m[3] = func(op int) int {return int(math.Pow(float64(op), 3))}
	t.Log(m[1](2), m[2](2), m[3](2))
}

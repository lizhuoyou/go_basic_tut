package c10func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// kep points:
// 1: in go, there could be multiple return values
// 2: in Go, all parameters are passed as value instead of reference
// 3: in go, a variable can be assigned a func as value
// 4: a func can be a paramter and a return value

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int)int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("Time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int{
	time.Sleep(time.Second * 1)
	return op * op
}

func TestFunc(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)

	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}

func sum(ops ...int) int {
	res := 0
	for _, op := range ops {
		res += op
	}
	return res
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("Clear resource")
	}()
	fmt.Println("Started")
	panic("Fatal error")	// will end, but defere will be executed
}

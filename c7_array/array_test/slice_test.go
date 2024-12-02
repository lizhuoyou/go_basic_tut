package array_test

import (
	"math"
	"testing"
)

func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	t.Log(append(s1, 6))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
}

func TestSliceGrowing(t *testing.T) {
	s0 := []int{}
	for i := 0; i < 10; i++ {
		s0 = append(s0, i)
		t.Log(len(s0), cap(s0))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := [12]int{}
	for i := range 12 {
		year[i] = i + 1
	}
	t.Log(year)

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = math.MaxInt
	t.Log(Q2)
}

func TestSliceComp(t *testing.T) {
	a := []int{1, 2, 3}
	// the following is invalid since slice can only be compared to nil
	// b := []int{2, 3, 4}
	// if a == b {
	// 	t.Log("equal")
	// }
	if len(a) > 0 {
		t.Logf("%v is not empty", a)
	}
}

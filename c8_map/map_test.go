package c8map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1:1, 2:4}
	t.Log(m1[2])
	t.Logf("len m1 is %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2 is %d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3 is %d", len(m3))
	// note: map cannot be visited with capacity cap()
}

func TestAccessNonExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Logf("key 3 has value %d", v)
	} else {
		t.Log("key 3 does not exist")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2:4, 3:9}
	for k, v := range m1 {
		t.Logf("%v's square is %v", k, v)
	}
}

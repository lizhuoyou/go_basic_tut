package c8map

import "testing"

func TestSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[3] = true
	for i := range 4 {
		if mySet[i] {
			t.Logf("%v is existing", i)
		} else {
			t.Logf("%v is not existing", i)
		}
	}
	t.Log(len(mySet))

	delete(mySet, 1)
	t.Log(len(mySet))
}

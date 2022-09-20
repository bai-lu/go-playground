package commonmistakes

import "testing"

func TestLoopIterator(t *testing.T) {
	loopIterator()

}

func TestLoopIterator1(t *testing.T) {
	loopIterator1()

}

func TestLoopIterator2(t *testing.T) {
	sliceIsNil()

}

func TestLoopIterator3(t *testing.T) {
	// sliceIsNil1()
	a := make([]int, 0, 100)
	t.Log(len(a))
	t.Log(cap(a))
}

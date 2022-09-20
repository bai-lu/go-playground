package commonmistakes

import "testing"

func TestBar(t *testing.T) {
	bar()
	bar1()

}

func TestA(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := a[:]
	t.Log(b)

}

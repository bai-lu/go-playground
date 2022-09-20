package commonmistakes

import "fmt"

func loopIterator() {
	var out []*int
	for i := 0; i < 3; i++ {
		i := i
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}

func loopIterator1() {
	var out [][]int
	for _, i := range [][1]int{{1}, {2}, {3}} {
		i := i
		out = append(out, i[:])
	}
	fmt.Println("Values:", out)
	fmt.Println("Values:", cap(out))
	var out1 []int
	// out1[0] = 0
	out1 = append(out1, 1)
	if out1 != nil {
		fmt.Println("out1 is nil")
		fmt.Println("out1 is nil", cap(out1))
	}
}

func sliceIsNil() {
	var a []int

	if a == nil {
		fmt.Println("a is nil")
	}
	fmt.Println(a)
}

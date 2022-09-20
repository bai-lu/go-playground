package commonmistakes

import "fmt"

func foo(a [2]int) {
	a[0] = 200
}

func bar() {
	a := [...]int{1, 2}
	foo(a)
	fmt.Println(a)
}

func foo1(a []int) {
	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8)
	a[0] = 200
}

func bar1() {
	// a := [2]int{1, 2}
	a := make([]int, 2, 100)
	a[0] = 10
	a[1] = 11

	foo1(a)
	fmt.Println(a)
}

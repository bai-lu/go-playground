package main

import "fmt"

type Num interface {
	~int | int8 | int16 | float32 | float64
}

func add[B Num](a B, b B) B {
	return a + b
}

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

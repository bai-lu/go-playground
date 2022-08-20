package tmp

import (
	"os"
	"time"
)

func F0() {
	defer func() {
		// recover()
		print("panic F0...\n")
	}()
	_, err := os.Open("/tmp/a.txt")
	if err != nil {
		panic(err)
	}

	println("F1")
	time.Sleep(time.Second * 2)
}

func F1() {
	defer func() {
		// recover()
		print("panic F1...\n")
	}()
	F0()

	println("F1")
}

func F2() {
	defer func() {
		print("panic F2...\n")
	}()
	go F1()

	println("F2")
	time.Sleep(time.Second * 2)
}

func F3() {
	defer func() {
		print("panic F3...\n")
	}()
	go F2()

	println("F3")
	time.Sleep(time.Second * 2)
}

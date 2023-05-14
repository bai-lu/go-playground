package main

import (
	"fmt"
	"sync"
)

func main() {

	go func() {
		// run
		for {

		}
	}()

	// map supports multi-thread read
	data := map[string]string{}
	data["a"] = "a"
	data["b"] = "b"
	data["c"] = "c"

	go fmt.Println(data["a"])
	go fmt.Println(data["b"])
	go fmt.Println(data["c"])
	m := sync.Mutex{}
	m.Lock()
	m.Lock()

	// unreachable code
	m.Unlock()
	m.Unlock()

}

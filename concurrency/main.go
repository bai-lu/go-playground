package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var lock1 sync.Mutex
var lock2 sync.Mutex

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

func printNum() {
	for _, num := range nums {
		lock2.Lock()
		fmt.Println(num)
		lock1.Unlock()
		time.Sleep(time.Second)
	}
}

func printChar() {
	for _, char := range chars {
		lock1.Lock()
		fmt.Println(char)
		lock2.Unlock()
		time.Sleep(time.Second)
	}
}

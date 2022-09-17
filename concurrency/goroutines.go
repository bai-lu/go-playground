package concurrency

import (
	"fmt"
	"log"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func CutSlice() {
	a := []int{}
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}

	log.Println(a)
	windowSize := 10
	windowStart, windowEnd := 0, 0
	for i := 0; i < 11; i++ {
		windowStart = windowEnd
		windowEnd = windowStart + windowSize
		sub := a[windowStart:windowEnd]
		log.Println(sub)
	}

	log.Println(a[100:120])

}

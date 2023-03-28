package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestPrint(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	for i := 0; i < 26; i++ {
		nums = append(nums, i)
		alphabet = append(alphabet, rune(97+i))
	}
	fmt.Println(len(nums), len(alphabet))
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	go func() {
		defer wg.Done()
		for _, i := range nums {
			<-ch1
			fmt.Printf("%d", i)
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for _, i := range alphabet {
			<-ch2
			fmt.Printf("%c", i)
			ch1 <- struct{}{}
		}
	}()
	ch1 <- struct{}{}
	wg.Wait()
}

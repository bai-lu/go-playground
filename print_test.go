// 交替打印数字和字符

package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestCrossPrint(t *testing.T) {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-ch1
			fmt.Print(i)
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 'a'; i <= 'z'; i++ {
			<-ch2
			fmt.Printf("%c", i)
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}
	wg.Wait()
}

func TestCrossPrintUseMutex(t *testing.T) {
	var mtx sync.Mutex
	var wg sync.WaitGroup
	cond := sync.NewCond(&mtx)
	wg.Add(2)

	printNums := func(wg *sync.WaitGroup, cond *sync.Cond, mtx *sync.Mutex) {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			mtx.Lock()
			cond.Signal()
			fmt.Print(i)
			cond.Wait()
			mtx.Unlock()
		}
	}

	printChars := func(wg *sync.WaitGroup, cond *sync.Cond, mtx *sync.Mutex) {
		defer wg.Done()
		for i := 'a'; i <= 'z'; i++ {
			mtx.Lock()
			cond.Wait()
			fmt.Printf("%c", i)
			cond.Signal()
			mtx.Unlock()
		}
	}
	go printNums(&wg, cond, &mtx)
	go printChars(&wg, cond, &mtx)
	wg.Wait()
}

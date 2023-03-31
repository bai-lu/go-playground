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

func printNum(wg *sync.WaitGroup, mtx *sync.Mutex, cond *sync.Cond) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		// 先获取互斥锁
		mtx.Lock()
		// 输出数字
		fmt.Printf("%d ", i)
		// 通知另一个协程可以开始打印字母了
		cond.Signal()
		// 等待另一个协程打印完字母后再继续输出数字
		cond.Wait()
		// 释放互斥锁
		mtx.Unlock()
	}
	// 通知另一个协程结束
	cond.Signal()
}

func printChar(wg *sync.WaitGroup, mtx *sync.Mutex, cond *sync.Cond) {
	defer wg.Done()
	for i := 'A'; i <= 'E'; i++ {
		// 先获取互斥锁
		mtx.Lock()
		// 等待另一个协程输出数字
		cond.Wait()
		// 输出字母
		fmt.Printf("%c ", i)
		// 通知另一个协程可以开始输出数字了
		cond.Signal()
		// 释放互斥锁
		mtx.Unlock()
	}
	// 通知另一个协程结束
	cond.Signal()
}

func TestXxx1(t *testing.T) {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	cond := sync.NewCond(&mtx)
	wg.Add(2)
	go printNum(&wg, &mtx, cond)
	go printChar(&wg, &mtx, cond)
	wg.Wait()
}

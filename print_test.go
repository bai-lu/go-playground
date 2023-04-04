package main

import (
	"fmt"
	"sync"
	"testing"
)

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

func TestA(t *testing.T) {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	cond := sync.NewCond(&mtx)
	wg.Add(2)
	go printNum(&wg, &mtx, cond)
	go printChar(&wg, &mtx, cond)
	wg.Wait()
}

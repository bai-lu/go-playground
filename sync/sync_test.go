package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	sm := sync.Map{}
	sm.Store("a", 1)
	// json.Marshal(sm)

}

func TestB(t *testing.T) {
	count := 0
	m := sync.Mutex{}
	go func() {
		m.Lock()
		count++
		println("in go routine", count)
		time.Sleep(time.Second)
		m.Unlock()
	}()
	m.Lock()
	count++
	println("in main routine", count)
	time.Sleep(3 * time.Second)
	m.Unlock()
	time.Sleep(2 * time.Second)

}

var locker sync.Mutex
var cond = sync.NewCond(&locker)

// NewCond(l Locker)里面定义的是一个接口,拥有lock和unlock方法。
// 看到sync.Mutex的方法,func (m *Mutex) Lock(),可以看到是指针有这两个方法,所以应该传递的是指针
func TestC(t *testing.T) {
	// 启动多个协程
	for i := 0; i < 10; i++ {
		go func(x int) {
			cond.L.Lock()         // 获取锁
			defer cond.L.Unlock() // 释放锁

			cond.Wait() // 等待通知，阻塞当前 goroutine

			// 通知到来的时候, cond.Wait()就会结束阻塞, do something. 这里仅打印
			fmt.Println(x)
		}(i)
	}

	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发一个通知给已经获取锁的 goroutine

	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发下一个通知给已经获取锁的 goroutine

	time.Sleep(time.Second * 1)
	cond.Broadcast() // 1 秒后下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 1) // 等待所有 goroutine 执行完毕
}

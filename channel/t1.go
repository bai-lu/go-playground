package channel

import (
	"fmt"
	"sync"
)

const MaxOutstanding = 100

type Request struct {
	name string
}

func process(r *Request) {
	// 处理请求
	fmt.Println(r.name)
}

func handle(queue chan *Request) {
	for r := range queue {
		process(r)
	}
}

func Serve(clientRequests chan *Request, quit chan bool) {
	// 启动处理程序
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<-quit // 等待通知退出。
	print("exiting")
}

func F1() {
	num := 100

	var wg sync.WaitGroup
	wg.Add(num)

	c := make(chan int) // channel 无限长?
	for i := 0; i < num; i++ {
		k := i
		go func() {
			c <- k // channl是协程安全的, 为什么不阻塞
			c <- 999
			wg.Done()
		}()
	}

	// 等待关闭channel
	go func() {
		wg.Wait()
		close(c)
	}()
	// c <- 999

	// 读取数据
	var a []int
	for i := range c {
		fmt.Println(i)
		a = append(a, i)
	}

	fmt.Println(len(c))
	fmt.Println(len(a))
}

func F2() {
	ch := make(chan struct{})
	done := make(chan struct{})
	l := []struct{}{}
	go func() {
		for {

			select {
			case metric, ok := <-ch:
				if !ok {
					return
				}
				l = append(l, metric)
			case _, ok := <-done:
				if !ok {
					return
				}
				close(ch)
				println("ch closed")
				return // end
				// default:
				// 	time.Sleep(time.Millisecond * 100) // sleep 100ms
			}
		}
	}()

	ch <- struct{}{}
	ch <- struct{}{}
	ch <- struct{}{}

	done <- struct{}{}
	println("done")

	println(len(l))
}

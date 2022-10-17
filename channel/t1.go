package channel

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func F1() {
	// sync.Map{}
	num := 100
	var a []int
	wg, _ := errgroup.WithContext(context.Background())
	c := make(chan int) // channel 无限长?
	for i := 0; i < num; i++ {
		k := i
		wg.Go(func() error {
			c <- k // channl是协程安全的, 为什么不阻塞
			return nil
		})
	}
	go func() {
		// 读取数据
		for i := range c {
			fmt.Println(i)
			a = append(a, i)
		}
		println("closed")
	}()
	wg.Wait()
	close(c)
	// time.Sleep(1 * 1e9)
	fmt.Println("channal len", len(c))
	// fmt.Println("array len", len(a))
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

package context

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// groutine exit seq

func m() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg, wgCtx := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		wg.Go(func() error {
			time.Sleep(time.Second)
			return nil
		})
	}
	wg.Wait()
	go f1(wgCtx) // 先执行
}

func f1(ctx context.Context) {
	// if ctx.Err() != nil {
	// 	return
	// }
	go f2(ctx) // 后执行
	<-ctx.Done()
	println("f1 done")

}

func f2(ctx context.Context) {
	<-ctx.Done()
	println("f2 done")
}

func f3() {
	b := time.AfterFunc(2*time.Second, func() {
		fmt.Println("hello")
	})
	a := <-time.After(1 * time.Second)
	println("1s pass")
	println(b.Reset(2 * time.Second))
	<-time.After(1 * time.Second)
	println("2s pass")
	<-time.After(1 * time.Second)
	println("3s pass")
	<-time.After(1 * time.Second)
	println("4s pass")
	// print(a)
	fmt.Println(a)
}

func f4() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func f5() {
	ch := f4()
	for i := range ch {
		fmt.Println(i)
	}
}

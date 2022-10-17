package channel

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)

// https://jishuin.proginn.com/p/763bfbd692b8
// 深入解析 Goroutine 泄露的场景：channel 发送者

type result struct {
	record string
	err    error
}

func search(term string) (string, error) {
	time.Sleep(1000 * time.Millisecond)
	return term, nil
}

// process 函数是一个用来寻找记录的函数
// 然后打印，如果超过 100 ms 就会失败 .
func process(term string) error {

	// 创建一个在 100 ms 内取消上下文的 context
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer func() {
		println("num of goroutine", runtime.NumGoroutine())
		cancel()
	}()

	// 为 Goroutine 创建一个传递结果的 channel
	// 阻塞式的 channel 与 非阻塞式的 channel 行为不同
	// ch := make(chan result, 1)
	ch := make(chan result)

	// 启动一个 Goroutine 来寻找记录，然后得到结果
	// 将返回值从 channel 中返回
	go func() {
		record, err := search(term)
		println("wait for search")
		ch <- result{record, err}
		println("search done")
	}()

	// 阻塞等待从 Goroutine 接收值
	// 通过 channel 和 context 来取消上下文操作
	var err error
	select {
	case <-ctx.Done():
		// time.Sleep(1200 * time.Millisecond)
		err = errors.New("search canceled")
	case result := <-ch:
		if result.err != nil {
			err = result.err
		}
		fmt.Println("Received:", result.record)
		// time.Sleep(1200 * time.Millisecond)
	}
	time.Sleep(5000 * time.Millisecond)
	return err

}

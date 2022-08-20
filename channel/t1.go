package channel

import "fmt"

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
}

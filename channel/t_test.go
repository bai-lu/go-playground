package channel

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

var NumCPU = runtime.GOMAXPROCS(2)
var NumCPU1 = runtime.NumCPU()

func TestNumCPU(t *testing.T) {
	var NumCPU = runtime.GOMAXPROCS(2)
	t.Log(NumCPU)
	t.Log(NumCPU1)

}

// 测试无缓冲 channal
func TestServer(t *testing.T) {
	ch := make(chan *Request, 100)
	quit := make(chan bool)
	go Serve(ch, quit)
	time.Sleep(time.Second)
	quit <- true
	time.Sleep(time.Second)
}

func TestSeq(t *testing.T) {
	// var ch1 chan string
	// var ch2 chan string
	// ch1 = make(chan string)
	// ch2 = make(chan string)
	a := ""

	go func() {
		a = "g1"

	}()
	go func() {
		a = "g2"
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}

func TestF1(t *testing.T) {
	F1()
}

func TestF2(t *testing.T) {
	F2()
}

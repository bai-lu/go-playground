package channel

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

var numCPU = runtime.GOMAXPROCS(0)
var numCPU1 = runtime.NumCPU()

func TestServer(t *testing.T) {
	t.Log(numCPU)
	t.Log(numCPU1)
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

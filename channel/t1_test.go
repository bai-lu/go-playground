package channel

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var NumCPU = runtime.GOMAXPROCS(2)
var NumCPU1 = runtime.NumCPU()

func TestNumCPU(t *testing.T) {
	var NumCPU = runtime.GOMAXPROCS(2)
	a := ""
	t.Log(NumCPU)
	t.Log(NumCPU1)
	t.Log(runtime.NumGoroutine())
	time.AfterFunc(2*time.Second, func() {
		t.Log(runtime.NumGoroutine())
		t.Log("hello", a)
	})
	time.AfterFunc(2*time.Second, func() {
		t.Log(runtime.NumGoroutine())
		t.Log("hello", a)
	})
	time.AfterFunc(2*time.Second, func() {
		t.Log(runtime.NumGoroutine())
		t.Log("hello", a)
	})
	time.Sleep(3 * time.Second)
	fmt.Println(a)
}

func TestSeq(t *testing.T) {
	var mutex sync.Mutex
	a := ""
	go func() {
		mutex.Lock()
		a = "g1"
		mutex.Unlock()
	}()
	go func() {
		mutex.Lock()
		a = "g2"
		mutex.Unlock()
	}()
	time.Sleep(1 * time.Second)
	mutex.Lock()
	fmt.Println(a)
	mutex.Unlock()
}

func TestF1(t *testing.T) {
	F1()
}

func TestF2(t *testing.T) {
	F2()
}

func TestT2(t *testing.T) {
	c := make(chan struct{}) // or buffered channel

	// The race detector cannot derive the happens before relation
	// for the following send and close operations. These two operations
	// are unsynchronized and happen concurrently.
	go func() { c <- struct{}{} }()
	<-c
	close(c)
}

package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func F1() {
	var lock sync.Mutex
	lock.Lock()
}

var Task = map[string]bool{
	"task1": true,
	"task2": false,
}

func TaskManager() {
	go func() {
		Task["task1"] = true
	}()
	// 后面的goruotine
	go func() {
		Task["task1"] = false
	}()

	time.Sleep(time.Second)
	fmt.Println(Task)

}

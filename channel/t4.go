package channel

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func makeCakeAndSend(cs chan string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		// time.Sleep(1 * 1e7)
		cs <- cakeName //send a strawberry cake
	}
}

func ReceiveCakeAndPack(a []string, cs chan string) {
	for s := range cs {
		fmt.Println("Packing received cake: ", s)
		a = append(a, s)
	}
	fmt.Println(a)
}

func t4() {
	// var lock sync.Mutex
	var a []*string
	cs := make(chan string)

	// go receiveCakeAndPack(a, cs)
	var wg errgroup.Group
	wg.Go(func() error {
		for s := range cs {
			s := s
			fmt.Println("Packing received cake: ", s)
			// lock.Lock()
			time.Sleep(time.Millisecond * 100)
			a = append(a, &s)
			// lock.Unlock()
		}
		fmt.Println("goroutine closed")
		return nil
	})
	var wg1 errgroup.Group
	wg1.Go(func() error {
		defer close(cs)
		makeCakeAndSend(cs, 5)
		return nil
	})
	wg1.Wait()
	wg.Wait()
	if s, ok := <-cs; !ok {
		fmt.Println("channel closed, s: ", s)
	}

	// time.Sleep(time.Second)
	// lock.Lock()
	// fmt.Printf("%v", a)
	for i, v := range a {
		fmt.Printf("%d, %v\n", i, *v)
	}
	// lock.Unlock()
}

func t6() {
	// var lock sync.Mutex
	var res []string
	cs := make(chan string)

	// go receiveCakeAndPack(a, cs)
	// var wg errgroup.Group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for s := range cs {
			fmt.Println("Packing received cake: ", s)
			// lock.Lock()
			res = append(res, s)
			// lock.Unlock()
		}
		fmt.Println("goroutine closed")
		wg.Done()
	}()

	var wg1 errgroup.Group
	wg1.Go(func() error {
		makeCakeAndSend(cs, 5)
		return nil
	})
	wg1.Wait()
	close(cs)
	wg.Wait()
	// time.Sleep(time.Second)
	// lock.Lock()
	fmt.Println(res)
	// lock.Unlock()
}

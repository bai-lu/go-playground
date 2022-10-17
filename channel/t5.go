package channel

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func makeCakeAndSend1(cs chan string, flavor string, count int) {
	defer println("makeCakeAndSend1 closed")
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		time.Sleep(500 * time.Millisecond)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack1(ctx context.Context, strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false

	for {
		//if both channels are closed then we can stop
		if strbry_closed && choco_closed {
			return
		}
		fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbry_ok := <-strbry_cs:
			if !strbry_ok {
				strbry_closed = true
				fmt.Println(" ... Strawberry channel closed!")
			} else {
				fmt.Println("Received from Strawberry channel.  Now packing", cakeName)
			}
		case cakeName, choco_ok := <-choco_cs:
			if !choco_ok {
				choco_closed = true
				fmt.Println(" ... Chocolate channel closed!")
			} else {
				fmt.Println("Received from Chocolate channel.  Now packing", cakeName)
			}
		case <-ctx.Done():
			fmt.Println(" ... Context cancelled!", ctx.Err())
			// close(strbry_cs)
			// close(choco_cs)
			return
		}
	}
}

func t5() {
	strbry_cs := make(chan string)
	choco_cs := make(chan string)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	//two cake makers
	go makeCakeAndSend1(choco_cs, "Chocolate", 10)   //make 3 chocolate cakes and send
	go makeCakeAndSend1(strbry_cs, "Strawberry", 10) //make 3 strawberry cakes and send

	//one cake receiver and packer
	go receiveCakeAndPack1(ctx, strbry_cs, choco_cs) //pack all cakes received on these cake channels

	//sleep for a while so that the program doesn’t exit immediately
	time.Sleep(3 * time.Second)
	println(runtime.NumGoroutine())

	// close(strbry_cs)
	// close(choco_cs)
	// time.Sleep(2 * time.Second)
}

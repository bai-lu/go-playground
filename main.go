package main

import (
	"fmt"
	"sync"
)

var num = make(chan int, 1)
var char = make(chan int, 1)
var wg sync.WaitGroup

func PrintNums() {
	defer wg.Done()
	for i := 0; i <= 12; i++ {
		// 使用双层循环实现一次性打印两个数字
		for j := 0; j < 2; j++ {
			fmt.Printf("%d", 2*i+j+1) // %d 格式化打印数字
		}
		// 1,执行一次内循环打印两个数字,num通道接收一个元素,阻塞,
		num <- 1
		// 4,再执行一次打印数字,char通道取走元素,继续执行
		<-char
	}
}

func PrintChars() {
	defer wg.Done()
	for i := 0; i <= 12; i++ {
		// 2,num通道取走元素,继续执行
		<-num
		// 使用双层循环实现一次性打印两个字符
		for j := 0; j < 2; j++ {
			fmt.Printf("%c", 'A'+(2*i+j)) // %c 格式化打印出字符(ACCIS码数字转字符)
		}
		// 3,执行一次内循环打印两个字符,char通道接收一个元素,阻塞
		char <- 1
	}
}
func main() {
	// 使用sync.WaitGroup实现主进程等待两个goroutine执行完再执行自己的程序
	wg.Add(2)
	go PrintNums()
	go PrintChars()
	wg.Wait()
	// 结果: 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ
}

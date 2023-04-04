package main

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	// 定义一个名为 `print` 的变参函数
	// var print func(string, ...interface{}) (n int, err error)
	// 将 fmt.Printf 函数赋值给变量 `print`
	print := fmt.Printf
	// 使用 `print` 打印字符串
	print("Hello, %s!\n", "World")
}

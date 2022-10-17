package base

import (
	"fmt"
	"unsafe"
)

func RuneA() {
	var a int = 1 // 默认64位 大部分情况
	var b int8 = 2
	var c int16 = 3
	var d int32 = 4
	var e int64 = 5
	var f uint64 = 100
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(e))
	fmt.Println(unsafe.Sizeof(f))

	var aa uint = 1

	aa = uint(a)
	m := aa << 60
	n := m << 1
	l := n << 1
	r := rune(l)
	fmt.Println(m, n, l, r)

	bb := *(*int)(unsafe.Pointer(&a))
	fmt.Println(bb)

}

func Sizeof() {
	var a int
	var b int8
	var c string = "go"
	var d string = "你好"
	// var e rune = "你好"
	var e = []rune("你好")
	var f = []int32("你好")
	println(unsafe.Sizeof(a))
	println(unsafe.Sizeof(b))
	println(unsafe.Sizeof(c))
	println(unsafe.Sizeof(d))
	println(unsafe.Sizeof(e))
	fmt.Println(f)
	l := unsafe.Sizeof(f)
	println(l)
}

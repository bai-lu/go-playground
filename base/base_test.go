package base

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	Rune()

}

func TestSizeof(t *testing.T) {
	Sizeof()

}

func TestTmp(t *testing.T) {
	fmt.Println(len("Go语音"))
	fmt.Println(len([]rune("Go语音")))
	var a int = 1
	b := a << 63
	fmt.Println(b)
	fmt.Printf("%T", b)

}

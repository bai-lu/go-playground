package base

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	RuneA()
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

func TestRuneA(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"TestRuneA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RuneA()
		})
	}
}

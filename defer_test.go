package main

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}

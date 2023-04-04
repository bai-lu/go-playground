package main

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	// We want to print numbers from 0 to 9. We use a for loop to get the numbers
	for i := 0; i < 10; i++ {
		// We use the defer keyword to print the numbers in reverse order
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Error:", err)
			}
		}()
		fmt.Println(i)
	}
}

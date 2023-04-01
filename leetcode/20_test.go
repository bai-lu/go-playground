package leetcode

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if s[0] == ']' || s[0] == '}' || s[0] == ')' {
		return false
	}
	stack := []rune{}
	top := rune(0)
	for _, i := range s {
		if i == '(' || i == '[' || i == '{' {
			stack = append(stack, i)
			top = stack[len(stack)-1]
		} else {
			if len(stack) == 0 {
				return false
			}
			if i == top+1 || i == top+2 {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					top = stack[len(stack)-1]
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	isValid("{{{[[]]}}}")
	for _, i := range "{[()]}" {
		fmt.Println(i)
	}
}

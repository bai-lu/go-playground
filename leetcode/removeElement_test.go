package leetcode

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}

func Test_removeElement(t *testing.T) {
	arr := []int{3, 2, 2, 3}
	removeElement(arr, 3)
	fmt.Println(arr)
}

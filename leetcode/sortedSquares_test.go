package leetcode

import (
	"container/list"
	"fmt"
	"testing"
)

func sortedSquares(nums []int) []int {
	result := []int{}
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[i]*nums[i] <= nums[j]*nums[j] {
			result = append(result, nums[j]*nums[j])
			j--
		} else {
			result = append(result, nums[i]*nums[i])
			i++
		}

	}

	sorted := []int{}
	for i := len(result) - 1; i >= 0; i-- {
		sorted = append(sorted, result[i])
	}
	return sorted

}

func Test_sortedSquares(t *testing.T) {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	arr := []int{1, 2, 3, 4, 5}
	arr1 := arr[0:1]
	fmt.Println(arr1)
}

func minSubArrayLen(target int, nums []int) int {
	left := 0
	min := len(nums) + 1
	sumArr := func(nums []int) int {
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
		}
		return sum
	}
	for i := 0; i <= len(nums); i++ {
		window := nums[left:i]
		for sumArr(window) >= target {

			if i-left < min {
				min = i - left
			}
			fmt.Println(min)
			left++
			window = nums[left:i]
		}

	}
	if min == len(nums)+1 {
		return 0
	}
	// fmt.Println(min)
	return min
}

func Test_minSubArrayLen(t *testing.T) {
	minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	list.New()
}

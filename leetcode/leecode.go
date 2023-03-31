package leetcode

import (
	"fmt"
	"sort"
)

func Check(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	// fmt.Println(nums)

	diff := nums[1] - nums[0] // 不等于就判断false
	i, j := 0, 1
	for j < len(nums) {
		if nums[j]-nums[i] != diff {
			return false
		}
		i += 1
		j += 1

	}
	return true
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	// 	nums := []int{4, 6, 5, 9, 3, 7}
	// l := []int{0, 0, 2}
	// r := []int{2, 3, 5}
	result := []bool{}
	for i := 0; i < len(l); i++ {

		subArray := nums[l[i] : r[i]+1]
		tmp := make([]int, len(subArray))
		copy(tmp, subArray)
		fmt.Println(tmp)
		result = append(result, Check(tmp))
	}
	return result

}

package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortInt(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
	// reverse
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}

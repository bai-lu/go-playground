package leetcode

import (
	"fmt"
	"testing"
)

func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t2.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func TestMergeTrees(t *testing.T) {
	var t1 = &TreeNode{Val: 1}
	var t2 = &TreeNode{Val: 2}
	var t3 = &TreeNode{Val: 3}
	var t4 = &TreeNode{Val: 4}
	var t5 = &TreeNode{Val: 5}
	var t6 = &TreeNode{Val: 6}
	var t7 = &TreeNode{Val: 7}

	t1.Left = t3
	t1.Right = t2
	t3.Left = t5
	t2.Right = t4
	t5.Left = t6
	t5.Right = t7

	t2 = &TreeNode{Val: 1}
	t3 = &TreeNode{Val: 2}
	t4 = &TreeNode{Val: 3}
	t5 = &TreeNode{Val: 4}
	t6 = &TreeNode{Val: 5}
	t7 = &TreeNode{Val: 6}

	t2.Left = t3
	t2.Right = t4
	t3.Left = t5
	t3.Right = t6
	t4.Right = t7

	var res = mergeTrees(t1, t2)
	fmt.Println(res)
}

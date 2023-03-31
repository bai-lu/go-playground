package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestCheck(t *testing.T) {

	nums := []int{5, 9, 3, 7}
	fmt.Println(Check(nums))

}

func Test_checkArithmeticSubarrays(t *testing.T) {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	fmt.Println(checkArithmeticSubarrays(nums, l, r))
}

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	n := 3
	nums2 := []int{2, 5, 6}
	m := 3
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums1)))
	fmt.Println(nums1)
}

func qs(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivot := nums[0]
	fmt.Println(pivot)
	subArray1 := []int{}
	subArray2 := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] <= pivot {
			subArray1 = append(subArray1, nums[i])
		} else {
			subArray2 = append(subArray2, nums[i])
		}
	}
	fmt.Println(subArray1, subArray2)
	// return append(subArray1, subArray2...)

	return append(append(qs(subArray1), pivot), qs(subArray2)...)
}

func Test_qs(t *testing.T) {
	nums := []int{5, 10, 1, 4, 9}
	// nums := []int{}
	fmt.Println(qs(nums))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 如果链表为空或链表只有一个节点，则链表已经反转
	if head == nil || head.Next == nil {
		return head
	}
	// 递归反转子链表
	newHead := reverseList(head.Next)
	// 连接当前节点和已反转的子链表
	head.Next.Next = head
	head.Next = nil
	fmt.Println("-")
	// 返回新的头节点
	return newHead
}

func TestReverseList(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	fmt.Println("Original List: ")
	printList(head)
	newHead := reverseList(head)
	fmt.Println("Reversed List: ")
	printList(newHead)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val, " ")
		head = head.Next
	}
}

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func Test_moveZeroes(t *testing.T) {
	nums := []int{1, 0, 0, 3, 12}
	moveZeroes(nums)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (result []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {

			stack = append(stack, root)
			result = append(result, root.Val)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return
}

func Test_preorderTraversal(t *testing.T) {
	// TreeNode* root = new TreeNode(1); TreeNode* node1 = new TreeNode(2); TreeNode* node2 = new TreeNode(3);
	// root->left = node1; root->right = node2;
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	root.Left = node1
	root.Right = node2
	result := preorderTraversal(root)
	fmt.Println(result)

}

// 递归过程
// maxDepth(3)
//     -> left = maxDepth(9)
//         -> return 1
//     -> right = maxDepth(20)
//         -> left = maxDepth(15)
//             -> return 1
//         -> right = maxDepth(7)
//             -> return 1
//         -> return max(1, 1) + 1 = 2
//     -> return max(1, 2) + 1 = 3

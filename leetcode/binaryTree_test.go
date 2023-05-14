package leetcode

import (
	"container/list"
	"fmt"
	"testing"
)

/*
*
102. 二叉树的递归遍历
*/
func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		// if the depth is equal to the number of rows, add a new row
		if len(arr) == depth {
			arr = append(arr, []int{})
		}

		// add the value to the correct row
		arr[depth] = append(arr[depth], root.Val)

		// recursively call the function on the left and right children
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, 0)

	return arr
}

/*
*
102. 二叉树的层序遍历
*/
func levelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val) //将值加入本层切片中
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []int{}          //清空层的数据
	}
	return res
}

func Test_levelOrder(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(levelOrder(root))

}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		tmp := arr[i]
		arr[i] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = tmp
	}
	fmt.Println(arr)
}

func Test_reverse(t *testing.T) {
	reverse([]int{1, 2, 3, 4, 5})
}

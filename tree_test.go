package main

import (
	"fmt"
	"strings"
	"testing"
)

// 定义一个哈希表型的文件树节点
type node map[string]node

// 创建文件树结构
func createTree(paths []string) node {
	// 创建一个根节点
	root := make(node)
	// 遍历所有文件路径，逐一插入节点
	for _, path := range paths {
		// 将文件路径拆分为各级目录，形成一个路径片段数组
		segments := strings.Split(path, "/")
		// 从根节点开始沿着路径构建子节点
		currNode := root
		for _, seg := range segments {
			// 如果当前节点不存在，就新建一个
			if _, ok := currNode[seg]; !ok {
				currNode[seg] = make(node)
			}
			// 进入子节点，进行下一轮循环
			currNode = currNode[seg]
		}
	}
	// 返回整个文件树结构
	return root
}

// 打印文件树结构
func printTree(tree node, indent int) {
	// 对节点进行深度优先遍历
	for name, node := range tree {
		// 根据节点的深度级别输出适当的缩进
		fmt.Printf("%s%s\n", strings.Repeat("  ", indent), name)
		// 如果当前节点有子节点，就递归打印它们
		if len(node) > 0 {
			printTree(node, indent+1)
		}
	}
}

func TestTree(t *testing.T) {
	// 准备一组文件路径
	paths := []string{
		"/usr/local/bin",
		"/usr/local/share/man",
		"/usr/local/include",
		"/usr/share/doc",
		"/usr/share/man",
		"/var/log",
	}
	// 构建文件树结构
	tree := createTree(paths)
	// 打印文件树结构
	printTree(tree, 0)
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
		1
		/\
	2		3
	/		/ \
4		   5	6
*/

// wrong
func inOrderWrong(root *TreeNode, array []int) {
	if root == nil {
		return
	}
	inOrderWrong(root.Left, array)
	array = append(array, root.Val)
	inOrderWrong(root.Right, array)
}

// array里那个指向数组的指针，也是以值的形式传进来的，也就是说，函数内部修改了这个指针（注意是修改了指针本身，而不是修改了指针指向的内容），外面是看不到的！！！
func inOrder(root *TreeNode, array *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, array)
	*array = append(*array, root.Val)
	inOrder(root.Right, array)
}

// PreOrderTraversal 前序遍历 1 -> 2 -> 4 -> 3 -> 5 -> 6 ->
func PreOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Printf("%d -> ", tree.Val)
	PreOrderTraversal(tree.Left)
	PreOrderTraversal(tree.Right)
}

// MidOrderTraversal 中序遍历 4 -> 2 -> 1 -> 5 -> 3 -> 6 ->
func MidOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}
	MidOrderTraversal(tree.Left)
	fmt.Printf(" %d -> ", tree.Val)
	MidOrderTraversal(tree.Right)
}

// PostOrderTraversal 后续遍历 4 -> 2 -> 5 -> 6 -> 3 -> 1 ->
func PostOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}
	PostOrderTraversal(tree.Left)
	PostOrderTraversal(tree.Right)
	fmt.Printf(" %d -> ", tree.Val)
}

// LevelOrderTraversal 按层遍历 1 -> 2 -> 3 -> 4 -> 5 -> 6 ->
func LevelOrderTraversal(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 采用队列实现
	queue := make([]*TreeNode, 0)
	queue = append(queue, tree)
	for len(queue) > 0 {
		tree = queue[0]
		fmt.Printf(" %d -> ", tree.Val)
		queue = queue[1:] // queue pop
		if tree.Left != nil {
			queue = append(queue, tree.Left)
		}
		if tree.Right != nil {
			queue = append(queue, tree.Right)
		}
	}
}

// 二叉树的中序遍历
// refer: https://blog.csdn.net/u013536232/article/details/105547626
func main() {
	var arr = []int{1, 3, 4, 2}
	var node TreeNode
	inOrderWrong(&node, arr)
	fmt.Println(node)
	fmt.Println(arr)
	inOrder(&node, &arr)
	fmt.Println(node)
	fmt.Println(arr)
}

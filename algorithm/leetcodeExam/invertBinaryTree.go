package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 226. 翻转二叉树：简单
// tags: 树、递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// tmp := root.Left
	// root.Left = root.Right
	// root.Right = tmp
	root.Left, root.Right = root.Right, root.Left // go语法交换元素，等价于上面3行代码
	root.Left = invertTree(root.Left)             // 递归左子树
	root.Right = invertTree(root.Right)           // 递归右子树
	return root
}

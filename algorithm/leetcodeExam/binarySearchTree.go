package main

// TreeNode is a binary tree's node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 700. 二叉搜索树中的搜索：简单
// tags: 二叉搜索树，二分查找

// 递归法
// 时间复杂度：O(N) 空间复杂度:O(N)
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val {
		// 去右子树搜索
		root = searchBST(root.Right, val)
	} else if root.Val > val {
		// 去左子树搜索
		root = searchBST(root.Left, val)
	}
	return root
}

// 迭代法
// 时间复杂度：O(N) 空间复杂度:O(1)
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val < val {
			root = root.Right
		} else if root.Val > val {
			root = root.Left
		}
	}
	return nil
}

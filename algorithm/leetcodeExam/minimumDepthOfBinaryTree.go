package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 111. 二叉树的最小深度：简单
// tags：二叉树，BFS，DFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	// root 本身就是一层， depth 初始化为 1
	depth := 1

	for len(q) != 0 {
		size := len(q)
		// 将当前队列中的所有节点向四周扩散
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			// 判断是否到达终点
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			// 将 cur 的相邻节点加入队列
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		// 每遍历一层，深度+1
		depth++
	}
	return depth
}

package structure

import "fmt"

// TreeNode 二叉树
type TreeNode struct {
	Data  int64
	Left  *TreeNode
	Right *TreeNode
}

// PreOrder 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印根节点
	fmt.Print(tree.Data, " ")
	// 再打印左子树
	PreOrder(tree.Left)
	// 再打印右字树
	PreOrder(tree.Right)
}

// MidOrder 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印左子树
	MidOrder(tree.Left)
	// 再打印根节点
	fmt.Print(tree.Data, " ")
	// 再打印右字树
	MidOrder(tree.Right)
}

// PostOrder 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印左子树
	MidOrder(tree.Left)
	// 再打印右字树
	MidOrder(tree.Right)
	// 再打印根节点
	fmt.Print(tree.Data, " ")
}

// Level 按层遍历
func Level(head *TreeNode) {

	if head == nil {
		return
	}

	// 用切片模仿队列
	var queue []*TreeNode
	queue = append(queue, head)

	for len(queue) != 0 {
		// 队头弹出，再把队头切掉，模仿队列的poll操作
		cur := queue[0]
		queue = queue[1:]

		fmt.Printf("%d", (*cur).Data)

		// 当前节点有左孩子，加入左孩子进队列
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		// 当前节点有右孩子，加入右孩子进队列
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}

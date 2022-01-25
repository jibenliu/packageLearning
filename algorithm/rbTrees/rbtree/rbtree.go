package rbtree

import "log"

// RBTree 实现红黑树
type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{root: nil}
}

func (t *RBTree) Insert(data int64) {
	if t.root == nil {
		rootNode := NewRBNode(data)
		rootNode.color = BLACK
		t.root = rootNode
	} else {
		t.insertNode(t.root, data)
	}
}

func (t *RBTree) insertNode(node *RBNode, data int64) {
	if node.value >= data {
		// 插入数据不大于父节点，插入左节点
		if node.left != nil {
			t.insertNode(node.left, data)
		} else {
			tmpNode := NewRBNode(data)
			tmpNode.parent = node
			node.left = tmpNode
			t.insertCheck(tmpNode)
		}
	} else {
		// 插入数据大于父节点，插入右节点
		if node.right != nil {
			t.insertNode(node.right, data)
		} else {
			tmpNode := NewRBNode(data)
			tmpNode.parent = node
			node.right = tmpNode
			t.insertCheck(tmpNode)
		}
	}
}

func (t *RBTree) rotateLeft(node *RBNode) {
	if tmpRoot, err := node.rotate(LeftRotate); err == nil {
		// 旋转时需验证是否有root的改动
		if tmpRoot != nil {
			t.root = tmpRoot
		}
	} else {
		log.Printf(err.Error())
	}
}

func (t *RBTree) rotateRight(node *RBNode) {
	if tmpRoot, err := node.rotate(RightRotate); err == nil {
		// 旋转时需验证是否有root的改动
		if tmpRoot != nil {
			t.root = tmpRoot
		}
	} else {
		log.Printf(err.Error())
	}
}

func (t *RBTree) insertCheck(node *RBNode) {
	if node.parent == nil {
		// 检查1：若插入节点没有父节点，则该节点为root
		t.root = node
		// 设置根节点为black
		t.root.color = BLACK
		return
	}

	// 父节点是黑色的话直接添加，红色则进行处理
	if node.parent.color == RED {
		if node.getUncle() != nil && node.getUncle().color == RED {
			// 父节点及叔父节点都是红色，则转为黑色
			node.parent.color = BLACK
			node.getUncle().color = BLACK
			// 祖父节点改成红色
			node.getGrandParent().color = RED
			// 递归处理
			t.insertCheck(node.getGrandParent())
		} else {
			// 父节点红色，父节点的兄弟节点不存在或为黑色
			isLeft := node == node.parent.left
			isParentLeft := node.parent == node.getGrandParent().left
			if !isLeft && isParentLeft {
				t.rotateLeft(node.parent)
				t.rotateRight(node.parent)

				node.color = BLACK
				node.left.color = RED
				node.right.color = RED
			} else if isLeft && !isParentLeft {
				t.rotateRight(node.parent)
				t.rotateLeft(node.parent)

				node.color = BLACK
				node.left.color = RED
				node.right.color = RED
			} else if isLeft && isParentLeft {
				node.parent.color = BLACK
				node.getGrandParent().color = RED
				t.rotateRight(node.getGrandParent())
			} else if !isLeft && !isParentLeft {
				node.parent.color = BLACK
				node.getGrandParent().color = RED
				t.rotateLeft(node.getGrandParent())
			}
		}
	}
}

func (t *RBTree) Delete(data int64) {
	t.deleteChild(t.root, data)
}

func (t *RBTree) deleteChild(n *RBNode, data int64) bool {
	if data < n.value {
		if n.left == nil {
			return false
		}
		return t.deleteChild(n.left, data)
	}
	if data > n.value {
		if n.right == nil {
			return false
		}
		return t.deleteChild(n.right, data)
	}

	if n.right == nil || n.left == nil {
		// 两个都为空或其中一个为空，转为删除一个子树节点的问题
		t.deleteOne(n)
		return true
	}

	//两个都不为空，转换成删除只含有一个子节点节点的问题
	mostLeftChild := n.right.getLeftMostChild()
	tmpVal := n.value
	n.value = mostLeftChild.value
	mostLeftChild.value = tmpVal

	t.deleteOne(mostLeftChild)

	return true
}

// 删除只有一个子节点的节点
func (t *RBTree) deleteOne(n *RBNode) {
	var child *RBNode
	isAdded := false
	if n.left == nil {
		child = n.right
	} else {
		child = n.left
	}

	if n.parent == nil && n.left == nil && n.right == nil {
		n = nil
		t.root = n
		return
	}
	if n.parent == nil {
		n = nil
		child.parent = nil
		t.root = child
		t.root.color = BLACK
		return
	}

	if n.color == RED {
		if n == n.parent.left {
			n.parent.left = child

		} else {
			n.parent.right = child
		}
		if child != nil {
			child.parent = n.parent
		}
		n = nil
		return
	}

	if child != nil && child.color == RED && n.color == BLACK {
		if n == n.parent.left {
			n.parent.left = child

		} else {
			n.parent.right = child
		}
		child.parent = n.parent
		child.color = BLACK
		n = nil
		return
	}

	// 如果没有孩子节点，则添加一个临时孩子节点
	if child == nil {
		child = NewRBNode(0)
		child.parent = n
		isAdded = true
	}

	if n.parent.left == n {
		n.parent.left = child
	} else {
		n.parent.right = child
	}

	child.parent = n.parent

	if n.color == BLACK {
		if !isAdded && child.color == RED {
			child.color = BLACK
		} else {
			t.deleteCheck(child)
		}
	}

	// 如果孩子节点是后来加的，需删除
	if isAdded {
		if child.parent.left == child {
			child.parent.left = nil
		} else {
			child.parent.right = nil
		}
		child = nil
	}
	n = nil
}

func (t *RBTree) deleteCheck(n *RBNode) {
	if n.parent == nil {
		n.color = BLACK
		return
	}
	if n.getSibling().color == RED {
		n.parent.color = RED
		n.getSibling().color = BLACK
		if n == n.parent.left {
			t.rotateLeft(n.parent)
		} else {
			t.rotateRight(n.parent)
		}
	}
	//注意：这里n的兄弟节点发生了变化，不再是原来的兄弟节点
	isParentRed := n.parent.color
	isSibRed := n.getSibling().color
	isSibLeftRed := BLACK
	isSibRightRed := BLACK
	if n.getSibling().left != nil {
		isSibLeftRed = n.getSibling().left.color
	}
	if n.getSibling().right != nil {
		isSibRightRed = n.getSibling().right.color
	}

	if !isParentRed && !isSibRed && !isSibLeftRed && !isSibRightRed {
		n.getSibling().color = RED
		t.deleteCheck(n.parent)
		return
	}
	if isParentRed && !isSibRed && !isSibLeftRed && !isSibRightRed {
		n.getSibling().color = RED
		n.parent.color = BLACK
		return
	}
	if n.getSibling().color == BLACK {
		if n.parent.left == n && isSibLeftRed && !isSibRightRed {
			n.getSibling().color = RED
			n.getSibling().left.color = BLACK
			t.rotateRight(n.getSibling())
		} else if n.parent.right == n && !isSibLeftRed && isSibRightRed {
			n.getSibling().color = RED
			n.getSibling().right.color = BLACK
			t.rotateLeft(n.getSibling())
		}
	}
	n.getSibling().color = n.parent.color
	n.parent.color = BLACK
	if n.parent.left == n {
		n.getSibling().right.color = BLACK
		t.rotateLeft(n.parent)
	} else {
		n.getSibling().left.color = BLACK
		t.rotateRight(n.parent)
	}
}

// log输出树
func printTreeInLog(n *RBNode, front string) {
	if n != nil {
		var colorStr string
		if n.color == RED {
			colorStr = "红"
		} else {
			colorStr = "黑"
		}
		log.Printf(front+"%d,%s\n", n.value, colorStr)
		printTreeInLog(n.left, front+"-(l)|")

		printTreeInLog(n.right, front+"-(r)|")
	}
}

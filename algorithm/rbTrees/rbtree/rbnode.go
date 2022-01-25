package rbtree

import "errors"

const (
	// RED 红树设为true
	RED bool = true
	// BLACK 黑树设为false
	BLACK bool = false
)

const (
	// LeftRotate 左旋
	LeftRotate bool = true
	// RightRotate 右旋
	RightRotate bool = false
)

// RBNode 红黑树
type RBNode struct {
	value               int64
	color               bool
	left, right, parent *RBNode
}

// NewRBNode 初始化红黑树
func NewRBNode(value int64) *RBNode {
	return &RBNode{color: RED, value: value}
}

// getGrandParent() 获取父级节点的父级节点
func (rb *RBNode) getGrandParent() *RBNode {
	if rb.parent == nil {
		return nil
	}
	return rb.parent.parent
}

// getSibling() 获取兄弟节点
func (rb *RBNode) getSibling() *RBNode {
	if rb.parent == nil {
		return nil
	}
	if rb == rb.parent.left {
		return rb.parent.right
	} else {
		return rb.parent.left
	}
}

// GetUncle() 父节点的兄弟节点
func (rb *RBNode) getUncle() *RBNode {
	if rb.getGrandParent() == nil {
		return nil
	}
	if rb.parent == rb.getGrandParent().right {
		return rb.getGrandParent().left
	} else {
		return rb.getGrandParent().right
	}
}

// rotate() 左旋/右旋
// 若有根节点变动则返回根节点
func (rb *RBNode) rotate(isRotateLeft bool) (*RBNode, error) {
	var root *RBNode
	if rb == nil {
		return root, nil
	}
	if !isRotateLeft && rb.left == nil {
		return root, errors.New("右旋左节点不能为空")
	} else if isRotateLeft && rb.right == nil {
		return root, errors.New("左旋右节点不能为空")
	}

	parent := rb.parent
	var isLeft bool
	if parent != nil {
		isLeft = parent.left == rb
	}

	if isRotateLeft {
		grandson := rb.right.left
		rb.right.left = rb
		rb.parent = rb.right
		rb.right = grandson
	} else {
		grandson := rb.left.right
		rb.left.right = rb
		rb.parent = rb.left
		rb.left = grandson
	}
	if parent == nil {
		rb.parent.parent = nil
		root = rb.parent

	} else {
		if isLeft {
			parent.left = rb.parent
		} else {
			parent.right = rb.parent
		}
		rb.parent.parent = parent
	}
	return root, nil
}

// 获取某节点最左侧的叶子，删除有2个孩子的节点时用
func (rb *RBNode) getLeftMostChild() *RBNode {
	if rb.left == nil {
		return rb
	}
	return rb.left.getLeftMostChild()
}

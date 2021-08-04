package main

// 判断链表是否是回文链表
// O(n) 时间复杂度和 O(1) 空间复杂度


//思路：
//定义快慢双指针
//关键点：找出mid，快指针前进2步，慢指针前进1步，最后循环结束慢指针为中间值
//以mid为head进行后半段的反转
//前半段与后半段的Val值依次比较，只要有对不上的立刻返回

type ListNode struct {
	Val  int
	Cur  *ListNode
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	p1, p2 := head, secondHalfStart
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 反转恢复链表，写不写都可
	// firstHalfEnd.Next = reverseList(secondHalfStart)
	return true
}

// 找出中间节点
func endOfFirstHalf(head *ListNode) *ListNode {
	slow, first := head, head
	for first.Next != nil && first.Next.Next != nil {
		slow = slow.Next
		first = first.Next.Next
	}
	return slow
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		tmpNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmpNext
	}
	return pre
}

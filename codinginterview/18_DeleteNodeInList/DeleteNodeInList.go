package deletenodeinlist

// 面试题18（一）：在O(1)时间删除链表结点
// 题目：给定单向链表的头指针和一个结点指针，定义一个函数在O(1)时间删除该
// 结点。
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Next == nil {
		return head
	} else if head.Val == val {
		return head.Next
	}
	pre := head
	next := head.Next
	for next != nil {
		if next.Val == val {
			pre.Next = next.Next
			break
		}
		pre = next
		next = next.Next
	}
	return head
}

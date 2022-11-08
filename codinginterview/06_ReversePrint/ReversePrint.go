package ReversePrint

type ListNode struct {
	Val  int
	Next *ListNode
}

// 面试题6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。

// 递归:
func reversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return reverseList1(head)
}

func reverseList1(node *ListNode) []int {
	if node.Next != nil {
		nums := reverseList1(node.Next)
		nums = append(nums, node.Val)
		return nums
	}
	return []int{node.Val}
}

// 递归优化,内存
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return reverseList2(head, 0)
}

func reverseList2(node *ListNode, depth int) []int {
	if node.Next != nil {
		list := reverseList2(node.Next, depth+1)
		list[len(list)-depth-1] = node.Val
		return list
	}
	list := make([]int, depth+1)
	list[0] = node.Val
	return list
}

// 反转链表
func reversePrint3(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var node *ListNode
	var newHead *ListNode
	for head != nil {
		node = head.Next
		head.Next = newHead
		newHead = head
		head = node
	}

	var list []int
	for newHead != nil {
		list = append(list, newHead.Val)
		newHead = newHead.Next
	}
	return list
}

// 直接遍历
func reversePrint4(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	for i, j := 0, len(list)-1; i < j; {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list
}

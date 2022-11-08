package NextNodeInBinaryTrees

import "offer/utils"

// 面试题8：二叉树的下一个结点
// 题目：给定一棵二叉树和其中的一个结点，如何找出中序遍历(左根右)顺序的下一个结点？
// 树中的结点除了有两个分别指向左右子结点的指针以外，还有一个指向父结点的指针。

func getNext(node *utils.TreeNode) *utils.TreeNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	if node.Parent != nil {
		current := node
		parent := current.Parent
		for parent != nil && current == parent.Right {
			current = parent
			parent = current.Parent
		}
		return parent
	}
	return nil
}

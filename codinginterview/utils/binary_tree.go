package utils

import "fmt"

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{
		Val: value,
	}
}

func ConnectTreeNodes(root, left, right *TreeNode) {
	if root == nil {
		return
	}
	root.Left = left
	if left != nil {
		left.Parent = root
	}
	root.Right = right
	if right != nil {
		right.Parent = root
	}
}

func PrintTreeNode(node *TreeNode) {
	if node == nil {
		fmt.Println("this node is null.")
		return
	}
	fmt.Printf("%v \n", node.Val)
}

func PrintTree(root *TreeNode) {
	PrintTreeNode(root)
	if root.Left != nil {
		PrintTree(root.Left)
	}
	if root.Right != nil {
		PrintTree(root.Right)
	}
}

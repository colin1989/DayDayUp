package BuildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题目：输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输
// 入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,
// 2, 4, 7, 3, 5, 6, 8}和中序遍历序列{4, 7, 2, 1, 5, 3, 8, 6}，则重建出
// 图2.6所示的二叉树并输出它的头结点。
func buildTreePreorderInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	return subTree1(preorder, inorder)
}

func subTree1(preorder []int, inorder []int) *TreeNode {
	// 查找中序遍历的根节点
	middle := 0
	for ; middle < len(inorder); middle++ {
		if preorder[0] == inorder[middle] {
			break
		}
	}
	if middle >= len(preorder) || middle >= len(inorder) {
		return nil
	}
	// 前序遍历,第一个数是根节点
	return &TreeNode{
		Val:   preorder[0],
		Left:  subTree1(preorder[1:middle+1], inorder[0:middle]),
		Right: subTree1(preorder[middle+1:], inorder[middle+1:]),
	}
}

func buildTreePreorderPostorder(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	return subTree2(inorder, postorder)
}

func subTree2(inorder, postorder []int) *TreeNode {
	// 查找中序遍历的根节点
	middle := 0
	for ; middle < len(inorder); middle++ {
		if postorder[len(postorder)-1] == inorder[middle] {
			break
		}
	}
	if middle >= len(postorder) || middle >= len(inorder) {
		return nil
	}
	// 前序遍历,第一个数是根节点
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  subTree2(inorder[0:middle], postorder[0:middle]),
		Right: subTree2(inorder[middle+1:], postorder[middle:len(postorder)-1]),
	}
}

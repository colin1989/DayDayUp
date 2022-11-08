package NextNodeInBinaryTrees

import (
	"offer/utils"
	"reflect"
	"testing"
)

//               8
//         4              12
//       2    6        10    13
//      1 3  5  7    9   11
func Test_getNextInorder(t *testing.T) {
	node1 := utils.CreateTreeNode(1)
	node2 := utils.CreateTreeNode(2)
	node3 := utils.CreateTreeNode(3)
	node4 := utils.CreateTreeNode(4)
	node8 := utils.CreateTreeNode(8)
	node6 := utils.CreateTreeNode(6)
	node10 := utils.CreateTreeNode(10)
	node5 := utils.CreateTreeNode(5)
	node7 := utils.CreateTreeNode(7)
	node9 := utils.CreateTreeNode(9)
	node11 := utils.CreateTreeNode(11)
	node12 := utils.CreateTreeNode(12)
	node13 := utils.CreateTreeNode(13)

	utils.ConnectTreeNodes(node8, node4, node12)
	utils.ConnectTreeNodes(node4, node2, node6)
	utils.ConnectTreeNodes(node2, node1, node3)
	utils.ConnectTreeNodes(node6, node5, node7)
	utils.ConnectTreeNodes(node12, node10, node13)
	utils.ConnectTreeNodes(node10, node9, node11)

	tests := []struct {
		name string
		node *utils.TreeNode
		want *utils.TreeNode
	}{
		// TODO: Add test cases.
		{"8-9", node8, node9},
		{"4-5", node4, node5},
		{"11-12", node11, node12},
		{"5-6", node5, node6},
		{"10-11", node10, node11},
		{"7-8", node7, node8},
		{"9-10", node9, node10},
		{"13", node13, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNext(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNextInorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

package BuildTree

import (
	"testing"
)

func Test_buildTreePreorderInorder(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		//              1
		//           /     \
		//          2       3
		//         /       / \
		//        4       5   6
		//         \         /
		//          7       8
		{"普通二叉树", args{[]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTreePreorderInorder(tt.args.preorder, tt.args.inorder)
			if got == nil {
				t.Errorf("buildTreePreorderInorder() = %v", got)
			}
			t.Logf("%+v", got)
		})
	}
}

func Test_buildTreePreorderPostorder(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		//              1
		//           /     \
		//          2       3
		//         /       / \
		//        4       5   6
		//         \         /
		//          7       8
		{"普通二叉树", args{[]int{4, 2, 7, 5, 8, 1, 3, 6}, []int{4, 7, 8, 5, 2, 6, 3, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTreePreorderPostorder(tt.args.inorder, tt.args.postorder)
			if got == nil {
				t.Errorf("buildTreePreorderInorder() = %v", got)
			}
			t.Logf("%+v", got)
		})
	}
}

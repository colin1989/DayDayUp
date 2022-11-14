package deletenodeinlist

import (
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{"123", args{head: &ListNode{-3,
			&ListNode{5,
				&ListNode{-99, nil}}}, val: -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNode(tt.args.head, tt.args.val)
			for got != nil {
				t.Logf("%v/n", got.Val)
				got = got.Next
			}
		})
	}
}

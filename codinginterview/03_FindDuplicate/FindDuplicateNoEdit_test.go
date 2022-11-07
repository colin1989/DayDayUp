package FindDuplicate

import (
	"sort"
	"testing"
)

func Test_findDuplicateNoEdit(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"多个重复的数字", args{[]int{2, 3, 5, 4, 3, 2, 6, 7}}, []int{2, 3}},
		{"一个重复的数字", args{[]int{3, 2, 1, 4, 4, 5, 6, 7}}, []int{4}},
		{"重复的数字是数组中最小的数字", args{[]int{1, 2, 3, 4, 5, 6, 7, 1, 8}}, []int{1}},
		{"重复的数字是数组中最大的数字", args{[]int{1, 7, 3, 4, 5, 6, 8, 2, 8}}, []int{8}},
		{"数组中只有两个数字", args{[]int{1, 1}}, []int{1}},
		{"重复的数字位于数组当中", args{[]int{3, 2, 1, 3, 4, 5, 6, 7}}, []int{3}},
		{"多个重复的数字", args{[]int{1, 2, 2, 6, 4, 5, 6, 4}}, []int{1, 2, 4, 6}},
		{"一个数字重复三次", args{[]int{1, 2, 2, 6, 4, 5, 2}}, []int{2}},
		{"没有重复的数字", args{[]int{1, 2, 6, 4, 5, 3}}, []int{-1}},
		{"无效的输入", args{[]int{}}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicateNoEdit(tt.args.nums); sort.IntSlice(tt.want).Search(got) == -1 {
				t.Errorf("findDuplicateNoEdit() = %v, want %v", got, tt.want)
			}
		})
	}
}

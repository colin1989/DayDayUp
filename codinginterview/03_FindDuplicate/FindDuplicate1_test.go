package FindDuplicate

import (
	"sort"
	"testing"
)

func Test_findDuplicate1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"重复的数字是数组中最小的数字", args{[]int{2, 1, 3, 1, 4}}, []int{1}},
		{"重复的数字是数组中最大的数字", args{[]int{2, 4, 3, 1, 4}}, []int{4}},
		{"数组中存在多个重复的数字", args{[]int{2, 4, 2, 1, 4}}, []int{4}},
		{"没有重复的数字", args{[]int{2, 1, 3, 0, 4}}, []int{-1}},
		{"没有重复的数字", args{[]int{2, 1, 3, 5, 4, 0}}, []int{-1}},
		{"无效的输入", args{[]int{}}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate1(tt.args.nums); sort.IntSlice(tt.want).Search(got) == -1 {
				t.Errorf("findDuplicate1() = %v, want %v", got, tt.want)
			}
		})
	}
}

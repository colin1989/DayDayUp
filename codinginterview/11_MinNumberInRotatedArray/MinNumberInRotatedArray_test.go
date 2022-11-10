package MinNumberInRotatedArray

import "testing"

func Test_minArray(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		// TODO: Add test cases.
		{"典型输入，单调升序的数组的一个旋转", []int{3, 4, 5, 1, 2}, 1},
		{"单调升序数组，旋转0个元素，也就是单调升序数组本身", []int{1, 2, 3, 4, 5}, 1},
		{"有重复的数字，并且重复的数字刚好是第一个数字和最后一个数字", []int{1, 0, 1, 1, 1}, 0},
		{"数组中只有一个数字", []int{2}, 2},
		{"leetcode[22201]", []int{2, 2, 2, 0, 1}, 0},
		{"leetcode[311]", []int{3, 1, 1}, 1},
		{"leetcode[313]", []int{3, 1, 3}, 1},
		{"leetcode[3313]", []int{3, 3, 1, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minArrayInRotatedArray(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		// TODO: Add test cases.
		{"典型输入，单调升序的数组的一个旋转", []int{3, 4, 5, 1, 2}, 1},
		{"单调升序数组，旋转0个元素，也就是单调升序数组本身", []int{1, 2, 3, 4, 5}, 1},
		{"有重复的数字，并且重复的数字刚好是第一个数字和最后一个数字", []int{1, 0, 1, 1, 1}, 0},
		{"数组中只有一个数字", []int{2}, 2},
		{"leetcode[22201]", []int{2, 2, 2, 0, 1}, 0},
		{"leetcode[311]", []int{3, 1, 1}, 1},
		{"leetcode[313]", []int{3, 1, 3}, 1},
		{"leetcode[3313]", []int{3, 3, 1, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArrayInRotatedArray(tt.numbers); got != tt.want {
				t.Errorf("minArrayInRotatedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

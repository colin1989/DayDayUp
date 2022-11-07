package FindNumberIn2DArray

import "testing"

func Test_findNumberIn2DArray(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"要查找的数在数组中1", args{[][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}, 7}, true},
		{"要查找的数在数组中2", args{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5}, true},
		{"要查找的数不在数组中", args{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20}, false},
		{"空的二维数组", args{[][]int{{}}, 1}, false},
		{"单个存在", args{[][]int{{-5}}, -5}, true},
		{"单个不存在", args{[][]int{{-5}}, 1}, false},
		{"单行存在", args{[][]int{{1, 1}}, 1}, true},
		{"单行不存在", args{[][]int{{1, 1}}, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArray(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

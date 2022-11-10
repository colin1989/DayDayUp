package RobotMove

import "testing"

func Test_movingCount(t *testing.T) {
	type args struct {
		threshold int
		rows      int
		cols      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"10行10列", args{threshold: 5, rows: 10, cols: 10}, 21},
		{"20行20列", args{threshold: 15, rows: 20, cols: 20}, 359},
		{"1行100列", args{threshold: 10, rows: 1, cols: 100}, 29},
		{"1行10列", args{threshold: 10, rows: 1, cols: 10}, 10},
		{"100行1列", args{threshold: 15, rows: 100, cols: 1}, 79},
		{"10行1列", args{threshold: 15, rows: 10, cols: 1}, 10},
		{"1行1列", args{threshold: 15, rows: 1, cols: 1}, 1},
		{"1行1列", args{threshold: 0, rows: 1, cols: 1}, 1},
		{"机器人不能进入任意一个方格", args{threshold: -10, rows: 10, cols: 10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.threshold, tt.args.rows, tt.args.cols); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

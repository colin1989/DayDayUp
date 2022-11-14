package Power

import (
	"testing"
)

func Test_power(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"2/3/8", args{2, 3}, 8},
		{"-2/3/-8", args{-2, 3}, -8},
		{"2/-3/0.125", args{2, -3}, 0.125},
		{"2/0/1", args{2, 0}, 1},
		{"-2/0/1", args{-2, 0}, 1},
		{"0/0/1", args{0, 0}, 1},
		{"0/4/0", args{0, 4}, 0},
		{"0/-4/0", args{0, -4}, 0},
		{"0.00001/2147483647/1", args{0.00001, 2147483647}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := power1(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("power() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myPower(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"2/3/8", args{2, 3}, 8},
		{"-2/3/-8", args{-2, 3}, -8},
		{"2/-3/0.125", args{2, -3}, 0.125},
		{"2/0/1", args{2, 0}, 1},
		{"-2/0/1", args{-2, 0}, 1},
		{"0/0/1", args{0, 0}, 1},
		{"0/4/0", args{0, 4}, 0},
		{"0/-4/0", args{0, -4}, 0},
		{"0.00001/2147483647/1", args{0.00001, 2147483647}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPower(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

package numberof1inbinary

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"11", args{11}, 3},
		{"5", args{5}, 2},
		{"0x7FFFFFFF", args{0x7FFFFFFF}, 31},
		{"0xFFFFFFFF", args{0xFFFFFFFF}, 32},
		{"0x80000000", args{0x80000000}, 1},
		{"-1", args{-1}, 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeight1(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"11", args{11}, 3},
		{"5", args{5}, 2},
		{"0x7FFFFFFF", args{0x7FFFFFFF}, 31},
		{"0xFFFFFFFF", args{0xFFFFFFFF}, 32},
		{"0x80000000", args{0x80000000}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight1(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight1() = %v, want %v", got, tt.want)
			}
		})
	}
}

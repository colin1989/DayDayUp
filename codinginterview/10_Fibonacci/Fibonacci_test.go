package Fibonacci

import "testing"

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	// TODO: Add test cases.
	{"0", args{0}, 0},
	{"1", args{1}, 1},
	{"2", args{2}, 1},
	{"3", args{3}, 2},
	{"4", args{4}, 3},
	{"5", args{5}, 5},
	{"6", args{6}, 8},
	{"7", args{7}, 13},
	{"8", args{8}, 21},
	{"9", args{9}, 34},
	{"10", args{10}, 55},
	{"11", args{11}, 89},
	{"12", args{12}, 144},
	{"13", args{13}, 233},
	{"14", args{14}, 377},
	{"43", args{43}, 433494437},
	{"44", args{44}, 701408733},
	{"45", args{45}, 134903163},
	{"95", args{95}, 407059028},
}

func Test_fibonacci1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci1(tt.args.n); got != tt.want {
				t.Errorf("fibonacci1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacci2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci2(tt.args.n); got != tt.want {
				t.Errorf("fibonacci2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacci3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci3(tt.args.n); got != tt.want {
				t.Errorf("fibonacci3() = %v, want %v", got, tt.want)
			}
		})
	}
}

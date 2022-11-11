package CuttingRope

import (
	"testing"
)

var tests = []struct {
	name string
	n    int
	want int
}{
	// TODO: Add test cases.
	{"1", 1, 0},
	{"2", 2, 1},
	{"3", 3, 2},
	{"4", 4, 4},
	{"5", 5, 6},
	{"6", 6, 9},
	{"7", 7, 12},
	{"8", 8, 18},
	{"9", 9, 27},
	{"10", 10, 36},
	{"50", 50, 86093442},
	//{"90", 90, 205891132094649},
	//{"100", 100, 7412080755407364},
	//{"105", 105, 50031545098999707},
	//{"110", 110, 300189270593998242},
	{"120", 120, 953271190},
	//{"120", 120, 8105110306037952534},
}

func Test_cuttingRopeDP(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRopeDP(tt.n); got != tt.want {
				t.Errorf("cuttingRopeDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cuttingRopeGA(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRopeGA(tt.n); got != tt.want {
				t.Errorf("cuttingRopeGA() = %v, want %v", got, tt.want)
			}
		})
	}
}

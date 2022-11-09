package algorithmclassic

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"12345678", args{[]int{4, 3, 5, 8, 1, 6, 2, 7}, 0, 7}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.nums, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort([]int{4, 3, 5, 8, 1, 6, 2, 7}, 0, 7)
	}
}

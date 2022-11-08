package ReversePrint

import (
	"reflect"
	"testing"
)

func Test_reversePrint1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"123", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}}}, []int{3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"12345", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}}}, []int{5, 4, 3, 2, 1}},
		{"123456", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							&ListNode{6, nil}}}}}}}, []int{6, 5, 4, 3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"nil", args{nil}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversePrint2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{"123", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}}}, []int{3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"12345", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}}}, []int{5, 4, 3, 2, 1}},
		{"123456", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							&ListNode{6, nil}}}}}}}, []int{6, 5, 4, 3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"nil", args{nil}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversePrint3(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{"123", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}}}, []int{3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"12345", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}}}, []int{5, 4, 3, 2, 1}},
		{"123456", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							&ListNode{6, nil}}}}}}}, []int{6, 5, 4, 3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"nil", args{nil}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint3(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversePrint4(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{"123", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}}}, []int{3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"12345", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}}}, []int{5, 4, 3, 2, 1}},
		{"123456", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							&ListNode{6, nil}}}}}}}, []int{6, 5, 4, 3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"nil", args{nil}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint4(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReversePrint1(b *testing.B) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"123", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}}}, []int{3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"12345", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}}}, []int{5, 4, 3, 2, 1}},
		{"123456", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							&ListNode{6, nil}}}}}}}, []int{6, 5, 4, 3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"nil", args{nil}, nil},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reversePrint1(tt.args.head)
		}
	}
}

func BenchmarkReversePrint2(b *testing.B) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"123", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}}}, []int{3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"12345", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}}}, []int{5, 4, 3, 2, 1}},
		{"123456", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							&ListNode{6, nil}}}}}}}, []int{6, 5, 4, 3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"nil", args{nil}, nil},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reversePrint2(tt.args.head)
		}
	}
}

func BenchmarkReversePrint3(b *testing.B) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"123", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}}}, []int{3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"12345", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}}}, []int{5, 4, 3, 2, 1}},
		{"123456", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							&ListNode{6, nil}}}}}}}, []int{6, 5, 4, 3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"nil", args{nil}, nil},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reversePrint3(tt.args.head)
		}
	}
}

func BenchmarkReversePrint4(b *testing.B) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"123", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3, nil}}}}, []int{3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"12345", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5, nil}}}}}}, []int{5, 4, 3, 2, 1}},
		{"123456", args{head: &ListNode{1,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{5,
							&ListNode{6, nil}}}}}}}, []int{6, 5, 4, 3, 2, 1}},
		{"1", args{head: &ListNode{1, nil}}, []int{1}},
		{"nil", args{nil}, nil},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			reversePrint4(tt.args.head)
		}
	}
}

package QueueWithTwoStacks

import (
	"reflect"
	"testing"
)

func TestCQueue_AppendTail(t *testing.T) {
	type fields struct {
		inStack  []int
		outStack []int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CQueue{
				inStack:  tt.fields.inStack,
				outStack: tt.fields.outStack,
			}
			c.AppendTail(tt.args.value)
		})
	}
}

func TestCQueue_DeleteHead(t *testing.T) {
	type fields struct {
		inStack  []int
		outStack []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{"3", fields{[]int{3}, []int{}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CQueue{
				inStack:  tt.fields.inStack,
				outStack: tt.fields.outStack,
			}
			if got := c.DeleteHead(); got != tt.want {
				t.Errorf("DeleteHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want CQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

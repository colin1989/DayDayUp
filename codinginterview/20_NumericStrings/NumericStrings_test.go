package NumericStrings

import (
	"math/rand"
	"testing"
)

func Test_isNumber(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		// TODO: Add test cases.
		{" ", " ", false},
		{"1  ", "1  ", true},
		{" 0", " 0", true},
		{"3.", "3.", true},
		{" .", " .", false},
		{". ", ". ", false},
		{"e", "e", false},
		{".", ".", false},
		{".1", ".1", true},
		{"2e0", "2e0", true},
		{".e1", ".e1", false},
		{"-1E-16", "-1E-16", true},
		{"D+ ", "D+ ", false},
		{".8+", ".8+", false},
		{"4e+", "4e+", false},
		{".-4", ".-4", false},
		{".-", ".-", false},
		{"-.", "-.", false},
		{"-.-", "-.-", false},
		{"-.8", "-.8", true},
		{"+.8", "+.8", true},
		{"1.1e+10", "1.1e+10", true},
		{"0e", "0e", false},
		{"..2", "..2", false},
		{"+E3", "+E3", false},
		{"-.3e6", "-.3e6", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args); got != tt.want {
				t.Errorf("isNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkShuffle(b *testing.B) {
	offset := 1
	for i := 0; i < b.N; i++ {
		s := []int{0, 1, 2, 3, 4}
		rand.Shuffle(len(s)-offset, func(i, j int) {
			s[i+offset], s[j+offset] = s[j+offset], s[i+offset]
		})
	}
}

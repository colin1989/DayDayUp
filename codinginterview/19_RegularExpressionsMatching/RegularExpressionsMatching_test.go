package regularexpressionsmatching

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"空", args{}, true},
		{"正则 .*", args{p: ".*"}, true},
		{"正则 .", args{p: "."}, false},
		{"正则 c*", args{p: "c*"}, true},
		{"a / .*", args{s: "a", p: ".*"}, true},
		{"a / a.", args{s: "a", p: "a."}, false},
		{"ab / .*c", args{s: "ab", p: ".*c"}, false},
		{"aaa / .a*", args{s: "aaa", p: ".a*"}, true},
		{" / c*c*", args{s: "", p: "c*c*"}, true},
		{"aaa / ab*.*", args{s: "aaa", p: "ab*.*"}, true},
		{"aa / a*", args{s: "aa", p: "a*"}, true},
		{"aa / a", args{s: "aa", p: "a"}, false},
		{"aa / a*", args{s: "aaaab", p: "ac*b*"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

type ts struct {
	data [1000]int
}

var (
	ts1  ts
	ts2  ts
	ts3  ts
	ts4  ts
	ts5  ts
	ts6  ts
	ts7  ts
	ts8  ts
	ts9  ts
	ts10 ts
)

func BenchmarkCleanSlice1(b *testing.B) {
	s := []*ts{&ts1, &ts2, &ts3, &ts4, &ts5, &ts6, &ts7, &ts8, &ts9, &ts10}
	for i := 0; i < b.N; i++ {
		s = s[:0]
		s = append(s, []*ts{&ts1, &ts2, &ts3, &ts4, &ts5}...)
		s = append(s, []*ts{&ts6, &ts7, &ts8, &ts9, &ts10}...)
	}
}

func BenchmarkCleanSlice2(b *testing.B) {
	s := []*ts{&ts1, &ts2, &ts3, &ts4, &ts5, &ts6, &ts7, &ts8, &ts9, &ts10}
	for i := 0; i < b.N; i++ {
		for j := 0; j < 5; j++ {
			s = append(s[:j], s[j+1:]...)
		}
		s = append(s, []*ts{&ts6, &ts7, &ts8, &ts9, &ts10}...)
	}
}

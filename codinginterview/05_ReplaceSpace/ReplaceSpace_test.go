package ReplaceSpace

import "testing"

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"空格在句子中间", args{"hello world"}, "hello%20world"},
		{"空格在句子开头", args{" helloworld"}, "%20helloworld"},
		{"空格在句子末尾", args{"helloworld "}, "helloworld%20"},
		{"连续有两个空格", args{"hello  world"}, "hello%20%20world"},
		{"传入内容为空的字符串", args{""}, ""},
		{"传入内容为一个空格的字符串", args{" "}, "%20"},
		{"传入的字符串没有空格", args{"helloworld"}, "helloworld"},
		{"传入的字符串全是空格", args{"    "}, "%20%20%20%20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

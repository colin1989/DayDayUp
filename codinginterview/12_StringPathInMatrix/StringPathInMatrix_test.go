package StringPathInMatrix

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"ABTG/CFCS/JDEH", args{
			board: [][]byte{{'A', 'B', 'T', 'G'}, {'C', 'F', 'C', 'S'}, {'J', 'D', 'E', 'H'}},
			word:  "BFCT",
		}, true},
		{"ABCE/SFSC/ADEE", args{
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'S', 'C'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCB",
		}, false},
		{"ABCE/SFCS/ADEE", args{
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
		}, true},
		{"CAA/AAA/BCD", args{
			board: [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}},
			word:  "AAB",
		}, true},
		{"ABCE/SFES/ADEE", args{
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCESEEEFS",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCESEEEFS")
	}
}

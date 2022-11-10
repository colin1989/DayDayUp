package StringPathInMatrix

// 面试题12：矩阵中的路径
// 题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有
// 字符的路径。路径可以从矩阵中任意一格开始，每一步可以在矩阵中向左、右、
// 上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入
// 该格子。例如在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字
// 母用下划线标出）。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个
// 字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
// A B T G
// C F C S
// J D E H

// 传参数比len()还耗时.内存消耗更严重
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	visit := make([]bool, len(board)*len(board[0]))
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == word[0] && hasPath(board, word, row, col, 0, visit) {
				return true
			}
		}
	}

	return false
}

func hasPath(board [][]byte, word string, row, col, length int, visit []bool) bool {
	if length == len(word) {
		return true
	}
	if row >= 0 && row < len(board) && col >= 0 && col < len(board[0]) && board[row][col] == word[length] && !visit[row*len(board[0])+col] {
		visit[row*len(board[0])+col] = true
		if hasPath(board, word, row, col+1, length+1, visit) ||
			hasPath(board, word, row, col-1, length+1, visit) ||
			hasPath(board, word, row-1, col, length+1, visit) ||
			hasPath(board, word, row+1, col, length+1, visit) {
			return true
		} else {
			visit[row*len(board[row])+col] = false
		}

	}
	return false
}

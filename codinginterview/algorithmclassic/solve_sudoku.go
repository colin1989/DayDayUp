package algorithmclassic

func solveSudoku(board [][]byte) {
	sudoku := make([]byte, len(board)*len(board[0]))

	// 填充
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			sudoku[row*len(board)+col] = board[row][col]
		}
	}
	// 回溯解数独
}

package RobotMove

// 面试题13：机器人的运动范围
// 题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动，它
// 每一次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的数位之和
// 大于k的格子。例如，当k为18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。
// 但它不能进入方格(35, 38)，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

func movingCount(threshold, rows, cols int) int {
	count := 0
	// 初始化访问数据
	visit := make([]bool, rows*cols)
	moving(threshold, rows, cols, 0, 0, &count, visit)
	return count
}

func moving(threshold, rows, cols, row, col int, count *int, visit []bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols && !visit[row*cols+col] && getDigitSum(row)+getDigitSum(col) <= threshold {
		visit[row*cols+col] = true
		*count += 1
		if moving(threshold, rows, cols, row, col+1, count, visit) ||
			moving(threshold, rows, cols, row+1, col, count, visit) ||
			moving(threshold, rows, cols, row, col-1, count, visit) ||
			moving(threshold, rows, cols, row-1, col, count, visit) {
			return true
		}
	}
	return false
}

func getDigitSum(num int) int {
	sum := 0
	for num%10 > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

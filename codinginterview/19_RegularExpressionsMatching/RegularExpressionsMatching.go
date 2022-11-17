package regularexpressionsmatching

// 面试题19：正则表达式匹配
// 题目：请实现一个函数用来匹配包含'.'和'*'的正则表达式。模式中的字符'.'
// 表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题
// 中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"
// 和"ab*ac*a"匹配，但与"aa.a"及"ab*a"均不匹配。
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	if m == 0 {
		return n == 0
	}

	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	// 默认空字符串匹配成功
	dp[0][0] = true

	// 初始化首行'*'的匹配位
	for j := 0; j < m; j++ {
		// 默认跳过匹配模式'a*',是否匹配由前两位决定
		// 模式 : 'a*a*a*'
		// dp[0]值为 : true false true false true false true
		// 模式 : 'a*aa*'
		// dp[0]值为 : true false true false false false
		if p[j] == '*' && dp[0][j-1] {
			dp[0][j+1] = true
		}
	}

	// 通过遍历取出s中的每个字符,与pattern每个匹配位进行匹配
	// 低位s的匹配影响着高位s的匹配结果
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if p[j] == '*' {
				if p[j-1] == '.' || p[j-1] == s[i] {
					// 状态转移方程:
					// p[i+1][j+1] 的状态是由以下结果决定
					dp[i+1][j+1] =
						// 等同s的上位匹配结果
						//in this case, a* counts as multiple a
						dp[i][j+1] ||
							// 等同s的当前位与p的上位匹配结果
							//in this case, a* counts as single a
							dp[i+1][j] ||
							// 最后这步,'a*'无匹配结果,跳过匹配.匹配结果等同匹配串的上位
							// in this case, a* counts as empty
							dp[i+1][j-1]
				} else {
					// 无法匹配模式串'a*',该标志位由串前一个标志位决定是否匹配
					//in this case, a* only counts as empty
					dp[i+1][j+1] = dp[i+1][j-1]
				}
				// 当遇到匹配符'.'或者字符相等时.
				// 当前匹配位等于 dp[i][j]
			} else if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			}
		}
	}

	return dp[n][m]
}

func isMatch1(s string, p string) bool {
	n, m := len(s), len(p)
	var matches func(i, j int) bool = nil
	matches = func(i, j int) bool {
		if j >= m {
			return i >= n
		}

		if j+1 < m && p[j+1] == '*' {
			if i < n && (p[j] == s[i] || p[j] == '.') {
				return matches(i+1, j+2) || matches(i+1, j) || matches(i, j+2)
			} else {
				return matches(i, j+2)
			}
		} else if i < n && (p[j] == s[i] || p[j] == '.') {
			return matches(i+1, j+1)
		}
		return false
	}

	return matches(0, 0)
}

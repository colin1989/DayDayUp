package NumericStrings

// 面试题20：表示数值的字符串
// 题目：请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，
// 字符串“+100”、“5e2”、“-123”、“3.1416”及“-1E-16”都表示数值，但“12e”、
// “1a3.14”、“1.2.3”、“+-5”及“12e+5.4”都不是

// 切割字符串 A[.B][e|EC] | [.B][e|EC]
// 确认A的状态,如果A为单符号,A＋B时．B不能为单符号．
// 有B+C时,C必须在B后面
func isNumeric(s string) bool {
	s = splitSpace(s)
	n := len(s)
	if n == 0 {
		return false
	}
	hasA := true
	hasB := false
	hasC := false

	a, b, c := s, "", ""
	i, j, k := 0, -1, -1
	for ; i < n; i++ {
		if s[i] == '.' {
			if j >= 0 || k >= 0 {
				return false
			}
			j = i
			a = s[0:j]
			hasA = len(a) > 0
			b = s[j:]
			hasB = true
		} else if s[i] == 'e' || s[i] == 'E' {
			if k >= 0 || i == 0 {
				return false
			}
			k = i

			if !hasB {
				a = s[0:k]
				hasA = len(a) > 0
			} else {
				b = s[j:k]
			}
			hasC = true
			c = s[k+1:]
			if len(c) == 0 {
				return false
			}
		}
	}

	AChar := false
	singleA := false
	BChar := false
	singleB := false
	CChar := false

	if hasA {
		singleA = checkIsSingleChar(a)
		AChar = !singleA && checkIsNumber(a)
		if !singleA && !AChar {
			return false
		}
	}
	if hasB {
		singleB = checkIsSingleChar(b)
		BChar = !singleB && checkIsNumber(b)
		if !singleB && !BChar {
			return false
		}
	}
	if hasC {
		CChar = checkIsNumber(c)
		if checkIsSingleChar(c) || !CChar {
			return false
		}
	}
	if AChar {
		return true
	} else if singleA && singleB || singleA && CChar && !BChar {
		return false
	} else if !AChar && singleB {
		return false
	}
	return true
}

func splitSpace(s string) string {
	// 切割头尾空格
	for {
		if len(s) == 0 {
			break
		} else if s[0] == ' ' {
			s = s[1:]
		} else if s[len(s)-1] == ' ' {
			s = s[:len(s)-1]
		} else {
			break
		}
	}
	return s
}

func checkIsSingleChar(s string) bool {
	if len(s) > 1 {
		return false
	}
	return s[0] == '+' || s[0] == '-' || s[0] == '.'
}

func checkIsNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' || s[i] == '.' {
			if i != 0 {
				return false
			}
		} else if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

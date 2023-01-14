package NumericStrings

func isNumber(s string) bool {
	states := []map[byte]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4}, // 0
		{'d': 2, '.': 4},                 // 1
		{'d': 2, '.': 3, 'e': 5, ' ': 8}, // 2
		{'d': 3, 'e': 5, ' ': 8},         // 3
		{'d': 3},                         // 4
		{'d': 7, 's': 6},                 // 5
		{'d': 7},                         // 6
		{'d': 7, ' ': 8},                 // 7
		{' ': 8},                         // 8
	}

	p := 0
	ok := false
	for i := 0; i < len(s); i++ {
		t := byte('?')
		if s[i] >= '0' && s[i] <= '9' {
			t = 'd'
		} else if s[i] == '+' || s[i] == '-' {
			t = 's'
		} else if s[i] == 'e' || s[i] == 'E' {
			t = 'e'
		} else if s[i] == ' ' {
			t = ' '
		} else if s[i] == '.' {
			t = '.'
		}
		if p, ok = states[p][t]; !ok {
			return false
		}
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}

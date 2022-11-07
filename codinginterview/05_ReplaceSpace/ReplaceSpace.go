package ReplaceSpace

func replaceSpace(s string) string {
	s1 := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			s1 += "%"
			s1 += "2"
			s1 += "0"
		} else {
			s1 = s1 + string(s[i])
		}
	}
	return s1
}

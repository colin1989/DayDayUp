package numberof1inbinary

func hammingWeight(num int) int {
	flag, count := 1, 0
	for flag > 0 {
		if num&flag > 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}

func hammingWeight1(num uint32) int {
	count := 0
	for ; num > 0; num &= (num - 1) {
		count++
	}
	return count
}

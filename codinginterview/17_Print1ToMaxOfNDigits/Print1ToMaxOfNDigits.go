package print1tomaxofndigits

// 面试题17：打印1到最大的n位数
// 题目：输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则
// 打印出1、2、3一直到最大的3位数即999。

func Print1ToMaxOfNDigits(n int) []int {
	count := 1
	for i := 0; i < n; i++ {
		count *= 10
	}
	numbers := make([]int, count-1, count)
	for i := 1; i < count; i++ {
		numbers[i-1] = i
	}
	return numbers
}

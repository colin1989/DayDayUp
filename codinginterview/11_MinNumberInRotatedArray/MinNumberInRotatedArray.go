package MinNumberInRotatedArray

// 面试题11：旋转数组的最小数字
// 题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如数组
// {3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组的最小值为1。
func minArrayInRotatedArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	// 使用二分查找
	start := 0              // 头
	end := len(numbers) - 1 // 尾
	middle := start         //  中间值

	// 判断头尾,有可能相等
	for numbers[start] >= numbers[end] {
		if start+1 == end {
			return numbers[end]
		}
		middle = (end-start)>>1 + start
		// 三者相等时,只能遍历
		if numbers[start] == numbers[end] &&
			numbers[end] == numbers[middle] {
			return MinInOrder(numbers, start, end)
			// 有可能相等
		} else if numbers[middle] >= numbers[start] {
			start = middle
		} else if numbers[middle] <= numbers[end] {
			end = middle
		}
	}

	return numbers[middle]
}

func MinInOrder(numbers []int, start, end int) int {
	out := numbers[start]
	for i := start + 1; i <= end; i++ {
		if numbers[i] < out {
			out = numbers[i]
		}
	}
	return out
}

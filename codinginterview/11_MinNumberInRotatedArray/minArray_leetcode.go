package MinNumberInRotatedArray

// 默认递增排序
func minArray(numbers []int) int {
	start, middle := 0, 0
	end := len(numbers) - 1
	for start < end {
		middle = ((end - start) >> 1) + start
		if numbers[middle] < numbers[end] {
			end = middle
		} else if numbers[middle] > numbers[end] {
			start = middle + 1
		} else {
			end--
		}
	}
	return numbers[start]
}

package FindDuplicate

// 面试题3（二）：不修改数组找出重复的数字
// 题目：在一个长度为n+1的数组里的所有数字都在1到n的范围内，所以数组中至
// 少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的
// 数组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的
// 输出是重复的数字2或者3。
func findDuplicateNoEdit(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 1
	end := len(nums) - 1
	for {
		middle := ((end - start) >> 1) + start
		count := countRange(nums, start, middle)
		if end == start {
			if count > 1 {
				return start
			} else {
				break
			}
		}
		if count > middle-start+1 {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return -1
}

func countRange(nums []int, start, end int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}

package algorithmclassic

// 平均情况下快速排序的时间复杂度是O(nlogn)，最坏情况是O(n^2)，但通过随机算法可以避免最坏情况。由于递归调用，快排的空间复杂度是O(logn)
func partition(nums []int, start, end int) int {
	// 直接取中间值性能最高
	idx := ((end - start) >> 1) + start
	// 把该值放到最后
	nums[idx], nums[end] = nums[end], nums[idx]

	// 左边比middle小,右边比middle大
	small := start - 1
	for idx = start; idx < end; idx++ {
		if nums[idx] < nums[end] {
			small++
			if small != idx {
				nums[small], nums[idx] = nums[idx], nums[small]
			}
		}
	}
	// 排序结束,把比较数放入后一位
	small++
	nums[small], nums[end] = nums[end], nums[small]
	// 返回分割的索引
	return small
}

func QuickSort(nums []int, start, end int) []int {
	if start == end {
		return nums
	}
	idx := partition(nums, start, end)
	if idx > start {
		QuickSort(nums, start, idx-1)
	}
	if idx < end {
		QuickSort(nums, idx+1, end)
	}
	return nums
}

package sort_m

// QuickSort 快速排序
/**
 * 快速排序
 * @param array 待排序数组
 * @return 排序后的数组
 * 时间复杂度：O(nlogn)
 */

func QuickSort(array []int) []int {
	// Implement quick sort_m here
	if len(array) < 2 {
		return array
	}

	// 获取左右边界
	left, right := 0, len(array)-1
	// 获取基准值
	pivot := array[(left+right)/2]

	for left < right {
		for array[left] < pivot {
			left++
		}

		for array[right] > pivot {
			right--
		}

		if left < right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		}
	}

	// 递归调用
	QuickSort(array[:left])
	QuickSort(array[left+1:])

	return array
}

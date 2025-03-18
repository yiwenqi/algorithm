package base

import "sort"

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)

	if len(nums) < 3 {
		return result
	}

	// 排序
	sort.Ints(nums)

	// 双重循环，双指针
	/**
	 * 1. 第一层循环，遍历数组，固定第一个数
	 * 2. 第二层循环，双指针，第二个数从第一个数的下一个开始，第三个数从最后一个开始
	 * 3. 如果三个数的和等于0，加入结果集
	 * 4. 如果三个数的和大于0，第三个数左移（由于 first + second + third = 0）
	 * 此时数组为有序数组，如果三个数的和小于0，第一个数和第二个数固定，因此 第三个数从最大的开始遍历；
	 * 时间复杂度：O(n^2) + O(nlogn) = O(n^2)
	 */
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		target := -nums[first]
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second < third && nums[second]+nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return result
}

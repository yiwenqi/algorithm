package base

import "math"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	// 采用滑动窗口模式
	/**
	 * 1. 定义一个 map，存储字符的索引
	 * 2. 判断子串结束的时候，当遇到重复字符，则表示当前子串结束，计算长度
	 * 3. 重复字符的索引，表示下一个子串的开始位置
	 */
	m := make(map[byte]int)
	left, maxIndex := 0, 0
	for i := 0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok {
			if index >= left {
				maxIndex = int(math.Max(float64(maxIndex), float64(i-left)))
				left = m[s[i]] + 1
				m[s[i]] = i
				continue
			}
		}
		m[s[i]] = i
		if i == len(s)-1 && left != len(s)-1 {
			maxIndex = int(math.Max(float64(maxIndex), float64(i-left+1)))
		}
	}

	if maxIndex == 0 {
		maxIndex = len(s)
	}

	return maxIndex
}

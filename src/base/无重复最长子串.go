package base

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
	tmpStr := ""
	left, max := 0, 0
	for i := 0; i < len(s); i++ {

	}
	return 0
}

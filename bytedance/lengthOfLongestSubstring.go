package bytedance

/// 最长无重复字串
/*
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

//https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

// 暴力破解思路： 找出所有的子串，判断每个是否有重复的
func LengthOfLongestSubstring(inputString string) int {
	if len(inputString) <= 1 {
		return len(inputString)
	}

	inputs := make([]string, 0)

	for i := 0; i < len(inputString); i++ {
		for j := i + 1; j <= len(inputString); j++ {
			inputs = append(inputs, inputString[i:j])
		}
	}
	maxLen := 0

	for _, v := range inputs {
		if !isDuplicated(v) {
			if len(v) > maxLen {
				maxLen = len(v)
			}
		}
	}

	return maxLen

}

func isDuplicated(xx string) bool {
	v := make(map[rune]string)
	for _, char := range xx {
		if _, ok := v[char]; ok {
			return true
		}
		v[char] = "x"
	}
	return false
}

/* 优化思路：两个指针，
1、左指针固定，逐步扩大右指针；如果都没有重复值，则返回最大值的；
2、有重复值的话，将左指针移动到；重复字符的位置，继续向右移动右指针
*/

func LengthOfLongestSubstring2(inputStr string) int {
	if len(inputStr) <= 1 {
		return len(inputStr)
	}

	maxLen := 0
	left := 0
	dup := make(map[string]int)
	for right := 0; right < len(inputStr); right++ {
		if index, ok := dup[inputStr[right:right+1]]; ok && index >= left {
			left = index + 1
		}
		dup[inputStr[right:right+1]] = right

		curLength := right - left + 1
		if curLength > maxLen {
			maxLen = curLength
		}
	}
	return maxLen
}

package string

// hard
// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func minWindow(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := make(map[rune]int)
	for _, v := range t {
		need[v]++
	}

	window := make(map[rune]int)
	left, right, valid := 0, 0, 0

	start, length := 0, len(s)+1

	for right < len(s) {

		c := s[right]
		right++

		if _, ok := need[rune(c)]; ok {
			window[rune(c)]++
			if window[rune(c)] == need[rune(c)] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++

			if _, ok := need[rune(d)]; ok {
				if window[rune(d)] == need[rune(d)] {
					valid--
				}
				window[rune(d)]--
			}

		}

	}
	if length == len(s)+1 {
		return ""
	}

	return s[start : start+length]
}

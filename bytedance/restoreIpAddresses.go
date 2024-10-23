package bytedance

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var result []string

	var path []string

	helper3(s, 0, path, &result)

	return result

}

func helper3(s string, start int, path []string, res *[]string) {
	if len(path) == 4 && start == len(s) {
		*res = append(*res, strings.Join(path, "."))
		return
	}

	if len(path) == 4 || start == len(s) {
		return
	}

	for i := 1; i <= 3; i++ {
		if start+i > len(s) {
			break
		}
		part := s[start : start+i]
		if isValid2(part) {
			path = append(path, part)
			helper3(s, start+i, path, res)
			path = path[:len(path)-1]
		}

	}

}
func isValid2(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	nums, _ := strconv.Atoi(s)

	return nums >= 0 && nums <= 255

}

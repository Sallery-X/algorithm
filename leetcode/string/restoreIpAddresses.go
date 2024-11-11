package string

import (
	"strconv"
	"strings"
)

//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
//给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
//
//
//示例 1：
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]

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

package string

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {

	version1s := strings.Split(version1, ".")
	version2s := strings.Split(version2, ".")

	minLen := 0
	if len(version1s) > len(version2s) {
		minLen = len(version2s)
	} else {
		minLen = len(version1s)
	}

	for i := 0; i < minLen; i++ {
		v1, _ := strconv.Atoi(version1s[i])

		v2, _ := strconv.Atoi(version2s[i])

		if v1 > v2 {
			return 1
		} else if v2 > v1 {
			return -1
		} else {
			continue
		}
	}

	if len(version1s) > len(version2s) {
		for i := minLen; i < len(version1s); i++ {
			v1, _ := strconv.Atoi(version1s[i])
			if v1 != 0 {
				return 1
			}
		}
	}

	if len(version1s) < len(version2s) {
		for i := minLen; i < len(version2s); i++ {
			v1, _ := strconv.Atoi(version2s[i])
			if v1 != 0 {
				return -1
			}
		}
	}
	return 0

}

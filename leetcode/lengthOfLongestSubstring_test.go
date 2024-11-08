package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testStr := "abba"
	maxLen := LengthOfLongestSubstring(testStr)
	fmt.Println(maxLen)
	maxlen2 := LengthOfLongestSubstring2(testStr)
	fmt.Println(maxlen2)
}

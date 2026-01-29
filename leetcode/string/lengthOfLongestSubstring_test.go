package string

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testStr := "abba"
	maxLen := LengthOfLongestSubstring(testStr)
	fmt.Println(maxLen)
	maxlen2 := longestSubstring(testStr)
	fmt.Println(maxlen2)
}

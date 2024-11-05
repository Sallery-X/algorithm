package bytedance

import (
	"fmt"
	"testing"
)

func Test_minDistance(t *testing.T) {
	word1 := "kitten"
	word2 := "sitting"
	fmt.Printf("The minimum edit distance between %s and %s is %d\n", word1, word2, minDistance(word1, word2))
}

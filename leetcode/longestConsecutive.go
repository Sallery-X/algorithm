package leetcode

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool, len(nums))

	for _, v := range nums {
		numMap[v] = true
	}

	longest := 0

	for k, _ := range numMap {
		if !numMap[k-1] {
			currentNum := k
			currentLen := 1

			for numMap[currentNum+1] {
				currentNum++
				currentLen++
			}

			if currentLen > longest {
				longest = currentLen
			}

		}

	}

	return longest

}

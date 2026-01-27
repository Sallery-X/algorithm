package array

// 最长连续子串
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
/*示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

*/
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

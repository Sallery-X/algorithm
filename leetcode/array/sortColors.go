package array

import "fmt"

/*
*	[0,1,2,0,1,2,1]->[0,0,1,1,1,2,2]
 */
func sortColors(nums []int) {
	l, r := 0, len(nums)-1

	for i := 0; i <= r; {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		} else if nums[i] == 2 {
			nums[r], nums[i] = nums[i], nums[r]
			r--
		} else {
			i++
		}

	}
	fmt.Println(nums)

}

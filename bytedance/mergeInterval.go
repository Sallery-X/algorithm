package bytedance

import (
	"math"
	"sort"
)

func mergeInterval(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		v1, v2 := intervals[i], intervals[j]
		return v1[0] < v2[0] || v1[0] == v2[0] && v1[1] < v2[1]
	})
	var res [][]int
	left := intervals[0][0]
	right := intervals[0][1]

	for _, v := range intervals {
		if v[0] <= right {
			right = int(math.Max(float64(right), float64(v[1])))
		} else {
			res = append(res, []int{left, right})
			left = v[0]
			right = v[1]
		}
	}
	res = append(res, []int{left, right})
	return res
}

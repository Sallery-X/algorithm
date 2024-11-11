package matrix

import (
	"math"
	"sort"
)

//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
//
//
//示例 1：
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

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

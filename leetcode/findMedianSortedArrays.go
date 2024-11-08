package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums2) < len(nums1) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)

	l, r := 0, x

	for l <= r {
		midX := (l + r) >> 1
		midY := (x+y+1)/2 - midX

		maxX := math.MinInt32
		if midX != 0 {
			maxX = nums1[midX-1]
		}

		maxY := math.MinInt32
		if midY != 0 {
			maxY = nums2[midY-1]
		}

		minX := math.MaxInt32
		if midX != x {
			minX = nums1[midX]
		}

		minY := math.MaxInt32

		if midY != y {
			minY = nums2[midY]
		}

		if maxX <= minY && maxY <= minX {
			if (x+y)%2 == 0 {
				return float64(max(maxY, maxX)+min(minX, minY)) / 2.0
			} else {
				return float64(max(maxX, maxY))
			}
		} else if maxX > minY {
			r = midX - 1
		} else {
			l = midX + 1
		}

	}

	panic("not sorted arrays")

}

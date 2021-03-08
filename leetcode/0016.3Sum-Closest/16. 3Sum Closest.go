package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res, diff := 0, math.MaxInt32
	if len(nums) > 2 {
		sort.Ints(nums)
		for i := 0; i < len(nums)-2; i++ {
			lo, hi := i+1, len(nums)-1
			for lo < hi {
				sum := nums[lo] + nums[hi] + nums[i]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					hi--
				} else {
					lo++
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

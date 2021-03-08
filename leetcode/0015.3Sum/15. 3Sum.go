package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) > 2 {
		sort.Ints(nums)
		for i := 0; i < len(nums)-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			lo, hi, target := i+1, len(nums)-1, -nums[i]
			for lo < hi {
				if nums[lo]+nums[hi] == target {
					res = append(res, []int{nums[lo], nums[hi], nums[i]})
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--
				} else if nums[lo]+nums[hi] > target {
					hi--
				} else {
					lo++
				}
			}
		}
	}
	return res
}

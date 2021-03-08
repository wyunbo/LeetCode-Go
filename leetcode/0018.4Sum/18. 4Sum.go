package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	res, length := make([][]int, 0), len(nums)
	if length > 3 {
		sort.Ints(nums)
		for i := 0; i < length-3; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j := i + 1; j < length-2; j++ {
				if j > i+1 && nums[j] == nums[j-1] {
					continue
				}
				lo, hi := j+1, length-1
				for lo < hi {
					sum := nums[i] + nums[j] + nums[lo] + nums[hi]
					if sum == target {
						res = append(res, []int{nums[i], nums[j], nums[lo], nums[hi]})
						for lo < hi && nums[lo] == nums[lo+1] {
							lo++
						}
						for lo < hi && nums[hi] == nums[hi-1] {
							hi--
						}
						lo++
						hi--
					} else if sum < target {
						lo++
					} else {
						hi--
					}
				}
			}
		}
	}
	return res
}

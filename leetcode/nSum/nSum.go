package leetcode

func nSum(nums []int, n int, target int) []int {
	return sumN(nums, n, target, 0, make([]int, 0))
}

func sumN(nums []int, n int, target int, i int, res []int) []int {
	if target == 0 {
		return res
	}
	if i == len(nums) || n == 0 || target < 0 {
		return nil
	}
	next := sumN(nums, n-1, target-nums[i], i+1, append(res, nums[i]))
	if next == nil {
		return sumN(nums, n, target, i+1, res)
	}
	return next
}

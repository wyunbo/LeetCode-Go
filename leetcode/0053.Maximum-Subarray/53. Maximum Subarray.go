package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	local, res := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		local = max(nums[i], nums[i]+local)
		res = max(res, local)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

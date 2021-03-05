package leetcode

func maxProduct(nums []int) int {
	maximum, minimum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		maximum = max(nums[i], nums[i]*maximum)
		minimum = min(nums[i], nums[i]*minimum)
		res = max(res, maximum)
	}
	return res
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

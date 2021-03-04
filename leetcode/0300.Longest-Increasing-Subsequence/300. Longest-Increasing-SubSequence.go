package leetcode

// Solution1 O(n^2)
func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)), 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		res = max(dp[i], res)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Solution2 O(nlogn)

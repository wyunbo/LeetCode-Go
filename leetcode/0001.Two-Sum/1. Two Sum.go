package leetcode

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		another := target - v
		if idx, ok := numMap[another]; ok {
			return []int{idx, k}
		}
		numMap[v] = k
	}
	return nil
}

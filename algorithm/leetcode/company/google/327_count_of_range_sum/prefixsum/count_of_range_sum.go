package prefixsum

// Use prefix sum and nested for loop to check all sub array
// Time: O(n^2)
// Space: O(n)
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	for i := 1; i < len(prefix); i++ {
		for j := 0; j < i; j++ {
			if prefix[i]-prefix[j] >= lower && prefix[i]-prefix[j] <= upper {
				res++
			}
		}
	}
	return res
}

package prefixsum

// Use prefix sum. For every number, we calculate the current prefix sum, if we want to know
// whether there is a subarray ended by this number sum to k, we need to find another number
// whose prefix sum is equal to (curSum - k) and the index of the number should be smaller than
// current number. Use can use a map to find the number quickly.
// Time: O(n)
// Space: O(n)
func subarraySum(nums []int, k int) int {
	// be careful about the initialization
	sumMap := map[int]int{0: 1}
	res := 0
	sum := 0
	for _, num := range nums {
		sum += num
		res += sumMap[sum-k]
		sumMap[sum]++
	}
	return res
}

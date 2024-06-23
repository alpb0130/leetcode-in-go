package linearandconstant

// Prefix sum. No need to store the prefix sum of each index because when we iterate over
// the array, we can calculate the current prefix sum.
// Time: O(n)
// Space: O(1)
func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	curSum := 0
	for i, num := range nums {
		if curSum == sum-curSum-num {
			return i
		}
		curSum += num
	}
	return -1
}

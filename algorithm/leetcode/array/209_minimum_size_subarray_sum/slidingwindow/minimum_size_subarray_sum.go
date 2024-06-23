package slidingwindow

// DP and sliding window are two methods that comes up. But DP is not good at this problem because
// it hard to solve this problem by sub problem. Sliding window is the best answer. The problem
// requires us to get the min length of subarray whose sum is no less than s. We can use sliding
// window. Calculate the sum from beginning till the sum is no less than s. Then move start idx
// until the next idx will not give us eligible solution. At that time, calculate the length we get.
// Constantly update the result as we move end pointer.
// Other solutions: https://leetcode.com/problems/minimum-size-subarray-sum/discuss/59123/O(N)O(NLogN)-solutions-both-O(1)-space
// Space: O(1)
// Time: O(n)
func minSubArrayLen(s int, nums []int) int {
	start, end := 0, 0
	sum := 0
	res := len(nums) + 1
	for end < len(nums) {
		sum += nums[end]
		if sum >= s {
			for start < end && sum-nums[start] >= s {
				sum -= nums[start]
				start++
			}
			if end-start+1 < res {
				res = end - start + 1
			}
		}
		end++
	}
	if res == len(nums)+1 {
		res = 0
	}
	return res
}

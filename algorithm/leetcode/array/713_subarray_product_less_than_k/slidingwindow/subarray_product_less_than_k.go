package slidingwindow

// All number are positive integers. Use sliding window. Move end until the current product is
// no less than than k. Then move start till the current product is smaller than k.
// Reference: https://leetcode.com/problems/subarray-product-less-than-k/solution/
// Time: O(n)
// Space: O(1)
func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	prefixProd := 1
	start := 0
	end := 0
	for end < len(nums) {
		prefixProd *= nums[end]
		for start <= end && prefixProd >= k {
			prefixProd /= nums[start]
			start++
		}
		res += (end - start + 1)
		end++
	}
	return res
}

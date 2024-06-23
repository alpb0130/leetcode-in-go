package extraspace

// Time: O(n)
// Space: O(k)
func rotate(nums []int, k int) {
	k %= len(nums)
	if k == 0 || len(nums) == 0 {
		return
	}
	i := len(nums) - k - 1
	dst := make([]int, i+1)
	copy(dst, nums[:i+1])
	for j := i + 1; j < len(nums); j++ {
		nums[j-i-1] = nums[j]
	}
	for j := 0; j <= i; j++ {
		nums[j+k] = dst[j]
	}
}

package inplace

// Time: O(n)
// Space: O(1)
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 || n == 0 {
		return
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

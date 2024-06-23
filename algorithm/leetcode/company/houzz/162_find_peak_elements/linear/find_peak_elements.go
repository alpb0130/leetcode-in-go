package linear

// We are required to find any peak. So we can just find the start of first
// descending position.
// Time: O(n)
// Space: O(1)
func findPeakElement(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

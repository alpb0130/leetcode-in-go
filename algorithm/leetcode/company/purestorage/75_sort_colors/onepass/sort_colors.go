package onepass

// Three way partition. Only need one pass.
// Time: O(n)
// Space: 0(1)
func sortColors(nums []int) {
	ptr, low, high := 0, 0, len(nums)-1
	for ptr <= high {
		if nums[ptr] == 0 {
			nums[ptr], nums[low] = nums[low], nums[ptr]
			low++
			ptr++
		} else if nums[ptr] == 2 {
			nums[ptr], nums[high] = nums[high], nums[ptr]
			high--
		} else {
			ptr++
		}
	}
}

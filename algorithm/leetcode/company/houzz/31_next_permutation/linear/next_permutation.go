package linear

// Example: [4, 6, 5, 3, 2, 1] -> ([5, 6, 4, 3, 2, 1]) (transition step) -> [5, 1, 2, 3, 4, 6]
// Time: O(n)
// Space: O(1)
func nextPermutation(nums []int) {
	// from end to start, find the index of first element that is smaller than its successor
	idx := len(nums) - 2
	for idx >= 0 && nums[idx] >= nums[idx+1] {
		idx--
	}
	// from back to front, search for the first element that is larger than the element we found
	// above
	if idx >= 0 {
		for i := len(nums) - 1; i >= idx+1; i-- {
			if nums[i] > nums[idx] {
				nums[i], nums[idx] = nums[idx], nums[i]
				break
			}
		}
	}
	// sort the remaining list in ascent order
	j := idx + 1
	k := len(nums) - 1
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
}

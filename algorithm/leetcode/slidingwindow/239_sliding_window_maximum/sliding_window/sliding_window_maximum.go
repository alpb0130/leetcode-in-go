package sliding_window

// Sliding window
// Time: O(n) - each element will be put and polled once
// Space: O(1)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}
	result := []int{}
	queue := []int{}
	for i, num := range nums {
		// remove index out of window
		if len(queue) != 0 && queue[0] == i-k {
			queue = queue[1:]
		}
		// remove index whose num is smaller than the current number
		j := 0
		for {
			if j == len(queue) || nums[queue[j]] <= num {
				break
			}
			j++
		}
		queue = queue[:j]
		queue = append(queue, i)
		// construct result
		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}
	}
	return result
}

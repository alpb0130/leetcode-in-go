package twopass

// Use map to count the number. Need two pass
// Time: O(n)
// Space: 0(1)
func sortColors(nums []int) {
	colors := make([]int, 3)
	for _, color := range nums {
		colors[color]++
	}
	i := 0
	for j, c := range colors {
		for {
			if c == 0 {
				break
			}
			nums[i] = j
			i++
			c--
		}
	}
}

package vn

// The problem requires to return the elevation array after we pour V water at position K.
// We just need to simulate the process for each water unit and update the current elevation.
// The key for this problem is to find the FIRST lowest elevation in the contiguous array ended
// with K in left/right direction.
// Time: O(V*N)
// Space: O(1)
func pourWater(heights []int, V int, K int) []int {
	for i := 1; i <= V; i++ {
		// fall left
		min := K
		for j := K; j >= 0; j-- {
			if j < K && heights[j] > heights[j+1] {
				break
			}
			if heights[j] < heights[min] {
				min = j
			}
		}
		if min != K {
			heights[min]++
			continue
		}
		for j := K; j < len(heights); j++ {
			if j > K && heights[j] > heights[j-1] {
				break
			}
			if heights[j] < heights[min] {
				min = j
			}
		}
		if min != K {
			heights[min]++
			continue
		}
		heights[K]++
	}
	return heights
}

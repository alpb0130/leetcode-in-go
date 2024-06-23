package linear

// One pass and simulation the process. If there is a spot whose left and right are empty, put a flower.
// We do greedy process and put flower as early as possible.
// Time: O(n)
// Space: O(1)
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i, bed := range flowerbed {
		if bed == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			count++
			flowerbed[i] = 1
		}
	}
	return count >= n
}

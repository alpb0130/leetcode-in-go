package slidingwindowii

import "math"

// Still a sliding window method.
// Time: O(n)
// Space: O(1)
func totalFruit(tree []int) int {
	res := 0
	first := 0
	second := -1
	count := 0
	for i := 0; i < len(tree); i++ {
		count++
		if tree[i] == tree[first] {
			first = i
		} else if second == -1 || tree[i] == tree[second] {
			second = first
			first = i
		} else {
			res = maxInt(count-1, res)
			count = first - second + 1
			second, first = first, i
		}
	}

	return maxInt(count, res)
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

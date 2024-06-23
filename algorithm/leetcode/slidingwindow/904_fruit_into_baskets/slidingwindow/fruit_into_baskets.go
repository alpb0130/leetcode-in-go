package slidingwindow

// The problem is actually equals to finding the longest contiguous subarray that contains only
// 2 unique elements. Use sliding window is straightforward. Use a map the record existing values
// in current window.
// Time: O(n)
// Space: O(1)
func totalFruit(tree []int) int {
	fruitMap := map[int]int{}
	res := 0
	start := 0
	end := 0
	for end < len(tree) {
		fruitMap[tree[end]]++
		for len(fruitMap) > 2 && start <= end {
			fruitMap[tree[start]]--
			if fruitMap[tree[start]] == 0 {
				delete(fruitMap, tree[start])
			}
			start++
		}
		if end-start+1 > res {
			res = end - start + 1
		}
		end++
	}
	return res
}

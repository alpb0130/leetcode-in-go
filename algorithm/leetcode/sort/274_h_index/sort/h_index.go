package sort

import "sort"

// Sort the array in desc order the iterate it
// Time: O(nlogn)
// Space: O(1)
func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	i := 0
	for i < len(citations) {
		if citations[i] < i+1 {
			break
		}
		i++
	}
	return i
}

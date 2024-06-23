package sort2

import "sort"

/**
* Definition for an interval.
* type Interval struct {
*	   Start int
*	   End   int
* }
 */
type Interval struct {
	Start int
	End   int
}

// Sort start and end array separately and then iterate (a little bit tricky)
// Time: O(nlogn)
// Space: O(n) - depend on sort algorithm
func merge(intervals []Interval) []Interval {
	results := []Interval{}
	if len(intervals) == 0 {
		return results
	}
	startArray := make([]int, len(intervals))
	endArray := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		startArray[i] = intervals[i].Start
		endArray[i] = intervals[i].End
	}
	sort.Ints(startArray)
	sort.Ints(endArray)
	for i, j := 0, 0; i < len(intervals); i++ {
		if i == len(intervals)-1 || startArray[i+1] > endArray[i] {
			results = append(results, Interval{startArray[j], endArray[i]})
			j = i + 1
		}
	}

	return results
}

package sort

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

// Sort first and then iterate
// Time: O(nlogn)
// Space: O(n) - depend on sort algorithm
func merge(intervals []Interval) []Interval {
	results := []Interval{}
	if len(intervals) == 0 {
		return results
	}
	sort.Sort(Intervals(intervals))
	start, end := intervals[0].Start, intervals[0].End
	for _, interval := range intervals {
		if interval.Start > end {
			results = append(results, Interval{start, end})
			start = interval.Start
			end = interval.End
		} else {
			if interval.Start < start {
				start = interval.Start
			}
			if interval.End > end {
				end = interval.End
			}
		}
	}
	// Very important!!! Don't miss it!
	results = append(results, Interval{start, end})
	return results
}

type Intervals []Interval

func (intervals Intervals) Len() int           { return len(intervals) }
func (intervals Intervals) Less(i, j int) bool { return intervals[i].Start < intervals[j].Start }
func (intervals Intervals) Swap(i, j int)      { intervals[i], intervals[j] = intervals[j], intervals[i] }

package auxiliararray

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

// Time: O(n)
// Space: O(n)
func insert(intervals []Interval, newInterval Interval) []Interval {
	results := []Interval{}
	if len(intervals) == 0 {
		return append(results, newInterval)
	}

	// Put unmerged interval in left and right array
	start, end := newInterval.Start, newInterval.End
	left, right := []Interval{}, []Interval{}
	for _, interval := range intervals {
		if interval.End < start {
			left = append(left, interval)
		} else if interval.Start > end {
			right = append(right, interval)
		} else {
			if interval.Start < start {
				start = interval.Start
			}
			if interval.End > end {
				end = interval.End
			}
		}
	}
	results = append(results, left...)
	results = append(results, Interval{start, end})
	results = append(results, right...)
	return results
}

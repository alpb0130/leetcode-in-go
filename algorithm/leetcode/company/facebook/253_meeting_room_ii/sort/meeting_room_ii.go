package sort

import (
	"math"
	"sort"
)

// The problem requires us the find the minimum num of meeting room required. We actually
// need to find the max # of meetings that occur at the same time. What we can do is
// sort the start time and end time separately and iterate over it. Every time we see a smaller
// start time, we increase counter by 1 which mean we need one more room. If we see a end time
// is smaller or equal to current start time, we know a meeting is over and we can reduce
// counter by 1.
// Time: O(nlogn)
// Space: O(n)
func minMeetingRooms(intervals [][]int) int {
	starts := []int{}
	ends := []int{}
	for _, interval := range intervals {
		starts = append(starts, interval[0])
		ends = append(ends, interval[1])
	}
	sort.Ints(starts)
	sort.Ints(ends)
	res := 0
	count := 0
	i, j := 0, 0
	for i < len(intervals) {
		if starts[i] < ends[j] {
			count++
			res = maxInt(res, count)
			i++
		} else {
			count--
			j++
		}
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

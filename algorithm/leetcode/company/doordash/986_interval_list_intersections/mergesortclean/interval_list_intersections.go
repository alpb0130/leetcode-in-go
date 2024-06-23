package mergesortclean

import "math"

type Interval struct {
	Start int
	End   int
}

// The problem gives us two sort disjoint intervals. We are required to find all intersections.
// Because they are sorted list and disjoint within a list, we can use something like merge sort
// to process two list. For every interval pair from two list, we get the max start and min end
// which is possible intersection. If start > end, it means two intervals don't have intersection.
// Then we move pointer according to the end. Move the pointer with small end one step
// forward. If two ends are same, move both.
// Time: O(m+n)
// Space: O(1)
func intervalIntersection(A []Interval, B []Interval) []Interval {
	res := []Interval{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		start := maxInt(A[i].Start, B[j].Start)
		end := minInt(A[i].End, B[j].End)
		if start <= end {
			res = append(res, Interval{start, end})
		}
		if A[i].End < B[j].End {
			i++
		} else if A[i].End > B[j].End {
			j++
		} else {
			i++
			j++
		}
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

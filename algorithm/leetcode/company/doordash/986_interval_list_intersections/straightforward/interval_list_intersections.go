package straightforward

import "math"

type Interval struct {
	Start int
	End   int
}

// Like merge sort.
// Time: O(m+n)
// Space: O(1)
func intervalIntersection(A []Interval, B []Interval) []Interval {
	res := []Interval{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		aStart, aEnd := A[i].Start, A[i].End
		bStart, bEnd := B[j].Start, B[j].End
		if aEnd > bEnd {
			if aStart <= bEnd {
				res = append(res, Interval{maxInt(aStart, bStart), minInt(aEnd, bEnd)})
			}
			j++
		} else if aEnd < bEnd {
			if bStart <= aEnd {
				res = append(res, Interval{maxInt(aStart, bStart), minInt(aEnd, bEnd)})
			}
			i++
		} else {
			res = append(res, Interval{maxInt(aStart, bStart), minInt(aEnd, bEnd)})
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

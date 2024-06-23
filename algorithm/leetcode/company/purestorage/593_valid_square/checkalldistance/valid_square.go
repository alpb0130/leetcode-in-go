package checkalldistance

// There are 6 possible distance in total ([pq, p2] is the same as [p2, p1]). We can compute
// them all and check whether the larger two are equal and smaller four are equal.
//
// For followup: given n vertices decide how many squares we can have. Brute force is O(n^4) and
// just enum all possible vertices combinations. Best way is O(n^2). Just enum two points combine
// and we can decide the other two points by this way:
// https://www.quora.com/Given-two-diagonally-opposite-points-of-a-square-how-can-I-find-out-the-other-two-points-in-terms-of-the-coordinates-of-the-known-points
// Given (x1, y1), (x2, y2), the mid is ((x1+x2)/2, (y1+y2)/2). Then the vector from (x1, y1) to
// mid is ((x1-x2)/2, (y1-y2)/2). We rotate the vector by 90 and get t = ((y2-y1)/2, (x1-x2)/2).
// So another two points are m-t (from m-p = t) and m+t (from p-m = t)
// Time: O(1)
// Space: O(1)
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	dists := []int{
		dist(p1, p2),
		dist(p1, p3),
		dist(p1, p4),
		dist(p2, p3),
		dist(p2, p4),
		dist(p3, p4),
	}
	diag := 0
	edge := dist(p1, p2)
	for _, dist := range dists {
		if dist > diag {
			diag = dist
		}
		if dist < edge {
			edge = dist
		}
	}
	if diag == 0 || edge == 0 || 2*edge != diag {
		return false
	}
	diagCnt := 0
	for _, dist := range dists {
		if dist == diag {
			diagCnt++
		} else if dist != edge {
			return false
		}
	}
	return diagCnt == 2
}

func dist(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

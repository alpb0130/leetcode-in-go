package sort

import "sort"

// The intuition is we need to compute the distance between points and verify whether the
// edges have same length and diagonals have the same length. But how do we know whether
// two points form any edge or diagnose? Sort them by x and break tie by y, we can easily
// found that 0-1, 0-2, 2-3, 1-3 will always be edge and 0-3, 1-3 will always be diagonals.
// 		# (2)                # (1)            # (1)   # (3)
// # (0)					      # (3)
//        # (3)            # (0)
//	 # (1)                      # (2)         # (0)   # (2)
// So we just very all edge and diagonals. Also make sure the distance is not 0
// Time: O(1)
// Space: O(1)
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	p := Points{p1, p2, p3, p4}
	sort.Sort(p)
	return dist(p[0], p[1]) != 0 &&
		dist(p[0], p[1]) == dist(p[0], p[2]) &&
		dist(p[0], p[2]) == dist(p[2], p[3]) &&
		dist(p[2], p[3]) != 0 &&
		dist(p[2], p[3]) == dist(p[1], p[3]) &&
		dist(p[0], p[3]) == dist(p[1], p[2])
}

func dist(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

type Points [][]int

func (p Points) Len() int {
	return len(p)
}
func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Points) Less(i, j int) bool {
	if p[i][0] != p[j][0] {
		return p[i][0] < p[j][0]
	}
	return p[i][1] < p[j][1]
}

package quickselect

import "math/rand"

// Quick select. Be careful about the edge case.
// Time: O(n) - shuffle + quick select
// Space: O(n)
func kClosest(points [][]int, K int) [][]int {
	// shuffle
	for i := 0; i < len(points); i++ {
		j := rand.Intn(i + 1)
		points[i], points[j] = points[j], points[i]
	}
	quickSelect(points, K-1)
	return points[:K]
}

func calcDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func quickSelect(points [][]int, k int) {
	start := 0
	end := len(points) - 1
	for {
		mid := partition(points, start, end)
		if mid < k {
			start = mid + 1
		} else if mid > k {
			end = mid - 1
		} else {
			break
		}
	}
}

func partition(points [][]int, start, end int) int {
	i := start
	j, k := start+1, end
	for {
		// Be careful here.
		for j <= k && calcDistance(points[j]) <= calcDistance(points[i]) {
			j++
		}
		for k >= j && calcDistance(points[k]) >= calcDistance(points[i]) {
			k--
		}
		if j > k {
			points[i], points[k] = points[k], points[i]
			break
		}
		points[j], points[k] = points[k], points[j]
		j++
		k--
	}
	return k
}

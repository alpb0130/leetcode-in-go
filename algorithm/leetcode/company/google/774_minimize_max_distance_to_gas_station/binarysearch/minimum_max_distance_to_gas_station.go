package binarysearch

import "math"

// Binary search. If the max interval is w, we can do binary search over [0, w]
// and try to verify whether the mid length is valid or not. By valid, I mean
// whether we can make the max gas station distance less than the mid by adding
// k station into intervals.
// Time: O(n*logw) - w is the range of the interval length
// Space: O(n)
func minmaxGasDist(stations []int, K int) float64 {
	max := 0.0
	intervals := []float64{}
	for i := 1; i < len(stations); i++ {
		dist := float64(stations[i] - stations[i-1])
		intervals = append(intervals, dist)
		max = math.Max(max, dist)
	}
	left, right := 0.0, max
	delta := 0.000001
	res := 0.0
	for math.Abs(left-right) >= delta {
		mid := (left + right) / 2
		if isValid(intervals, mid, K) {
			res = mid
			right = mid - delta
		} else {
			left = mid + delta
		}
	}
	return res
}

func isValid(intervals []float64, target float64, K int) bool {
	numStations := 0.0
	// min stations required in order to reduce the max interval below target
	for _, interval := range intervals {
		numStations += math.Max(0, math.Ceil(interval/target)-1)
	}
	return int(numStations) <= K
}

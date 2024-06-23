package easyslope

import "math"

type Point struct {
	X int
	Y int
}

// Brute force. For each point, use a map to record how many point after it on the same line.
// The problem is hwo to map two node in a line. We can use the slope as the key. The y-intercept
// is not important because for a line passing a given node. It's not possible two line with same
// slope having the same y-intercept. Why we don't need look at the line map globally? Because for
// point k, we have look at all points after it and when we look at point k+1, no need to look
// back. If k, k+1, k+2 are on the same line, then when looking at k+1, we can only look at
// k+1 and k+2. One problem with this is when we use the absolution value to represent slope,
// there might be precision loss.
// Time: O(n)
// Space: O(n)
func maxPoints(points []Point) int {
	max := 0
	for i := 0; i < len(points); i++ {
		lineMap := map[float64]int{}
		curMax := 0
		vertical := 1
		dup := 1
		for j := i + 1; j < len(points); j++ {
			if points[i].X == points[j].X && points[i].Y != points[j].Y {
				vertical++
			} else if points[i].X == points[j].X && points[i].Y == points[j].Y {
				dup++
			} else {
				s := getSlope(points[i], points[j])
				lineMap[s]++
				curMax = maxInt(curMax, lineMap[s])
			}
		}
		curMax = maxInt(curMax+dup, vertical)
		max = maxInt(max, curMax)
	}
	return max
}

func getSlope(point1, point2 Point) float64 {
	return float64(point1.Y-point2.Y) / float64(point1.X-point2.X)
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

package slopenoprecisionloss

import (
	"math"
	"strconv"
)

type Point struct {
	X int
	Y int
}

// The same as another method but use string to represent precision loss. We can calculate the
// GCD of numerator and denominator and get the lowest term of the fraction and represent in
// string.
// Time: O(n)
// Space: O(n)
func maxPoints(points []Point) int {
	max := 0
	for i := 0; i < len(points); i++ {
		lineMap := map[string]int{}
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

func getSlope(point1, point2 Point) string {
	gcd := getGCD(point1.Y-point2.Y, point1.X-point2.X)
	a := (point1.Y - point2.Y) / gcd
	b := (point1.X - point2.X) / gcd
	return strconv.Itoa(a) + "+" + strconv.Itoa(b)

}

func getGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return getGCD(b, a%b)
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

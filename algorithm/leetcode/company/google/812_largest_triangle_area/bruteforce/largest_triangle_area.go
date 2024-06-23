package bruteforce

import (
	"fmt"
	"math"
)

// Enumerate all triangles and compute area
// O(n^3)
// O(1)
func largestTriangleArea(points [][]int) float64 {
	if len(points) < 3 {
		return 0.0
	}
	n := len(points)
	res := 0.0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				res = math.Max(res, getArea(points[i], points[j], points[k]))
				if getArea(points[i], points[j], points[k]) == 18.5 {
					fmt.Println(points[i], points[j], points[k])
				}
			}
		}
	}
	return res
}

// https://en.wikipedia.org/wiki/Shoelace_formula
func getArea(p1, p2, p3 []int) float64 {
	floatP1 := []float64{float64(p1[0]), float64(p1[1])}
	floatP2 := []float64{float64(p2[0]), float64(p2[1])}
	floatP3 := []float64{float64(p3[0]), float64(p3[1])}
	return 0.5 * math.Abs(floatP1[0]*floatP2[1]+floatP2[0]*floatP3[1]+floatP3[0]*floatP1[1]-
		floatP1[0]*floatP3[1]-floatP3[0]*floatP2[1]-floatP2[0]*floatP1[1])
}

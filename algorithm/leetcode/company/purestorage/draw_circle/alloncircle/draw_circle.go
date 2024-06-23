package alloncircle

import (
	"math"
)

// Time O(logrr)
func Draw(rr int) [][]int {
	if rr == 0 {
		return [][]int{}
	}
	res := [][]int{}
	r := int(math.Sqrt(float64(rr)))
	for x := 0; x <= r; x++ {
		yy := math.Sqrt(float64(rr - x*x))
		y := int(yy)
		if yy == float64(y) {
			res = append(res, []int{x, y})
			res = append(res, []int{-x, -y})
			if y != 0 && x != 0 {
				res = append(res, []int{x, -y})
				res = append(res, []int{-x, y})
			}
		}
	}
	return res
}

// Better
func Draw2(rr int) [][]int {
	if rr == 0 {
		return [][]int{}
	}
	res := [][]int{}
	r := math.Sqrt(float64(rr))
	startX := int(math.Ceil(math.Sqrt(float64(rr / 2))))
	for x := startX; x <= int(r); x++ {
		yy := math.Sqrt(float64(rr - x*x))
		y := int(yy)
		if yy == float64(y) {
			if y == x || x == 0 || y == 0 {
				res = append(res, [][]int{{x, y}, {-x, -y}, {-y, x}, {y, -x}}...)
			} else {
				res = append(res, [][]int{{x, y}, {x, -y}, {-x, y}, {-x, -y}, {y, x}, {y, -x}, {-y, x}, {-y, -x}}...)
			}
		}
	}
	return res
}

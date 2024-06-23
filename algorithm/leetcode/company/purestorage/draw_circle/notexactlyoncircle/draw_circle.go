package notexactlyoncircle

import "math"

func Draw(rr int) [][]int {
	if rr == 0 {
		return [][]int{}
	}
	res := [][]int{}
	startY := int(math.Sqrt(float64(rr / 2)))
	for y := 0; y <= int(startY); y++ {
		xx := math.Sqrt(float64(rr - y*y))
		x := int(math.Round(xx))
		if y == int(y) || x == 0 || y == 0 {
			res = append(res, [][]int{{x, y}, {-x, -y}, {-y, x}, {y, -x}}...)
		} else {
			res = append(res, [][]int{{x, y}, {x, -y}, {-x, y}, {-x, -y}, {y, x}, {y, -x}, {-y, x}, {-y, -x}}...)
		}
	}
	return res
}

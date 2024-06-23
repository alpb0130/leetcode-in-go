package overlap

import "math"

// res = rect1+rect2-overlap
// Time: O(1)
// Space: O(1)
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	res := 0
	res += areaRect(A, B, C, D) + areaRect(E, F, G, H)
	// compute overlap rectangle boundary
	// Take max left, min right, max bottom, min up and see whether they are valid
	left := maxInt(A, E)
	right := minInt(C, G)
	bottom := maxInt(B, F)
	up := minInt(D, H)
	overlap := 0
	if left < right && bottom < up {
		overlap = areaRect(left, bottom, right, up)
	}
	return res - overlap
}

func areaRect(A, B, C, D int) int {
	return absInt(A-C) * absInt(B-D)
}

func absInt(a int) int {
	return int(math.Abs(float64(a)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

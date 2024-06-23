package byposition

import "math"

func Sqrt(n float64, delta float64) float64 {
	if n == 0 {
		return 0
	}
	left, right := 1, int(n)
	res := 0
	for left <= right {
		mid := (left + right) / 2
		if mid <= int(n)/mid {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if math.Abs(float64(res-left)) <= delta {
		return float64(res)
	}

	final := float64(res)
	pres := 0.1
	for pres >= delta {
		tmp := final
		left, right = 0, 9
		for left <= right {
			mid := (left + right) / 2
			if tmp+pres*float64(mid) <= n/(tmp+pres*float64(mid)) {
				final = tmp + pres*float64(mid)
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		pres /= 10
	}
	return final
}

package deltamove

func Sqrt(n float64, delta float64) float64 {
	if n == 0 {
		return 0
	}
	left, right := 1.0, n
	res := 0.0
	for left <= right {
		mid := (left + right) / 2
		if mid <= n/mid {
			res = mid
			left = mid + delta
		} else {
			right = mid - delta
		}
	}

	return res
}

package generatelevelbylevel

func grayCode(n int) []int {
	res := []int{0}
	for i := 0; i < n; i++ {
		size := len(res)
		for j := size - 1; j >= 0; j-- {
			res = append(res, res[j]|1<<uint(i))
		}
	}
	return res
}

package backtracking

// Get all the factors first and then backtracking. Time and space complexity are really
// depend on
// Time: O(logn*logn) -  for number 2^x, the total # of calls are x + (x-1) + ... + 1 = O(x^2)
//       which is O(logn*logn)
// Space: O(logn) - the worst case like 4, 8, 16...
func getFactors(n int) [][]int {
	factors := []int{}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	res := [][]int{}
	if len(factors) == 0 {
		return res
	}
	getFactorsHelper(factors, n, 0, []int{}, &res)
	return res
}

func getFactorsHelper(factors []int, n int, start int, comb []int, res *[][]int) {
	if n == 1 {
		dst := make([]int, len(comb))
		copy(dst, comb)
		*res = append(*res, dst)
		return
	}
	for i := start; i < len(factors); i++ {
		if n%factors[i] == 0 {
			getFactorsHelper(factors, n/factors[i], i, append(comb, factors[i]), res)
		}
	}
}

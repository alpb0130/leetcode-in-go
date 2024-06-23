package best

// Use K as carry on and append number with high significant to the end of existing result.
// Reverse the results finally. This would improve time and space a lot we don't need to
// create a lot of new slices.
// Time: O(max(len(A), len(K))
// Space: O(max(len(A), len(K)))
func addToArrayForm(A []int, K int) []int {
	res := []int{}
	for i := len(A)-1; i >= 0; i-- {
		K += A[i]
		res = append(res, K%10)
		K /= 10
	}
	for K > 0 {
		res = append(res, K%10)
		K /= 10
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

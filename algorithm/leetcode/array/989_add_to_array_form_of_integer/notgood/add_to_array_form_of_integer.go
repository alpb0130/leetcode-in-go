package notgood

// The code is simpler than the best method but it need to create a lot of new
// temporary slice. The time and space are not optimized.
// Time: O(max(len(A), len(K))
// Space: O(max(len(A), len(K))^2)
func addToArrayForm(A []int, K int) []int {
	res := []int{}
	for i := len(A) - 1; i >= 0; i-- {
		K += A[i]
		res = append([]int{K % 10}, res...)
		K /= 10
	}
	for K > 0 {
		res = append([]int{K % 10}, res...)
		K /= 10
	}
	return res
}

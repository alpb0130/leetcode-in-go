package linear

import (
	"strconv"
	"strings"
)

// Build the permutation position by position. As we can imagine, all permutations of n are:
// 1, P(2,...,n) - include (n-1) permutations
// 2, P(1,3,...,n) - include (n-1) permutations
// 3, P(1,2,3,...,n) - include (n-1) permutations
// ...
// n, P(1,2,...,n-1) - include (n-1) permutations
// For kth seq, we can get the index of first number by using k/fact(n-1) and the new order
// of next level bucket is k%fact(n-1). Do that till n = 1.
// Time: O(n)
// Space: O(n) - factorials and nums
func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i+1
	}
	factorials := make([]int, n+1)
	factorials[0] = 1
	for i := 1; i <= n; i++ {
		factorials[i] = factorials[i-1]*i
	}
	var str strings.Builder
	k--
	for i := 1; i <= n; i++ {
		idx := k/factorials[n-i]
		k = k%factorials[n-i]
		str.WriteString(strconv.Itoa(nums[idx]))
		nums = append(nums[:idx], nums[idx+1:]...)
	}
	return str.String()
}

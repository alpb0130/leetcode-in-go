package second_best

// Start from 2 (skip 1 because 1 will always be a factor). For every number,
// we find all numbers (< n) that has this number as a factor and mark it as
// non-prime. But when we check 5, we don't need to check 2*5 ane 3*5 because
// those are already been checked when we check 2 and 3. We start from j == i
// and keep increase j by 1
// Time: O(nlogn)
// Space: O(n)
func countPrimes(n int) int {
	nonPrime := make([]bool, n)
	res := 0
	for i := 2; i < n; i++ {
		if nonPrime[i] {
			continue
		}
		res++
		for j := i; j*i < n; j++ {
			nonPrime[j*i] = true
		}
	}
	return res
}

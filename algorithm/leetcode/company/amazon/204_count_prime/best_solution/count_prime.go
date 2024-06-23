package best_solution

// Time: O(lognlogn)
// Space: O(n)
func countPrimes(n int) int {
	nonPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if nonPrime[i] {
			continue
		}
		for j := i; j*i < n; j++ {
			nonPrime[j*i] = true
		}
	}
	res := 0
	for i := 2; i < n; i++ {
		if !nonPrime[i] {
			res++
		}
	}
	return res
}

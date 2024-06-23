package recursive

// Pretty bad solution
// Time: O(Catalan(n))
// Space: O(n)
func numTrees(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	res := 0
	for i := 1; i <= n; i++ {
		res += numTrees(i-1) * numTrees(n-i)
	}
	return res
}

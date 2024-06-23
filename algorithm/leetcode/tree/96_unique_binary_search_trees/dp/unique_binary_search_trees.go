package dp

// F(n) = sum(F(i-1)*F(n-i)) (i = 1, 2,...,n)
// Time: O(n^2)
// Space: O(n)
func numTrees(n int) int {
	numMap := map[int]int{0: 1}
	for i := 1; i <= n; i++ {
		num := 0
		for j := 1; j <= i; j++ {
			num += numMap[j-1] * numMap[i-j]
		}
		numMap[i] = num
	}
	return numMap[n]
}

package countingsortbetter

// Improved counting sort
// Time: O(n) - n is len(citations)
// Space: O(n)
func hIndex(citations []int) int {
	counter := map[int]int{}
	n := len(citations)
	for _, citation := range citations {
		if citation > n {
			counter[n]++
		} else {
			counter[citation]++
		}
	}
	num := 0
	i := n
	for i >= 0 {
		num += counter[i]
		if num >= i {
			break
		}
		i--
	}
	return i
}

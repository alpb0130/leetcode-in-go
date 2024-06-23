package iterative2

// Time: O(logn)??
// Space: O(1)
func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	div := []int{2, 3, 5}
	for _, n := range div {
		for num%n == 0 {
			num /= n
		}
	}
	return num == 1
}

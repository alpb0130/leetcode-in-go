package twopointer

// Reference: https://leetcode.com/problems/swap-adjacent-in-lr-string/solution/
// Time: O(n)
// Space: O(1)
func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}
	n := len(start)
	i, j := 0, 0
	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}
		if (i < n) != (j < n) {
			return false
		}
		if i >= n {
			break
		}
		if start[i] != end[j] || (start[i] == 'L' && i < j) || (start[i] == 'R' && i > j) {
			return false
		}
		i++
		j++
	}
	return true
}

package _map

// Simulation the process and record all numbers we see. If we see it again, it means we
// have cycles.
// Time: O(1)
// Space: O(1) because we use extra space but the total number of extra space should be small
func isHappy(n int) bool {
	numMap := map[int]bool{}
	for n != 1 {
		if numMap[n] {
			return false
		}
		numMap[n] = true
		m := 0
		for n != 0 {
			r := n % 10
			m += r * r
			n /= 10
		}
		n = m
	}
	return true
}

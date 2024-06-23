package linear

// If the roman is followed by a larger roman, then we subtract it from sum. Otherwise
// we add to sum.
// Time: O(n)
// Space: O(1)
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	valMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && valMap[s[i]] < valMap[s[i+1]] {
			sum -= valMap[s[i]]
		} else {
			sum += valMap[s[i]]
		}
	}
	return sum
}

package twomap

// Use two separate map for s and t and map them to the same index. If they are not mapping to the
// same index, return false.
// Time: O(n)
// Space: O(n)
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := map[byte]int{}
	tMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}
		// use i+1 so that we don't reuse the zero value 0 as a valid value
		sMap[s[i]] = i + 1
		tMap[t[i]] = i + 1
	}
	return true
}

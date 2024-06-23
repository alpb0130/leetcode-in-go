package map1

// Use one map to record char in s to char in t. Another map to record whether the dst char in t
// ha been used. If used, should be false.
// Time: O(n)
// Space: O(n)
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charMap := map[byte]byte{}
	existMap := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]
		c, ok := charMap[charS]
		if ok && c != charT {
			return false
		}
		if !ok && existMap[charT] {
			return false
		}
		existMap[charT] = true
		charMap[charS] = charT
	}
	return true
}

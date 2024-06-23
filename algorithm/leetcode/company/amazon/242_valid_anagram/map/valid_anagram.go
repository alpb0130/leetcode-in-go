package _map

// Time: O(n)
// Space: O(n)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		charMap[t[i]]--
		if charMap[t[i]] == 0 {
			delete(charMap, t[i])
		}
	}
	return len(charMap) == 0
}

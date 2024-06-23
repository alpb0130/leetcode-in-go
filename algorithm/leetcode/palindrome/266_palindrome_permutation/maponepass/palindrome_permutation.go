package maponepass

// In the string, there are at most 1 letter with odd count
// Time: O(n)
// Space: O(1)
func canPermutePalindrome(s string) bool {
	counter := map[rune]bool{}
	for _, c := range s {
		if counter[c] {
			delete(counter, c)
		} else {
			counter[c] = true
		}
	}
	return len(counter) <= 1
}

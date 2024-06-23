package maptwopass

// In the string, there are at most 1 char with odd count
// Time: O(n)
// Space: O(1)
func canPermutePalindrome(s string) bool {
	counter := map[rune]int{}
	for _, c := range s {
		counter[c]++
	}
	oddCnt := 0
	for _, cnt := range counter {
		oddCnt += cnt % 2
	}
	return oddCnt <= 1
}

package twopointer

// We iterate over each letter and use two pointers to check whether whether the current
// substring is palindromic. We keep expanding the two point till it's not palindromic.
// But for every iteration, we need check twice. One for only one pivotal string and
// another for two neighbor pivotal string.
// Time: O(n^2)
// Space: O(1)
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	var left, right int
	max := 0
	for i := 0; i < len(s); i++ {
		l, r := longestSubPalindrome(s, i, i)
		if r-l+1 > max {
			max = r - l + 1
			left = l
			right = r
		}
		l, r = longestSubPalindrome(s, i, i+1)
		if r-l+1 > max {
			max = r - l + 1
			left = l
			right = r
		}
	}
	return s[left : right+1]
}

func longestSubPalindrome(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			return i + 1, j - 1
		}
		i--
		j++
	}
	return i + 1, j - 1
}

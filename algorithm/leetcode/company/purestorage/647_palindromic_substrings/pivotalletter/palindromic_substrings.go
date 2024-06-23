package pivotalletter

// For the string, start from position 0 to len(s)-1 and set it as pivotal letter.
// Then we use two pointers and expand the substring and check whether current letters are
// the same or not. If the same, we can update res and keep expanding. Otherwise, stop here
// and move to the next pivotal string.
// Also take care the cases where we have two pivotal letters.
// Time: O(n^2)
// Space: O(1)
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(s); i++ {
		j, k := i, i
		for j >= 0 && k < len(s) && s[j] == s[k] {
			res++
			j--
			k++
		}
		if i < len(s)-1 {
			j, k = i, i+1
		}
		for j >= 0 && k < len(s) && s[j] == s[k] {
			res++
			j--
			k++
		}
	}
	return res
}

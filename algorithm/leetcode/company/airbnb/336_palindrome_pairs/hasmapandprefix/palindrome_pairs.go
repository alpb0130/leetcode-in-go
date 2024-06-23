package hasmapandprefix

// Use prefix and suffix to match string.
// Firstly, we use a map to record the index of each string.
// Then, if str1+str2 is a palindrome (assume l1 > l2), then reverse(substr1) must = str2.
// We use this feature to go over all strings. But for inner loop, instead go over all other
// string, we can go over each char of current string and get different prefix and suffix
// combination. If reverse(prefix) is equal to str2, then (str1, str2). If reverse(suffix)
// is equal to str2, then (str2, str1).
// Remember to dedupe for the case where two strings have the same length and can be palindrome.
// For all other case, only when we visit the longer string, we can get the pairs.
// Time: O(n*k^2) - the outer loop is n. The inner loop is k. For each loop, we need to reverse
// and check whether a string is palindrome. That's k. So totally O(n*k^2)
// Space: O(n*k) - the map
func palindromePairs(words []string) [][]int {
	res := [][]int{}
	if len(words) == 0 {
		return res
	}
	// reverse string and build map
	strMap := map[string]int{}
	for i, word := range words {
		strMap[word] = i
	}

	for i := 0; i < len(words); i++ {
		str := words[i]
		for j := 0; j < len(str)+1; j++ {
			prefix := str[:j]
			reversePrefix := reverseString(prefix)
			suffix := str[j:]
			reverseSuffix := reverseString(suffix)

			// i first
			if k, ok := strMap[reversePrefix]; ok && isPalindrome(suffix) && i != k {
				res = append(res, []int{i, k})
			}
			// i second. We use len(prefix) != 0 to dedupe because the case where duplication happen is
			// two string has the same length and for str1, it would match with str2 twice when the prefix
			// is str1 or the suffix is str
			if k, ok := strMap[reverseSuffix]; ok && isPalindrome(prefix) && i != k && len(prefix) != 0 {
				res = append(res, []int{k, i})
			}
		}
	}
	return res
}

func isPalindrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverseString(str string) string {
	bytes := []byte(str)
	i, j := 0, len(str)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

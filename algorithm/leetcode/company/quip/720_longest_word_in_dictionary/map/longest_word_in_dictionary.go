package _map

// Time: O(m*n) - m is # of word. n is average length of words.
// Space: O(m)
func longestWord(words []string) string {
	strMap := map[string]bool{}
	for _, word := range words {
		strMap[word] = true
	}
	var res string
	for _, word := range words {
		// check all prefix is in word list
		if len(word) > len(res) || (len(word) == len(res) && word < res) {
			valid := true
			for i := 0; i < len(word); i++ {
				if !strMap[word[:i+1]] {
					valid = false
				}
			}
			if valid {
				res = word
			}
		}
	}
	return res
}

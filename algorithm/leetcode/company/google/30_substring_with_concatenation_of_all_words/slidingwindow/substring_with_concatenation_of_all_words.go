package slidingwindow

// For every start point, we move len(word) and use sliding window to keep track of
// the words in the window. After we process a start point, we move to the next start point
// by increase by 1. We only have len(word) start points.
// Reference: https://leetcode.wang/leetCode-30-Substring-with-Concatenation-of-All-Words.html
// Time: O(s) - every bytes in s only get visited once and when we visit every byte, we only take
//              the substring of len(word) which should be O(1) or O(word)??
// Space: O(k*w)
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 || len(s) < len(words[0])*len(words) {
		return []int{}
	}
	wordMap := map[string]int{}
	for _, word := range words {
		wordMap[word]++
	}
	wordNum := len(wordMap)
	wordLen := len(words[0])
	totalLen := wordLen * (len(words))
	res := []int{}
	for i := 0; i < wordLen; i++ {
		windowNum := 0
		windowMap := map[string]int{}
		for j := i; j <= len(s)-wordLen; j += wordLen {
			if j-i >= totalLen {
				start := j - totalLen
				removedWord := s[start : start+wordLen]
				if wordMap[removedWord] > 0 {
					windowMap[removedWord]--
					if windowMap[removedWord] < wordMap[removedWord] {
						windowNum--
					}
				}

			}
			w := s[j : j+wordLen]
			if wordMap[w] > 0 {
				windowMap[w]++
				if windowMap[w] == wordMap[w] {
					windowNum++
				}
			}
			if windowNum == wordNum {
				res = append(res, j-totalLen+wordLen)
			}
		}
	}
	return res
}

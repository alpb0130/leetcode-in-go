package _map

// Start from every position of the given string and we check teh substring in the
// window. The window length is len(words)*len(words[0]). If there are word in substring
// that is not in words or more word, we jump out of the for loop
// Time: O(w+s*w) - w is # of words and s is length of s
// Space: O(w*k) - k is length of every word
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
	for i := 0; i <= len(s)-totalLen; i++ {
		windowNum := 0
		windowMap := map[string]int{}
		for j := i; j < i+totalLen; j += wordLen {
			w := s[j : j+wordLen]
			windowMap[w]++
			if windowMap[w] > wordMap[w] {
				break
			}
			if windowMap[w] == wordMap[w] {
				windowNum++
			}
		}
		if windowNum == wordNum {
			res = append(res, i)
		}
	}
	return res
}

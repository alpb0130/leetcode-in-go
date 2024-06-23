package improvedslidingwindow

// Similar to the general sliding window method but we extract all the letters (and index) in s
// which are also in t so that we can get a smaller slice to iterate over.
// Time: O(# chars in s which are in T as well)
// Space: O(# unique chars in s and t) + O(# chars in s which are in T as well)
func minWindow(s string, t string) string {
	charMap := map[byte]int{}
	for _, c := range t {
		charMap[byte(c)]++
	}
	charNum := len(charMap)

	processedS := []character{}
	for i, c := range s {
		if charMap[byte(c)] > 0 {
			processedS = append(processedS, character{byte(c), i})
		}
	}
	windowMap := map[byte]int{}
	windowNum := 0
	min := len(s) + 1
	var l, r int
	for start, end := 0, 0; end < len(processedS); end++ {
		char := processedS[end]
		windowMap[char.c]++
		if windowMap[char.c] == charMap[char.c] {
			windowNum++
		}
		for start <= end && windowNum == charNum {
			startChar := processedS[start]
			endChar := processedS[end]
			if endChar.idx-startChar.idx+1 < min {
				min = endChar.idx - startChar.idx + 1
				l = startChar.idx
				r = endChar.idx
			}
			windowMap[startChar.c]--
			if windowMap[startChar.c] < charMap[startChar.c] {
				windowNum--
			}
			start++
		}
	}
	if min == len(s)+1 {
		return ""
	}
	return s[l : r+1]
}

type character struct {
	c   byte
	idx int
}

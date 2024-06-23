package linear

import "math"

// The input two words could be the same and they must be in the words list.
// One pass and use two vars to keep track of the index of matched words.
// Every time we update the index, we update the result then. We need to take
// care of the case where two words are the same. If they are same, still use
// two pointer but p1 must be the word whose index is larger.
// Time: O(n)
// Space: O(1)
func shortestWordDistance(words []string, word1 string, word2 string) int {
	p1, p2 := -1, -1
	res := len(words)
	for i, w := range words {
		if w == word1 && w != word2 {
			p1 = i
		} else if w == word2 && w != word1 {
			p2 = i
		} else if w == word1 && w == word2 {
			// deal with the case where two words are same
			p2 = p1
			p1 = i
		} else {
			continue
		}
		if p1 != -1 && p2 != -1 {
			res = minInt(res, absInt(p1-p2))
		}
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func absInt(a int) int {
	return int(math.Abs(float64(a)))
}

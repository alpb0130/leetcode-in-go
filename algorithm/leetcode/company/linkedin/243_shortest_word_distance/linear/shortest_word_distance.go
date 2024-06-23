package linear

import "math"

// The input two words are different and they must be in the words list.
// One pass and use two vars to keep track of the index of matched words.
// Every time we update the index, we update the result then.
// Time: O(n)
// Space: O(1)
func shortestDistance(words []string, word1 string, word2 string) int {
	p1, p2 := -1, -1
	res := len(words)
	for i, w := range words {
		if w == word1 {
			p1 = i
		} else if w == word2 {
			p2 = i
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

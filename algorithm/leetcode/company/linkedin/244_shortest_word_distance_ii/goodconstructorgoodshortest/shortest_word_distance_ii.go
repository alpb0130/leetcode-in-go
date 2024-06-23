package goodconstructorgoodshortest

import "math"

type WordDistance struct {
	wordMap  map[string][]int
	wordsLen int
}

// All solutions:
// 1. Best constructor: O(1). Worst Shortest: O(n) - only store the original words list.
//    Use method in I to get shortest distance.
// 2. Good constructor: O(n). Good Shortest: O(min(s, t))
// 3. Worst constructor: O(n!). Best Shortest: O(1) - precalculate distance between all word
//    pairs and constant time to get shortest distance. But space is super bad.

// The input two words are different and they must be in the words list.
// This method is a trade-off method. Constructor is good but not best. Shortest is good but
// not best either.
// Use a map to store all the index for every word and when Shortest is call, we only need to
// process the two index slice.
// Time: O(n)
// Space: O(n)
func Constructor(words []string) WordDistance {
	wordMap := map[string][]int{}
	for i, word := range words {
		if _, ok := wordMap[word]; ok {
			wordMap[word] = append(wordMap[word], i)
		} else {
			wordMap[word] = []int{i}
		}
	}
	return WordDistance{
		wordMap:  wordMap,
		wordsLen: len(words),
	}
}

// Time: O(min(s, t)) where s, t is the number of word1 and word2 in the original input words list.
func (this *WordDistance) Shortest(word1 string, word2 string) int {
	list1 := this.wordMap[word1]
	list2 := this.wordMap[word2]
	i, j := 0, 0
	res := this.wordsLen
	// No need to check another list if one list has been all processed.
	// If one list has all been process, which means the last index of the list
	// is smaller than the current index of the unfinished list. No need to keep
	// extending the diff.
	for i < len(list1) && j < len(list2) {
		res = minInt(res, absInt(list1[i]-list2[j]))
		if list1[i] < list2[j] {
			i++
		} else {
			j++
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

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */

package greedyandexpandprefix

import "strconv"

// We firstly generate the best abbreviations for every word. If we find duplication from the
// abbre list, we take them out and expand the prefix and try to fix duplicate. We keep expanding
// the prefix until there is no duplicate from the set. Then we can move to the next abbre and try
// to remove duplicate.
// Better solution (group and trie): https://leetcode.com/problems/word-abbreviation/solution/
// Time: O(C^2)
// Space: O(C)
func wordsAbbreviation(dict []string) []string {
	if len(dict) == 0 {
		return []string{}
	}
	res := []string{}
	for _, word := range dict {
		res = append(res, getAbbre(word, 0))
	}
	prefix := make([]int, len(dict))
	for i := 0; i < len(dict); i++ {
		for {
			dup := []int{}
			for j := i + 1; j < len(dict); j++ {
				if res[i] == res[j] {
					dup = append(dup, j)
				}
			}
			if len(dup) == 0 {
				break
			}
			dup = append(dup, i)
			for _, d := range dup {
				prefix[d] += 1
				res[d] = getAbbre(dict[d], prefix[d])
			}
		}
	}
	return res
}

func getAbbre(word string, i int) string {
	if len(word) <= 3 || len(word)-i-2 == 1 {
		return word
	}
	return word[:i+1] + strconv.Itoa(len(word)-i-2) + string(word[len(word)-1])
}

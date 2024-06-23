package _map

import "strings"

// The indexed list is not in order, be careful.
// Time: O(len(S))
// Space: O(n) - n is # of replacement
func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	if len(S) == 0 || len(indexes) == 0 {
		return S
	}
	// build source-target map
	indexMap := map[int][]string{}
	for i := 0; i < len(indexes); i++ {
		indexMap[indexes[i]] = []string{sources[i], targets[i]}
	}
	var res strings.Builder
	i := 0
	for i < len(S) {
		if indexMap[i] != nil {
			source := indexMap[i][0]
			j := i + len(source)
			if j <= len(S) && S[i:j] == source {
				res.WriteString(indexMap[i][1])
				i = j
			} else {
				res.WriteString(string(S[i]))
				i++
			}
		} else {
			res.WriteString(string(S[i]))
			i++
		}
	}
	return res.String()
}

package bfsanddfsreverse

// Time: O(m*n) - m is the length of wordList, n is the length of every words
// Space: O(m*n) - the map
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	patternMap := map[string][]string{}
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := wordToPattern(word, i)
			patternMap[pattern] = append(patternMap[pattern], word)
		}
	}
	predMap := map[string][]string{}
	q := &queue{beginWord}
	isVisited := map[string]bool{beginWord: true}
	reachEnd := false
	// BFS to build the path to endWord
	for q.Len() > 0 && !reachEnd {
		size := q.Len()
		// We need a temporary visited map to record the visited info for node of
		// this level. We need this because we don't what the node in the same level
		// to be blocked to map to the same word
		tmpIsVisited := map[string]bool{}
		for i := 0; i < size; i++ {
			str := q.Poll()
			for j := 0; j < len(str); j++ {
				pattern := wordToPattern(str, j)
				for _, next := range patternMap[pattern] {
					if !isVisited[next] {
						if !tmpIsVisited[next] {
							q.Offer(next)
						}
						predMap[next] = append(predMap[next], str)
						tmpIsVisited[next] = true
					}
					if next == endWord {
						reachEnd = true
					}
				}
			}
		}
		for str, _ := range tmpIsVisited {
			isVisited[str] = true
		}
	}
	// Build the result recursively
	res := [][]string{}
	buildResults(predMap, endWord, beginWord, []string{}, &res)
	return res
}

func buildResults(predMap map[string][]string, start, end string, list []string, res *[][]string) {
	list = append(list, start)
	if start == end {
		reverse := reverseList(list)
		*res = append(*res, reverse)
	}
	for _, pred := range predMap[start] {
		buildResults(predMap, pred, end, list, res)
	}
}

func reverseList(list []string) []string {
	dst := make([]string, len(list))
	for i := 0; i < len(list); i++ {
		dst[i] = list[len(list)-1-i]
	}
	return dst
}

func wordToPattern(word string, i int) string {
	bytes := []byte(word)
	bytes[i] = '*'
	return string(bytes)
}

type queue []string

func (q *queue) Poll() string {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *queue) Offer(s string) {
	*q = append(*q, s)
}
func (q *queue) Len() int {
	return len(*q)
}

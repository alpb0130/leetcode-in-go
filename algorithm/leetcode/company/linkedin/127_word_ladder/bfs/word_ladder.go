package bfs

// The problem requires to find the length of min sequence to convert beginWord to
// endWord. For each change, we can only change one letter and the word after changing
// should be int the word list. If we regard the word conversion relationship as an
// unweighted undirected graph, the problem is convert to use dfs to find the shortest
// path from one node to another. The hard part is how to construct the connection.
// We use "*" to represent all possible pattern for a word (replace letters with "*"), then
// every word would have n patterns (n is the length of the word). The all words share the
// same pattern could be in the same list and we can easily find the adjacent string (node).
// Time: O(m*n) - m is the length of wordList, n is the length of every words
// Space: O(m*n)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// build matching pattern
	stringPattern := map[string][]string{}
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			bytes := []byte(word)
			bytes[i] = '*'
			pattern := string(bytes)
			if _, ok := stringPattern[pattern]; ok {
				stringPattern[pattern] = append(stringPattern[pattern], word)
			} else {
				stringPattern[pattern] = []string{word}
			}
		}
	}
	isVisited := map[string]bool{beginWord: true}
	dist := 1
	q := &queue{beginWord}
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			str := q.Poll()
			for j := 0; j < len(str); j++ {
				bytes := []byte(str)
				bytes[j] = '*'
				pattern := string(bytes)
				nextStr := stringPattern[pattern]
				for _, next := range nextStr {
					if next == endWord {
						return dist + 1
					} else if !isVisited[next] {
						isVisited[next] = true
						q.Offer(next)
					}
				}
			}
		}
		dist++
	}
	return 0
}

type queue []string

func (q *queue) Poll() string {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
func (q *queue) Offer(x string) {
	*q = append(*q, x)
}
func (q *queue) Len() int {
	return len(*q)
}

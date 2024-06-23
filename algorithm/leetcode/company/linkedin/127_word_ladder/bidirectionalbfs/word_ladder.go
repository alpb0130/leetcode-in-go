package bidirectionalbfs

// Similar to general bfs but search from both end so that we can avoid fanout as much as possible.
// Reference: https://leetcode.com/problems/word-ladder/solution/
// Time: O(m*n) - m is the length of wordList, n is the length of every words
// Space: O(m*n)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// build matching pattern
	stringPattern := map[string][]string{}
	hasEndWord := false
	for _, word := range wordList {
		if word == endWord {
			hasEndWord = true
		}
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
	if !hasEndWord {
		return 0
	}
	beginQueue := &queue{beginWord}
	endQueue := &queue{endWord}
	beginDist := map[string]int{beginWord: 1}
	endDist := map[string]int{endWord: 1}
	for beginQueue.Len() != 0 && endQueue.Len() != 0 {
		dist := nextLevel(beginQueue, beginDist, endDist, stringPattern)
		if dist != 0 {
			return dist
		}
		dist = nextLevel(endQueue, endDist, beginDist, stringPattern)
		if dist != 0 {
			return dist
		}
	}
	return 0
}

func nextLevel(q *queue, dist, counterpartDist map[string]int, stringPattern map[string][]string) int {
	size := q.Len()
	for i := 0; i < size; i++ {
		str := q.Poll()
		d := dist[str]
		for j := 0; j < len(str); j++ {
			bytes := []byte(str)
			bytes[j] = '*'
			pattern := string(bytes)
			nextStr := stringPattern[pattern]
			for _, next := range nextStr {
				if counterpartDist[next] != 0 {
					return d + counterpartDist[next]
				} else if dist[next] == 0 {
					dist[next] = d + 1
					q.Offer(next)
				}
			}
		}
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

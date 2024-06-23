package bfsanddfs

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
	neighborMap := map[string][]string{}
	distMap := map[string]int{beginWord: 1}
	q := &queue{beginWord}
	isVisited := map[string]bool{beginWord: true}
	dist := 1
	reachEnd := false
	for q.Len() > 0 && !reachEnd {
		size := q.Len()
		for i := 0; i < size; i++ {
			str := q.Poll()
			for j := 0; j < len(str); j++ {
				pattern := wordToPattern(str, j)
				for _, next := range patternMap[pattern] {
					if next != str {
						neighborMap[str] = append(neighborMap[str], next)
					}
					if !isVisited[next] {
						q.Offer(next)
						isVisited[next] = true
						distMap[next] = dist + 1
					}
					if next == endWord {
						reachEnd = true
					}
				}
			}
		}
		dist++
	}
	res := [][]string{}
	buildResults(neighborMap, distMap, beginWord, endWord, []string{}, &res)
	return res
}

func buildResults(neighborMap map[string][]string, distMap map[string]int, start, end string, list []string, res *[][]string) {
	list = append(list, start)
	if start == end {
		dst := make([]string, len(list))
		copy(dst, list)
		*res = append(*res, dst)
	}
	for _, neighbor := range neighborMap[start] {
		if distMap[start]+1 == distMap[neighbor] {
			buildResults(neighborMap, distMap, neighbor, end, list, res)
		}
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

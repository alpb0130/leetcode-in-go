package bfstopologicalsort

import "strings"

// Construct the relationship graph and use bfs to get the topology order.
// It's very important to include all chars instead the ones we can only figure out the order.
// So the first for loop is to make all chars as keys
// Time: O(m*n) - m is # words. n is average length of word. Construct the initial map, we need m*n
//       Time. Build relationship, we need m*n. BFS to get the topological sort, we need O(V+E). E
//       is bounded by m because we can only extract at most O(m) relationship from m words. V is
//       bounded by m*n.
// Space: O(m*n) - The map
func alienOrder(words []string) string {
	letterMap := map[byte][]byte{}
	degrees := map[byte]int{}
	// build graph
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if letterMap[word[i]] == nil {
				letterMap[word[i]] = []byte{}
				degrees[word[i]] = 0
			}
		}
	}
	for i := 1; i < len(words); i++ {
		r := getRelationship(words[i-1], words[i])
		if len(r) == 2 {
			letterMap[r[0]] = append(letterMap[r[0]], r[1])
			degrees[r[1]]++
		}
	}
	q := &queue{}
	for b, degree := range degrees {
		if degree == 0 {
			q.Offer(b)
		}
	}

	var res strings.Builder
	for q.Len() > 0 {
		b := q.Poll()
		res.WriteString(string(b))
		for _, next := range letterMap[b] {
			degrees[next]--
			if degrees[next] == 0 {
				q.Offer(next)
			}
		}
	}
	if res.Len() == len(letterMap) {
		return res.String()
	}
	return ""
}

func getRelationship(w1, w2 string) []byte {
	for i := 0; i < len(w1) && i < len(w2); i++ {
		if w1[i] != w2[i] {
			return []byte{w1[i], w2[i]}
		}
	}
	return []byte{}
}

type queue []byte

func (q *queue) Len() int {
	return len(*q)
}
func (q *queue) Offer(x byte) {
	*q = append(*q, x)
}
func (q *queue) Poll() byte {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

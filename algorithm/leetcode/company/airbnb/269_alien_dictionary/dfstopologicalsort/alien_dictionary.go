package dfstopologicalsort

// Construct the relationship graph and use dfs to get the topology order (reverse of DFS
// postorder).
// It's very important to include all chars instead the ones we can only figure out the order.
// So the first for loop is to make` all chars as keys
// Time: O(m*n) - m is # words. n is average length of word. Construct the initial map, we need m*n
//       Time. Build relationship, we need m*n. DFS to get the topological sort, we need O(V+E) E
//       is bounded by m because we can only extract at most O(m) relationship from m words. V is
//       bounded by m*n.
// Space: O(m*n) - The map
func alienOrder(words []string) string {
	letterMap := map[byte][]byte{}
	// build graph
	// O(m*n)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if letterMap[word[i]] == nil {
				letterMap[word[i]] = []byte{}
			}
		}
	}
	// O(m*n)
	for i := 1; i < len(words); i++ {
		r := getRelationship(words[i-1], words[i])
		if len(r) == 2 {
			letterMap[r[0]] = append(letterMap[r[0]], r[1])
		}
	}
	// dfs to get the order
	isVisited := map[byte]bool{}
	res := []byte{}
	for c, _ := range letterMap {
		if !isVisited[c] {
			if !dfs(letterMap, c, isVisited, map[byte]bool{}, &res) {
				return ""
			}
		}
	}
	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func getRelationship(w1, w2 string) []byte {
	for i := 0; i < len(w1) && i < len(w2); i++ {
		if w1[i] != w2[i] {
			return []byte{w1[i], w2[i]}
		}
	}
	return []byte{}
}

func dfs(letterMap map[byte][]byte, start byte, isVisited, onStack map[byte]bool, res *[]byte) bool {
	isVisited[start] = true
	onStack[start] = true
	adj := letterMap[start]
	for _, c := range adj {
		if onStack[c] {
			return false
		}
		if !isVisited[c] && !dfs(letterMap, c, isVisited, onStack, res) {
			return false
		}
	}
	*res = append(*res, start)
	onStack[start] = false
	return true
}

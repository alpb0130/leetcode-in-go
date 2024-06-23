package eularpath

import "strings"

// Find Euler path using Hierholzer's algorithm. One thing to notice is we actually not adding
// the node into bytes array, we actually add edge into byte array because we need it instead of
// node.
// Time: O(k^n)
// Space: O(k^(n-1)*k) = O(k^n)
func crackSafe(n int, k int) string {
	if n == 0 || k == 0 {
		return ""
	}
	var start strings.Builder
	for i := 0; i < n-1; i++ {
		start.WriteString("0")
	}
	isVisited := map[string]bool{start.String(): true}
	// DFS
	bytes := []byte{}
	dfs(start.String(), n, k, isVisited, &bytes)
	reverseBytes(bytes)
	return start.String() + string(bytes)
}

func reverseBytes(bytes []byte) {
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
}

func dfs(node string, n, k int, isVisited map[string]bool, bytes *[]byte) {
	for i := 0; i < k; i++ {
		number := string('0' + i)
		next := node + number
		if !isVisited[next] {
			isVisited[next] = true
			dfs(next[1:], n, k, isVisited, bytes)
			*bytes = append(*bytes, byte('0'+i))
		}
	}
}

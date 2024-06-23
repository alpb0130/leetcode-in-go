package trieandbacktracking

// The problem gives a letter board and a list of words. We are required to find all word in the
// list that can be constructed by searching from the board. We can brute force the problem by
// DFS the board for every word. But it takes a lot of time and we do many repeated searches.
// Like if two words have shared prefix, we need to do the same search twice.
// We can use trie to save those info. We have two benefits by using trie:
// 1. Search words sharing the same prefix at the same time.
// 2. Limit search over the board to existing words only. We can prune the search if we find
//    the next letter is not in the trie.
// First we build the trie first from the word list which takes O(l*k) - l is # words, k is average
// length of word.
// Then we do dfs search from every position (m*n) over the trie. For each position, we have 4
// directions. We only search the direction which there is a node in children. This would limit
// our search space a lot.
// Time: O(m*n*3^(average word length)) - worst case time. The worst case is the give board is
//       [["a","a"],["a","a"]] and the word list is: ["aaab", "aaac", "aaad"]
// Space: O(m*n) - the search call depth, the trie and the visit map
func findWords(board [][]byte, words []string) []string {
	res := []string{}
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return res
	}
	root := buildTrie(words)
	// search
	m := len(board)
	n := len(board[0])
	diffs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if root.children[board[i][j]] != nil {
				dfs(board, i, j, m, n, map[int]bool{i*n + j: true}, diffs, root.children[board[i][j]], &res)
			}
		}
	}
	return res
}

func dfs(board [][]byte, i, j, m, n int, isVisited map[int]bool, diffs [][]int, root *TrieNode, res *[]string) {
	if root.isWord {
		*res = append(*res, root.word)
		// dedupe
		root.isWord = false
	}
	for _, diff := range diffs {
		ii := i + diff[0]
		jj := j + diff[1]
		if isVisited[ii*n+jj] || !validIndex(ii, jj, m, n) || root.children[board[ii][jj]] == nil {
			continue
		}
		isVisited[ii*n+jj] = true
		dfs(board, ii, jj, m, n, isVisited, diffs, root.children[board[ii][jj]], res)
		isVisited[ii*n+jj] = false
	}
}

func validIndex(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}

type TrieNode struct {
	isWord bool
	// store actually word so that we can easily add a word to result set without
	// tracking current word bytes.
	word     string
	children map[byte]*TrieNode
}

func buildTrie(words []string) *TrieNode {
	root := newNode()
	for _, word := range words {
		addWord(root, word)
	}
	return root
}

func addWord(root *TrieNode, word string) {
	cur := root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if cur.children[c] == nil {
			cur.children[c] = newNode()
		}
		cur = cur.children[c]
	}
	cur.isWord = true
	cur.word = word
}

func newNode() *TrieNode {
	return &TrieNode{
		children: map[byte]*TrieNode{},
	}
}

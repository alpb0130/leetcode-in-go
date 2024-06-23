package dfsandtrie

import "strings"

// Trie and DFS. One import tip is all the words have same length.
// Time: O(n*k+n^k) - this is the upper bound. n is number of words and k is the length of word.
//                    The actual time should be less than that because we have trie to help use do
//                    pruning.
// Space: O(???)
func wordSquares(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}
	n := len(words[0])
	tree := &Trie{NewTrieNode(' ')}
	tree.AddWords(words)
	res := [][]string{}
	// dfs
	dfs(tree, n, []string{}, &res)
	return res
}

func dfs(tree *Trie, n int, cur []string, res *[][]string) {
	if len(cur) == n {
		dst := make([]string, n)
		copy(dst, cur)
		*res = append(*res, dst)
		return
	}
	m := len(cur)
	var prefix strings.Builder
	for i := 0; i < m; i++ {
		prefix.WriteString(string(cur[i][m]))
	}
	next := tree.GetWords(prefix.String())
	for _, w := range next {
		cur = append(cur, w)
		dfs(tree, n, cur, res)
		cur = cur[:m]
	}
}

type TrieNode struct {
	Char  byte
	Words []string
	Word  string
	Next  map[byte]*TrieNode
}

func NewTrieNode(c byte) *TrieNode {
	return &TrieNode{
		Char:  c,
		Words: []string{},
		Word:  "",
		Next:  map[byte]*TrieNode{},
	}
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) AddWords(words []string) {
	for _, word := range words {
		t.addWord(word)
	}
}

func (t *Trie) addWord(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		node.Words = append(node.Words, word)
		next, ok := node.Next[word[i]]
		if !ok {
			node.Next[word[i]] = NewTrieNode(word[i])
			next = node.Next[word[i]]
		}
		node = next
	}
	node.Word = word
	node.Words = append(node.Words, word)
}

func (t *Trie) GetWords(prefix string) []string {
	node := t.root
	for i := 0; i < len(prefix); i++ {
		next, ok := node.Next[prefix[i]]
		if !ok {
			return []string{}
		}
		node = next
	}
	return node.Words
}

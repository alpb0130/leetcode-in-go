package trie

// Time: O(n*k^2) - m is the number of words because we need to check every words. k is the
//       avg length of word. For every prefix substring of a word, it's a word, we need to
//       check whether remaining word is a palindrome at line 40. So for each word, the time
//       would be (n-1)+(n-2)+...+1. Build trie also take time.
// Space: O(n*k) - m is the number of words. n is the avg length of word
func palindromePairs(words []string) [][]int {
	res := [][]int{}
	if len(words) <= 1 {
		return res
	}
	root := createTrie(words)
	for i, word := range words {
		searchWordAndAddToRes(root, reverseStr(word), i, &res)
	}
	return res
}

type TrieNode struct {
	c        byte
	children map[byte]*TrieNode
	isEnd    bool
	index    int
	list     []int
}

func searchWordAndAddToRes(root *TrieNode, bytes []byte, index int, res *[][]int) {
	if root == nil {
		return
	}
	if len(bytes) == 0 {
		for _, idx := range root.list {
			if idx != index {
				*res = append(*res, []int{idx, index})
			}
		}
		return
	}
	if root.isEnd && isPalindrome(bytes) && root.index != index {
		*res = append(*res, []int{root.index, index})
	}
	c := bytes[0]
	child := root.children[c]
	searchWordAndAddToRes(child, bytes[1:], index, res)
}

func createTrie(words []string) *TrieNode {
	root := newNode(' ')
	for i, word := range words {
		addWord(root, []byte(word), i)
	}
	return root
}

func addWord(root *TrieNode, word []byte, index int) {
	if len(word) == 0 {
		root.isEnd = true
		root.index = index
		root.list = append(root.list, index)
		return
	}
	if isPalindrome(word) {
		root.list = append(root.list, index)
	}
	c := word[0]
	if root.children[c] == nil {
		root.children[c] = newNode(c)
	}
	addWord(root.children[c], word[1:], index)
}

func newNode(c byte) *TrieNode {
	return &TrieNode{
		c:        c,
		children: map[byte]*TrieNode{},
		index:    -1,
		list:     []int{},
	}
}

func isPalindrome(bytes []byte) bool {
	i, j := 0, len(bytes)-1
	for i < j {
		if bytes[i] != bytes[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverseStr(str string) []byte {
	bytes := []byte(str)
	i, j := 0, len(str)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return bytes
}

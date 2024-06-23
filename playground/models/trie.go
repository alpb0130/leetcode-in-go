package models

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			Next: map[rune]*TrieNode{},
		},
	}
}

func (t *Trie) Encode() string {
	return t.EncodeNode(t.root)
}

func (t *Trie) EncodeNode(node *TrieNode) string {
	if node == nil {
		return ""
	}
	var str strings.Builder
	str.WriteString("{")
	str.WriteString("rune: " + string(node.Char) + ", ")
	str.WriteString("isWord: " + fmt.Sprintf("%v", node.IsWord) + ", ")
	for r, node := range node.Next {
		str.WriteString(fmt.Sprintf("%s: %s, ", string(r), t.EncodeNode(node)))
	}
	str.WriteString("}")
	return str.String()

}

func (t *Trie) AddWord(line string, title string) {
	root := t.root
	for len(line) > 0 {
		runeValue, width := utf8.DecodeRuneInString(line)
		if node, ok := root.Next[runeValue]; ok {
			root = node
		} else {
			root.Next[runeValue] = &TrieNode{
				Char:   runeValue,
				IsWord: false,
				Next:   map[rune]*TrieNode{},
			}
			root = root.Next[runeValue]
		}
		line = line[width:]
	}
	root.IsWord = true
	root.Title = title
}

func (t *Trie) SearchPrefix(prefix string) []string {
	root := t.root
	if root == nil {
		return []string{}
	}

	p := prefix
	// find the last node
	for len(prefix) > 0 {
		runeValue, width := utf8.DecodeRuneInString(prefix)
		if node, ok := root.Next[runeValue]; ok {
			root = node
		} else {
			return []string{}
		}
		prefix = prefix[width:]
	}

	// root is the last prefix node
	// dfs
	res := map[string]bool{}
	t.getAllWords(root, []rune(p), res)
	list := []string{}
	for t, _ := range res {
		list = append(list, t)
	}
	return list
}

func (t *Trie) getAllWords(root *TrieNode, cur []rune, res map[string]bool) {
	if root == nil {
		return
	}
	if root.IsWord {
		//dst := make([]rune, len(cur))
		//copy(dst, cur)
		if !res[root.Title] {
			res[root.Title] = true
		}
		//*res = append(*res, root.Title)
	}
	for r, node := range root.Next {
		t.getAllWords(node, append(cur, r), res)
	}
}

package models

type TrieNode struct {
	Char   rune
	IsWord bool
	Title  string
	Next   map[rune]*TrieNode
}

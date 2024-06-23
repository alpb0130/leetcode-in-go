package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"playground/models"
	"strings"
)

func main() {
	/*trie := models.NewTrie()
	trie.AddWord("abc")
	trie.AddWord("abd")
	trie.AddWord("a")
	trie.AddWord("bdd")
	strs := trie.SearchPrefix("abcd")
	fmt.Println(strs)*/

	file, err := os.Open("/Users/alpb0130/Desktop/test1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	trie := models.NewTrie()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		title := scanner.Text()
		for _, word := range strings.Split(title, " ") {
			trie.AddWord(word, title)
		}

	}
	//title := "abc abe"
	//for _, word := range strings.Split(title, " ") {
	//	trie.AddWord(word, title)
	//}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter prefix: ")
	prefix, _ := reader.ReadString('\n')
	prefix = prefix[:len(prefix)-1]
	res := map[string]int{}
	for _, pre := range strings.Split(prefix, " ") {
		tmp := trie.SearchPrefix(pre)
		for _, title := range tmp {
			res[title]++
		}
	}

	for s, cnt := range res {
		fmt.Println(s, cnt)
	}
}

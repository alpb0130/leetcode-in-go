package heuristic

import (
	"math"
)

type Master struct {
}

func (this *Master) Guess(word string) int { return 1 }

// Every time, we guess the word with minimun possible children for all matches
// Time: O(n*n*6 + n*n*logn)
/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
func findSecretWord(wordlist []string, master *Master) {
	if len(wordlist) == 0 {
		return
	}
	// build word matrix
	wordMatch := make([][]int, len(wordlist))
	for i := range wordMatch {
		wordMatch[i] = make([]int, len(wordlist))
		wordMatch[i][i] = len(wordlist[0])
	}
	guess := map[int]bool{}
	visit := map[int]bool{}
	for i := 0; i < len(wordlist); i++ {
		guess[i] = true
		for j := i + 1; j < len(wordlist); j++ {
			match := countWordMatch(wordlist[i], wordlist[j])
			wordMatch[i][j] = match
			wordMatch[j][i] = match
		}
	}
	// guess
	// This for log decrease exponentially
	for len(guess) > 0 {
		// pick guess
		idx := pick(wordMatch, guess, visit)
		match := master.Guess(wordlist[idx])
		if match == len(wordlist[idx]) {
			return
		}
		guess1 := map[int]bool{}
		visit[idx] = true
		str := []string{}
		for word := range guess {
			// be careful here!!! You only want to add word that has [match] with
			// the chosen word to the guess list
			// Also, you always want to choose the next guess set from the current guess set
			// because if you don't pick from current set, it will not be eligible from our
			// previous validation.
			if !visit[word] && wordMatch[idx][word] == match {
				guess1[word] = true
				str = append(str, wordlist[word])
			}
		}
		guess = guess1
	}
	return
}

// O(n*n) - over guess and over wordMatch
func pick(wordMatch [][]int, guess, visit map[int]bool) int {
	if len(guess) <= 2 {
		for wordIdx := range guess {
			return wordIdx
		}
	}

	min := len(wordMatch)
	idx := 0
	for i := range guess {
		max := 0
		matchCount := make([]int, 7)
		for j, cnt := range wordMatch[i] {
			if !visit[j] && j != i {
				matchCount[cnt]++
				max = maxInt(max, matchCount[cnt])
			}
		}
		if max < min {
			min = max
			idx = i
		}
	}
	return idx
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func countWordMatch(s1, s2 string) int {
	res := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			res++
		}
	}
	return res
}

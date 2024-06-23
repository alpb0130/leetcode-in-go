package twppassii

import "fmt"

// One pass: https://leetcode.com/problems/bulls-and-cows/discuss/74621/One-pass-Java-solution
func getHint(secret string, guess string) string {
	if len(secret) == 0 {
		return ""
	}
	numBulls := 0
	numCows := 0
	secretMap := map[byte]int{}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			numBulls++
		} else {
			secretMap[secret[i]]++
		}
	}
	for i := 0; i < len(guess); i++ {
		if guess[i] != secret[i] && secretMap[guess[i]] > 0 {
			numCows++
			secretMap[guess[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", numBulls, numCows)
}

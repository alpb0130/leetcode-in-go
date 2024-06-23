package twopass

import "fmt"

func getHint(secret string, guess string) string {
	if len(secret) == 0 {
		return ""
	}
	numBulls := 0
	numCows := 0
	secretMap := map[byte]int{}
	for i := 0; i < len(secret); i++ {
		secretMap[secret[i]]++
	}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			numBulls++
			if secretMap[guess[i]] <= 0 {
				numCows--
			}
		} else if secretMap[guess[i]] > 0 {
			numCows++
		}
		secretMap[guess[i]]--
	}
	return fmt.Sprintf("%dA%dB", numBulls, numCows)
}

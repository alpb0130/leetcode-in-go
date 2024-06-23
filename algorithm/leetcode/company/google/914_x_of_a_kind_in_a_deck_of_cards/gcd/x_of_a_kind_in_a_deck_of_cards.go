package gcd

// Get GCD of all deck counts
func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}
	deckMap := map[int]int{}
	for _, d := range deck {
		deckMap[d]++
	}
	counts := []int{}
	for _, c := range deckMap {
		counts = append(counts, c)
	}
	gcd := counts[0]
	for i := 1; i < len(counts); i++ {
		gcd = getGCD(gcd, counts[i])
	}
	return gcd >= 2
}

func getGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return getGCD(b, a%b)
}

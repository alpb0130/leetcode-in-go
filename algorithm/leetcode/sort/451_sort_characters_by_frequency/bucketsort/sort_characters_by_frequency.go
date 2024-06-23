package bucketsort

import "bytes"

// Time: O(n) - n is the length of string.
// Space: O(1)
func frequencySort(s string) string {
	frequency := map[rune]int{}
	for _, c := range s {
		frequency[c]++
	}
	freqCharsMap := map[int][]rune{}
	maxFreq := 0
	for c, f := range frequency {
		chars, ok := freqCharsMap[f]
		if ok {
			freqCharsMap[f] = append(chars, c)
		} else {
			freqCharsMap[f] = []rune{c}
		}
		if f > maxFreq {
			maxFreq = f
		}
	}
	// Access the bucket sort result - iterate from high freq to low freq
	var str bytes.Buffer
	for i := maxFreq; i > 0; i-- {
		chars := freqCharsMap[i]
		for _, c := range chars {
			for j := 0; j < i; j++ {
				str.WriteString(string(c))
			}
		}
	}
	return str.String()
}

package sort

import (
	"bytes"
	"sort"
)

// Time: O(n) - n is the length of string. Because we are always sorting a list with constant length,
// the time of sorting is O(1)
// Space: O(1)
func frequencySort(s string) string {
	frequency := map[rune]int{}
	for _, c := range s {
		frequency[c]++
	}
	frequencyList := freqList{}
	for c, f := range frequency {
		frequencyList = append(frequencyList, charFreq{c, f})
	}
	sort.Sort(frequencyList)
	var str bytes.Buffer
	for _, charFreq := range frequencyList {
		for i := 0; i < charFreq.freq; i++ {
			str.WriteString(string(charFreq.char))
		}
	}
	return str.String()
}

type charFreq struct {
	char rune
	freq int
}

type freqList []charFreq

func (f freqList) Len() int { return len(f) }
func (f freqList) Less(i, j int) bool {
	return f[i].freq > f[j].freq
}
func (f freqList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
